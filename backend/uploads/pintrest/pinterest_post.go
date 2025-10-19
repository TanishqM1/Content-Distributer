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
// need imagepath, title, boardID. compatible with jpg as of now.

func UploadPinterest(title string, caption string, imagePath string, sourceType string, imageURL string) {

	fmt.Printf("\n Retrieved the following fields: \n Title: %v \n Caption: %v \n imagePath: %v \n SourceType: %v \n ImageURL: %v ", title, caption, imagePath, sourceType, imageURL)

	apiURL := "https://api.upload-post.com/api/upload_photos"
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Cannot Load .ENV (UploadPinterest())")
	}
	apiKey := os.Getenv("UploadsAPI")
	fmt.Println(apiKey)
	user := "SocialContentDistributer"
	boardID := "1126462994236750396"

	// === Create multipart body ===
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Attach image file
	file, err := os.Open(imagePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("photos[]", imagePath)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}

	// === Required Pinterest fields ===
	writer.WriteField("user", user)
	writer.WriteField("title", caption)
	writer.WriteField("platform[]", "pinterest")
	writer.WriteField("async_upload", "false")

	// Pinterest-specific metadata
	writer.WriteField("pinterest_board_id", boardID)
	writer.WriteField("pinterest_title", title)
	writer.WriteField("pinterest_cover_image_content_type", "image/jpeg")
	// writer.WriteField("pinterest_cover_image_data", imagePath)
	writer.WriteField("pinterest_cover_image_key_frame_time", "0e5171886886126120206342978")

	writer.Close()

	// === Build and send request ===
	req, err := http.NewRequest("POST", apiURL, body)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", "Apikey "+apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("Status:", resp.Status)
	fmt.Println("Response:", string(respBody))
}
