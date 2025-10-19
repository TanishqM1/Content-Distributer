package instagram

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// WORKS INDEPENDENTLY, NEED TO HOOKUP W/ FRONTEND AND BACKEND

func UploadInstagram(mediaPath string, title string, userTags string) {

	fmt.Printf("\n Retrieved the following fields: \n MediaPath: %v \n Title: %v  \n userTags: %v", mediaPath, title, userTags)

	apiURL := "https://api.upload-post.com/api/upload"
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Cannot Load .ENV (UploadPinterest())")
	}
	apiKey := os.Getenv("UploadsAPI")
	fmt.Println(apiKey)
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
