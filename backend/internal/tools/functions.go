package tools

import (
	"fmt"

	"github.com/TanishqM1/SocialContentDistributer/uploads/instagram"
	"github.com/TanishqM1/SocialContentDistributer/uploads/linkedin"
	"github.com/TanishqM1/SocialContentDistributer/uploads/pinterest"
	"github.com/TanishqM1/SocialContentDistributer/uploads/reddit"
	"github.com/TanishqM1/SocialContentDistributer/uploads/youtube"
)

// youtube
func (y YouTubeUploader) BuildAPI() map[string]interface{} {
	return map[string]interface{}{
		"access_token":   y.AccessToken,
		"platform_name":  y.PlatformName,
		"title":          y.Title,
		"description":    y.Description,
		"tags":           y.Tags,
		"category_id":    y.CategoryID,
		"privacy_status": y.PrivacyStatus,
		"media_file":     y.MediaFile,
	}
}

// instagram
func (i InstagramUploader) BuildAPI() map[string]interface{} {
	return map[string]interface{}{
		"access_token":  i.AccessToken,
		"platform_name": i.PlatformName,
		"image_url":     i.ImageURL,
		"caption":       i.Caption,
		"location_id":   i.LocationID,
		"user_tags":     i.UserTags,
	}
}

// pinterest
func (p PinterestUploader) BuildAPI() map[string]interface{} {
	return map[string]interface{}{
		"access_token":  p.AccessToken,
		"platform_name": p.PlatformName,
		"board_id":      p.BoardID,
		"title":         p.Title,
		"description":   p.Description,
		"link":          p.Link,
		"media_source": map[string]interface{}{
			"source_type": p.SourceType, // e.g. "image_url"
			"url":         p.ImageURL,
		},
	}
}

// reddit
func (r RedditUploader) BuildAPI() map[string]interface{} {
	body := map[string]interface{}{
		"access_token":  r.AccessToken,
		"platform_name": r.PlatformName,
		"sr":            r.Subreddit,
		"kind":          r.PostType, // "self", "link", or "image"
		"title":         r.Title,
		"resubmit":      r.Resubmit,
		"nsfw":          r.NSFW,
	}

	// Conditional fields depending on post type
	if r.PostType == "self" {
		body["text"] = r.Text
	} else if r.PostType == "link" || r.PostType == "image" {
		body["url"] = r.URL
	}

	return body
}

// linkedin
func (l LinkedInUploader) BuildAPI() map[string]interface{} {
	return map[string]interface{}{
		"access_token":   l.AccessToken,
		"platform_name":  l.PlatformName,
		"author":         l.Author,
		"lifecycleState": l.LifecycleState, // usually "PUBLISHED"
		"specificContent": map[string]interface{}{
			"com.linkedin.ugc.ShareContent": map[string]interface{}{
				"shareCommentary": map[string]interface{}{
					"text": l.Text,
				},
				"shareMediaCategory": l.MediaType, // "IMAGE" or "VIDEO"
				"media": []map[string]interface{}{
					{
						"status": l.MediaStatus, // "READY"
						"media":  l.MediaPath,   // URN or URL
					},
				},
			},
		},
		"visibility": map[string]interface{}{
			"com.linkedin.ugc.MemberNetworkVisibility": l.Visibility, // "PUBLIC"
		},
	}
}

// basic implementation
func SendAPI(u UploadContent) {
	// implement SendAPI
	body := u.BuildAPI()
	// jsonData, _ := json.Marshal(body) --> this is a byte array!
	platform := body["platform_name"]

	switch platform {
	case "youtube":
		title := getStringValue(body, "title")
		description := getStringValue(body, "description")
		category := getStringValue(body, "category_id")
		privacy := getStringValue(body, "privacy_status")
		filename := getStringValue(body, "media_file")
		tagslist := getStringArrayValue(body, "tags")
		var tags string
		for _, v := range tagslist {
			tags += v
		}

		// Only proceed if we have a valid filename
		if filename != "" && filename != "blank" {
			// Use the file path from uploads folder
			filePath := fmt.Sprintf("uploads/media/%s", filename)
			youtube.UploadYoutube(title, description, category, privacy, filePath, tags)
		} else {
			fmt.Println("Skipping YouTube upload - no valid filename provided")
			fmt.Printf("Filename received: '%s'\n", filename)
		}

	case "instagram":
		imageURL := getStringValue(body, "image_url")
		caption := getStringValue(body, "caption")
		locationID := getStringValue(body, "location_id")
		userTags := getStringValue(body, "user_tags")

		instagram.UploadInstagram(&imageURL, &caption, &locationID, &userTags)

	case "pinterest":
		boardID := body["board_id"].(string)
		title := body["title"].(string)
		description := body["description"].(string)
		link := body["link"].(string)
		sourceType := body["media_source"].(map[string]interface{})["source_type"].(string)
		imageURL := body["media_source"].(map[string]interface{})["url"].(string)

		pinterest.UploadPinterest(boardID, title, description, link, sourceType, imageURL)

	case "reddit":
		subreddit := body["sr"].(string)
		postType := body["kind"].(string)
		title := body["title"].(string)
		resubmit := true
		nsfw := body["nsfw"].(bool)

		var text, url string
		if postType == "self" {
			text = body["text"].(string)
		} else if postType == "link" || postType == "image" {
			url = body["url"].(string)
		}

		reddit.UploadReddit(subreddit, postType, title, text, url, resubmit, nsfw)

	case "linkedin":
		linkedin.UploadLinkedIn()

	}

}

// Helper functions to safely extract values from map
func getStringValue(body map[string]interface{}, key string) string {
	if val, ok := body[key]; ok && val != nil {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func getStringArrayValue(body map[string]interface{}, key string) []string {
	if val, ok := body[key]; ok && val != nil {
		if arr, ok := val.([]string); ok {
			return arr
		}
	}
	return []string{}
}
