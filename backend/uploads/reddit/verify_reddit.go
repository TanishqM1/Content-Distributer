package reddit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// Replace these with your credentials
const (
	clientID     = "eSffeXojzcwY07XiIZMyRA"
	clientSecret = "HjfMyVo9KnCtSwKCi5cbqT-5f5-fPg"
	username     = "Happy_Bookkeeper6491"
	password     = "Social#2510"
)

func UploadReddit() {
	fmt.Print("Reddit!")
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
	req.Header.Set("User-Agent", "SocialContentVerifier/0.1 by "+username)
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

	fmt.Println("🔹 Reddit API response:")
	fmt.Printf("%+v\n", tokenResp)

	// Step 2: Optional test request to verify token
	if accessToken, ok := tokenResp["access_token"].(string); ok {
		req2, _ := http.NewRequest("GET", "https://oauth.reddit.com/api/v1/me", nil)
		req2.Header.Set("Authorization", "bearer "+accessToken)
		req2.Header.Set("User-Agent", "SocialContentVerifier/0.1 by "+username)

		resp2, err := http.DefaultClient.Do(req2)
		if err != nil {
			log.Fatal(err)
		}
		defer resp2.Body.Close()

		var me map[string]interface{}
		json.NewDecoder(resp2.Body).Decode(&me)
		fmt.Println("✅ Authenticated as:")
		fmt.Printf("%+v\n", me)
	}

	post()
}
