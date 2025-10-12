package main

// These structs represent the data needed to upload to each platform.
// Platforms: YouTube, Instagram, Reddit, Pinterest, LinkedIn

// the `json` is called STRUCT TAGS.

type YouTubeUploader struct {
	AccessToken   string   `json:"access_token"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Tags          []string `json:"tags,omitempty"`
	CategoryID    string   `json:"category_id"`
	PrivacyStatus string   `json:"privacy_status"` // "public", "private", or "unlisted"
	MediaFile     []byte   `json:"media_file"`     // binary video data
}

type InstagramUploader struct {
	AccessToken string `json:"access_token"`
	ImageURL    string `json:"image_url"`
	Caption     string `json:"caption"`
	LocationID  string `json:"location_id,omitempty"`
	UserTags    string `json:"user_tags,omitempty"`
}

type PinterestUploader struct {
	AccessToken string `json:"access_token"`
	BoardID     string `json:"board_id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
	SourceType  string `json:"source_type"` // always "image_url" for MVP
	ImageURL    string `json:"image_url"`
}

type RedditUploader struct {
	AccessToken string `json:"access_token"`
	Subreddit   string `json:"subreddit"`
	PostType    string `json:"post_type"` // "self", "link", or "image"
	Title       string `json:"title"`
	Text        string `json:"text,omitempty"` // used if PostType == "self"
	URL         string `json:"url,omitempty"`  // used if PostType == "link" or "image"
	Resubmit    bool   `json:"resubmit,omitempty"`
	NSFW        bool   `json:"nsfw,omitempty"`
}

type LinkedInUploader struct {
	AccessToken    string `json:"access_token"`
	Author         string `json:"author"`          // URN of person/org
	LifecycleState string `json:"lifecycle_state"` // "PUBLISHED"
	Text           string `json:"text"`
	MediaType      string `json:"media_type,omitempty"`   // e.g. "IMAGE" or "VIDEO"
	MediaStatus    string `json:"media_status,omitempty"` // e.g. "READY"
	MediaPath      string `json:"media_path,omitempty"`   // URN or URL
	Visibility     string `json:"visibility"`             // e.g. "PUBLIC"
}

// we need a method for each to actually build the api
// then an interface, that has 1 method. UploadContent() for ALL platforms.

// example, type UploadContent interface{
// 		BuildAPI()
// 	}

// Then func SendAPI(u UploadContent) string{
// 		logic to actually send api content
// }
