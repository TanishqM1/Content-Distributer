package instagram

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// WORKS INDEPENDENTLY, NEED TO HOOKUP W/ FRONTEND AND BACKEND

func main() {
	apiURL := "https://api.upload-post.com/api/upload"
	apiKey := "placeholder"
	mediaPath := "images.jpg" // or "myvideo.mp4" — must be in same folder
	title := "My Instagram Post"
	user := "SocialContentDistributer"

	// === Create multipart form ===
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Attach the media file (image or video)
	file, err := os.Open(mediaPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("video", mediaPath)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}

	// Add form fields
	writer.WriteField("title", title)
	writer.WriteField("user", user)
	writer.WriteField("platform[]", "instagram") // ✅ target platform

	writer.Close()

	// === Build HTTP request ===
	req, err := http.NewRequest("POST", apiURL, body)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", "Apikey "+apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// === Send request ===
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status:", resp.Status)
	fmt.Println("Response:", string(respBody))
}
