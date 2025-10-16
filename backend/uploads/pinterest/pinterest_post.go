package pinterest

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// WORKS INDEPENDENTLY, NEED TO HOOKUP W/ FRONTEND AND BACKEND
// need imagepath, title, boardID.

func main() {
	apiURL := "https://api.upload-post.com/api/upload_photos"
	apiKey := "placeholder"
	imagePath := "images.jpg" // your local file (in same folder)
	user := "SocialContentDistributer"
	title := "Testing"
	caption := "This is my post!"
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
	writer.WriteField("pinterest_cover_image_data", "https://media.istockphoto.com/id/517188688/photo/mountain-landscape.jpg?s=612x612&w=0&k=20&c=A63koPKaCyIwQWOTFBRWXj_PwCrR4cEoOw2S9Q7yVl8=")
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
