package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/TanishqM1/SocialContentDistributer/api"
	"github.com/TanishqM1/SocialContentDistributer/internal/tools"
)

var wg = sync.WaitGroup{}

func PostContent(w http.ResponseWriter, r *http.Request) {
	var params = api.TotalFields{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Error(err)
		api.HandleRequestError(w, err)
		return
	}

	// now params has all the values from our JSON.
	// we need params.platforms (an array) and we need to create a struct for each platform that lives inside there.

	// We append these finished, made structs to a "our_structs" array. Then, we loop through this array & call SendAPI() on each one (which does the underlying logic like building the api too.)
	// this is it! additional steps to add are concurrency and error logging

	uploads, err := BuildUploadStructs(params)
	if err != nil {
		log.Error(err)
		api.HandleInternalError(w)
		return
	}
	// no we have an "uploads" folder with struct objects. We need to call SendAPI() on all of these struct objects.
	for _, v := range uploads {
		wg.Add(1)
		go tools.SendAPI(v, &wg)
	}
	wg.Wait()
	fmt.Printf("\n All Completed!")
	// once all the api uploads are done (running concurrently), we can send the success response back to the frontend.

	// Send success response back to frontend
	response := map[string]interface{}{
		"success":   true,
		"message":   "Content uploaded successfully",
		"platforms": params.Platforms,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func BuildUploadStructs(params api.TotalFields) ([]tools.UploadContent, error) {
	var uploads []tools.UploadContent

	for _, p := range params.Platforms {
		switch p {
		case "youtube":
			fmt.Println("Building Youtube Struct!")
			uploads = append(uploads, tools.YouTubeUploader{
				AccessToken:   "123",
				PlatformName:  "youtube",
				Title:         params.Title,
				Description:   params.Description,
				Tags:          params.Tags,
				CategoryID:    params.CategoryID,
				PrivacyStatus: params.PrivacyStatus,
				MediaFile:     params.MediaFile,
			})

		case "instagram":
			uploads = append(uploads, tools.InstagramUploader{
				AccessToken:  "123",
				PlatformName: "instagram",
				ImageURL:     params.ImageURL,
				Caption:      params.Caption,
				LocationID:   params.LocationID,
				UserTags:     params.UserTags,
			})

		case "pinterest":
			uploads = append(uploads, tools.PinterestUploader{
				AccessToken:  "123",
				PlatformName: "pinterest",
				BoardID:      params.BoardID,
				Title:        params.Title,
				Description:  params.Description,
				Link:         params.Link,
				SourceType:   params.SourceType,
				ImageURL:     params.ImageURL,
			})

		case "reddit":
			uploads = append(uploads, tools.RedditUploader{
				AccessToken:  "123",
				PlatformName: "reddit",
				Subreddit:    params.Subreddit,
				PostType:     params.PostType,
				Title:        params.Title,
				Text:         params.Text,
				URL:          params.URL,
				Resubmit:     params.Resubmit,
				NSFW:         params.NSFW,
			})

		case "linkedin":
			uploads = append(uploads, tools.LinkedInUploader{
				AccessToken:    "123",
				PlatformName:   "linkedin",
				Author:         params.Author,
				LifecycleState: params.LifecycleState,
				Text:           params.TextLinkedIn,
				MediaType:      params.MediaType,
				MediaStatus:    params.MediaStatus,
				MediaPath:      params.MediaPath,
				Visibility:     params.Visibility,
			})
		}
	}

	return uploads, nil
}
