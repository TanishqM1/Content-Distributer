package tools

import (
	"github.com/TanishqM1/SocialContentDistributer/uploads/linkedin"
	"github.com/TanishqM1/SocialContentDistributer/uploads/pinterest"
	"github.com/TanishqM1/SocialContentDistributer/uploads/reddit"
	"github.com/TanishqM1/SocialContentDistributer/uploads/youtube"
)

// youtube
func (y YouTubeUploader) BuildAPI() map[string]interface{} {
	return map[string]interface{}{
		"access_token":  y.AccessToken,
		"platform_name": y.PlatformName,
		"snippet": map[string]interface{}{
			"title":       y.Title,
			"description": y.Description,
			"tags":        y.Tags,
			"categoryId":  y.CategoryID,
		},
		"status": map[string]interface{}{
			"privacyStatus": y.PrivacyStatus,
		},
		"media": map[string]interface{}{
			"file_data": y.MediaFile,
		},
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
		title := body["snippet"].(map[string]interface{})["title"].(string)
		description := body["snippet"].(map[string]interface{})["description"].(string)
		category := body["snippet"].(map[string]interface{})["categoryId"].(string)
		privacy := body["status"].(map[string]interface{})["privacyStatus"].(string)
		filename := body["media"].(map[string]interface{})["file_data"].(string)
		tagslist := body["snippet"].(map[string]interface{})["tags"].([]string)
		var tags string
		for _, v := range tagslist {
			tags += v
		}

		youtube.UploadYoutube(title, description, category, privacy, filename, tags)

	case "pinterest":
		boardID := body["board_id"].(string)
		title := body["title"].(string)
		description := body["description"].(string)
		link := body["link"].(string)
		sourceType := body["media_source"].(map[string]interface{})["source_type"].(string)
		imageURL := body["media_source"].(map[string]interface{})["url"].(string)

		pinterest.UploadPinterest(boardID, title, description, link, sourceType, imageURL)

	case "reddit":
		reddit.UploadReddit()
	case "linkedin":
		linkedin.UploadLinkedIn()

	}

}
