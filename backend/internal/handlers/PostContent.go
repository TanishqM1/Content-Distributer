package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/TanishqM1/SocialContentDistributer/api"
	"github.com/TanishqM1/SocialContentDistributer/internal/tools"
)

func PostContent(w http.ResponseWriter, r *http.Request) {
	var params = api.TotalFields{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Error(err)
		api.HandleRequestError(w, err)
	}

	// now params has all the values from our JSON.
	// we need params.platforms (an array) and we need to create a struct for each platform that lives inside there.

	// We append these finished, made structs to a "our_structs" array. Then, we loop through this array & call SendAPI() on each one (which does the underlying logic like building the api too.)
	// this is it! additional steps to add are concurrency and error logging

	uploads := []tools.UploadContent{}
	fmt.Println(uploads)

	// for p := range params.Platforms{
	// 	if p == "Youtube"{
	// 		uploads = append(uploads, tools.YouTubeUploader{
	// 			AccessToken: "123",
	// 			PlatformName: "Youtube",
	// 			Title: ,
	// 		})
	// 	}
	// }
}
