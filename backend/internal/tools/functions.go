package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
func SendAPI(u UploadContent) (string, bool, error) {
	// implement SendAPI
	body := u.BuildAPI()
	jsonData, _ := json.Marshal(body)
	var endpointURL string
	platform := body["platform_name"]

	switch platform {
	case "youtube":
		endpointURL = "https://www.googleapis.com/upload/youtube/v3/videos?part=snippet,status"
	case "pinterest":
		endpointURL = "https://api.pinterest.com/v5/pins"
	case "reddit":
		endpointURL = "https://oauth.reddit.com/api/submit"
	case "linkedin":
		endpointURL = "https://api.linkedin.com/v2/ugcPosts"
	}

	req, err := http.NewRequest("POST", endpointURL, bytes.NewReader(jsonData))
	req.Header.Set("Authorization", "Bearer "+body["access_token"].(string))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return platform.(string), false, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return platform.(string), false, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return platform.(string), false, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return platform.(string), false, fmt.Errorf("API Request Failed: %s", string(respBody))
	}

	return platform.(string), true, nil
}
