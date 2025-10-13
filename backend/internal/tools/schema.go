package tools

// These structs represent the data needed to upload to each platform.
// You are no longer depending on JSON tags because you’ll handle decoding manually
// (e.g., via schema.Decoder, r.ParseForm, or json.Unmarshal in your own logic).

// ===== YouTube =====
type YouTubeUploader struct {
	AccessToken   string
	PlatformName  string
	Title         string
	Description   string
	Tags          []string
	CategoryID    string
	PrivacyStatus string // "public", "private", or "unlisted"
	MediaFile     []byte // binary video data
}

// ===== Instagram =====
type InstagramUploader struct {
	AccessToken  string
	PlatformName string
	ImageURL     string
	Caption      string
	LocationID   string
	UserTags     string
}

// ===== Pinterest =====
type PinterestUploader struct {
	AccessToken  string
	PlatformName string
	BoardID      string
	Title        string
	Description  string
	Link         string
	SourceType   string // e.g., "image_url"
	ImageURL     string
}

// ===== Reddit =====
type RedditUploader struct {
	AccessToken  string
	PlatformName string
	Subreddit    string
	PostType     string // "self", "link", or "image"
	Title        string
	Text         string // used if PostType == "self"
	URL          string // used if PostType == "link" or "image"
	Resubmit     bool
	NSFW         bool
}

// ===== LinkedIn =====
type LinkedInUploader struct {
	AccessToken    string
	PlatformName   string
	Author         string // URN of person/org
	LifecycleState string // "PUBLISHED"
	Text           string
	MediaType      string // "IMAGE" or "VIDEO"
	MediaStatus    string // "READY"
	MediaPath      string // URN or URL
	Visibility     string // "PUBLIC"
}

type UploadContent interface {
	BuildAPI() map[string]interface{} // each struct needs to implement the buildapi function.
}

// need to implement buildapi for each platform!
// we use "interface" as our value type because it can literally hold every value type.

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
	platformname := "temp"
	success := true

	return platformname, success, nil
}
