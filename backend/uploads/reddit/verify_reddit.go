package reddit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	clientID     = "eSffeXojzcwY07XiIZMyRA"
	clientSecret = "HjfMyVo9KnCtSwKCi5cbqT-5f5-fPg"
	username     = "Happy_Bookkeeper6491"
	password     = "Social#2510"
)

func UploadReddit(subreddit, postType, title, text, link string, resubmit, nsfw bool) {
	fmt.Println("📢 UploadReddit() called with:")
	fmt.Printf("  Subreddit: %s\n", subreddit)
	fmt.Printf("  Post Type: %s\n", postType)
	fmt.Printf("  Title: %s\n", title)
	fmt.Printf("  Text: %s\n", text)
	fmt.Printf("  URL: %s\n", link)
	fmt.Printf("  Resubmit: %v\n", resubmit)
	fmt.Printf("  NSFW: %v\n", nsfw)
	fmt.Println("------------------------------------")

	// Step 1: Get access token
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", username)
	data.Set("password", password)
	data.Set("scope", "identity submit")

	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Set("User-Agent", "windows:SocialContentDistributer:v1.0 (by /u/"+username+")")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var tokenResp map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		log.Fatal(err)
	}

	token, ok := tokenResp["access_token"].(string)
	if !ok {
		log.Fatalf("❌ Could not get access_token: %+v\n", tokenResp)
	}

	fmt.Println("✅ Access token received!")

	// Step 2: Verify token
	req2, _ := http.NewRequest("GET", "https://oauth.reddit.com/api/v1/me", nil)
	req2.Header.Set("Authorization", "bearer "+token)
	req2.Header.Set("User-Agent", "windows:SocialContentDistributer:v1.0 (by /u/"+username+")")

	resp2, err := http.DefaultClient.Do(req2)
	if err != nil {
		log.Fatal(err)
	}
	defer resp2.Body.Close()

	var me map[string]interface{}
	json.NewDecoder(resp2.Body).Decode(&me)
	fmt.Println("👤 Authenticated as:", me["name"])

	// Step 3: Create the post
	post(token, subreddit, postType, title, text, link, resubmit, nsfw)
}

func post(accessToken, subreddit, postType, title, text, link string, resubmit, nsfw bool) {
	data := url.Values{}
	data.Set("sr", subreddit)
	data.Set("kind", postType)
	data.Set("title", title)
	data.Set("resubmit", fmt.Sprintf("%v", resubmit))
	data.Set("nsfw", fmt.Sprintf("%v", nsfw))
	data.Set("sendreplies", "true")

	if postType == "self" {
		data.Set("text", text)
	} else if postType == "link" || postType == "image" {
		data.Set("url", link)
	}

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
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Printf("🧩 Reddit Response: %+v\n", res)
}
