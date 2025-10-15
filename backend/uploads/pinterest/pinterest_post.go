package pinterest

import "fmt"

func UploadPinterest(boardID string, title string, description string, link string, sourceType string, imageURL string) {
	// implement the pinterest API call here, and sync with there parameters!
	// currently only works with web-hosted images.
	fmt.Println("📸 UploadPinterest() called with:")
	fmt.Printf("  Board ID: %s\n", boardID)
	fmt.Printf("  Title: %s\n", title)
	fmt.Printf("  Description: %s\n", description)
	fmt.Printf("  Link: %s\n", link)
	fmt.Printf("  Source Type: %s\n", sourceType)
	fmt.Printf("  Image URL: %s\n", imageURL)
	fmt.Println("------------------------------------")
}
