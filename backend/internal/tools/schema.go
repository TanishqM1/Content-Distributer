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
	MediaFile     string // binary video data
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
	Link         string // unneeded if locally upoading
	SourceType   string // e.g., "image/jpeg"
	ImageURL     string // unneeded if locally uploading
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
