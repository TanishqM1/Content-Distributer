package api

import (
	"encoding/json"
	"net/http"
)

// this file is used as a helper file.
// it defines error messages that we will output in the project, as well as the structure for our json responses.

// struct of error responses.
// this struct will either be an internal error (program fault) or user error (something incorrect is passed in) -> though this error handling should be handled in the front too.
type Error struct {
	Code    int
	Message string
}

// TotalFields represents the complete request body received from the frontend.
// It includes all fields (common + platform-specific) for all supported platforms.
type TotalFields struct {
	// --- Common fields ---
	Platforms   []string `json:"platforms"`   // e.g. ["youtube", "pinterest", "linkedin"]
	Title       string   `json:"title"`       // shared by YouTube, Pinterest, Reddit
	Description string   `json:"description"` // shared by YouTube, Pinterest
	Caption     string   `json:"caption"`     // Instagram, optional for others
	Tags        []string `json:"tags"`        // YouTube
	CategoryID  string   `json:"category_id"` // YouTube
	MediaFile   string   `json:"media_file"`  // base64 or URL from frontend

	// --- YouTube-specific ---
	PrivacyStatus string `json:"privacy_status"` // "public", "private", or "unlisted"

	// --- Instagram-specific ---
	ImageURL   string `json:"image_url"`   // URL if remote; could reuse media_file
	LocationID string `json:"location_id"` // optional
	UserTags   string `json:"user_tags"`   // optional, comma-separated

	// --- Pinterest-specific ---
	BoardID    string `json:"board_id"`
	Link       string `json:"link"`
	SourceType string `json:"source_type"` // e.g. "image_url"

	// --- Reddit-specific ---
	Subreddit string `json:"subreddit"`
	PostType  string `json:"post_type"` // "self", "link", or "image"
	Text      string `json:"text"`      // for "self" posts
	URL       string `json:"url"`       // for "link" or "image" posts
	Resubmit  bool   `json:"resubmit"`
	NSFW      bool   `json:"nsfw"`

	// --- LinkedIn-specific ---
	Author         string `json:"author"`          // URN of person/org
	LifecycleState string `json:"lifecycle_state"` // usually "PUBLISHED"
	TextLinkedIn   string `json:"text_linkedin"`   // avoids name clash with Reddit "text"
	MediaType      string `json:"media_type"`      // "IMAGE" or "VIDEO"
	MediaStatus    string `json:"media_status"`    // "READY"
	MediaPath      string `json:"media_path"`      // URN or URL
	Visibility     string `json:"visibility"`      // "PUBLIC"
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	// response header (error case)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// json for output (using Encode)
	json.NewEncoder(w).Encode(resp)
}

// wrapper function for writeError. This lets us use it with two different use cases.

var (
	// error is not internal.
	HandleRequestError = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	// internal error. we log it speerately!
	HandleInternalError = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occured", http.StatusInternalServerError)
	}
)
