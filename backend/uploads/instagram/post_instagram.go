package instagram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Instagram Graph API configuration
const (
	BaseURL = "https://graph.facebook.com/v18.0"
)

// Instagram API response structures
type MediaContainerResponse struct {
	ID string `json:"id"`
}

type PublishResponse struct {
	ID string `json:"id"`
}

type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Code    int    `json:"code"`
	} `json:"error"`
}

func UploadInstagram(imageURL *string, caption *string, locationID *string, userTags *string) {
	fmt.Printf("\n UploadInstagram() function")

	// Validate required parameters
	if *imageURL == "" {
		log.Fatalf("You must provide an image URL for Instagram upload")
	}

	if *caption == "" {
		log.Fatalf("You must provide a caption for Instagram upload")
	}

	// Get credentials from environment variables
	pageAccessToken := os.Getenv("INSTAGRAM_PAGE_ACCESS_TOKEN")
	instagramAccountID := os.Getenv("INSTAGRAM_ACCOUNT_ID")

	if pageAccessToken == "" {
		log.Fatalf("INSTAGRAM_PAGE_ACCESS_TOKEN environment variable not set")
	}

	if instagramAccountID == "" {
		log.Fatalf("INSTAGRAM_ACCOUNT_ID environment variable not set")
	}

	// Log the parameters being used
	fmt.Printf("Image URL: %s\n", *imageURL)
	fmt.Printf("Caption: %s\n", *caption)
	fmt.Printf("Location ID: %s\n", *locationID)
	fmt.Printf("User Tags: %s\n", *userTags)
	fmt.Printf("Instagram Account ID: %s\n", instagramAccountID)

	// Step 1: Create media container
	containerID, err := createMediaContainer(instagramAccountID, pageAccessToken, *imageURL, *caption, *locationID, *userTags)
	if err != nil {
		log.Fatalf("Failed to create media container: %v", err)
	}

	fmt.Printf("Media container created with ID: %s\n", containerID)

	// Step 2: Wait for container to be ready (optional - you can implement polling)
	// For now, we'll proceed directly to publish

	// Step 3: Publish the media
	mediaID, err := publishMedia(instagramAccountID, pageAccessToken, containerID)
	if err != nil {
		log.Fatalf("Failed to publish media: %v", err)
	}

	fmt.Printf("Instagram upload successful! Media ID: %s\n", mediaID)
}

func createMediaContainer(accountID, accessToken, imageURL, caption, locationID, userTags string) (string, error) {
	url := fmt.Sprintf("%s/%s/media", BaseURL, accountID)

	// Build the request body
	body := map[string]string{
		"image_url": imageURL,
		"caption":   caption,
	}

	// Add optional parameters
	if locationID != "" {
		body["location_id"] = locationID
	}

	if userTags != "" {
		body["user_tags"] = userTags
	}

	// Add access token
	body["access_token"] = accessToken

	// Convert to JSON
	jsonData, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	// Make the request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Check for errors
	if resp.StatusCode != http.StatusOK {
		var errorResp ErrorResponse
		json.Unmarshal(responseBody, &errorResp)
		return "", fmt.Errorf("API error: %s", errorResp.Error.Message)
	}

	// Parse success response
	var containerResp MediaContainerResponse
	err = json.Unmarshal(responseBody, &containerResp)
	if err != nil {
		return "", err
	}

	return containerResp.ID, nil
}

func publishMedia(accountID, accessToken, containerID string) (string, error) {
	url := fmt.Sprintf("%s/%s/media_publish", BaseURL, accountID)

	// Build the request body
	body := map[string]string{
		"creation_id":  containerID,
		"access_token": accessToken,
	}

	// Convert to JSON
	jsonData, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	// Make the request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Check for errors
	if resp.StatusCode != http.StatusOK {
		var errorResp ErrorResponse
		json.Unmarshal(responseBody, &errorResp)
		return "", fmt.Errorf("API error: %s", errorResp.Error.Message)
	}

	// Parse success response
	var publishResp PublishResponse
	err = json.Unmarshal(responseBody, &publishResp)
	if err != nil {
		return "", err
	}

	return publishResp.ID, nil
}
