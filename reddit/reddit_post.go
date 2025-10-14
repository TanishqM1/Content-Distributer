package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Replace these with your credentials
const (
	accessToken = "YOUR_ACCESS_TOKEN" // from your verification step
)

func post() {
	// Create POST form data
	data := url.Values{}
	data.Set("sr", "test")                        // subreddit name
	data.Set("kind", "self")                      // text post
	data.Set("title", "Hello Reddit from Go!")    // post title
	data.Set("text", "Testing Reddit API upload") // post body
	data.Set("resubmit", "true")
	data.Set("sendreplies", "true")

	req, err := http.NewRequest("POST", "https://oauth.reddit.com/api/submit", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "bearer "+accessToken)
	req.Header.Set("User-Agent", "windows:SocialContentDistributer:v1.0 (by /u/Happy_Bookkeeper6491)")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("🔹 Status:", resp.Status)
	resp.Write(os.Stdout)
}
