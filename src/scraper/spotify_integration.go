package main

import (
	"net/http"
	"strings"
	"encoding/base64"
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
)

type spotifyAuthResponse struct {
	access_token string
	token_type string
	expires_in int
}

func getSpotifyAuth() string {
	auth_request_body := strings.NewReader("grant_type=client_credentials")
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", auth_request_body)
	auth_payload := []byte(os.Getenv("client_id") + ":" + os.Getenv("client_secret"))
	encoded_auth_payload := base64.StdEncoding.EncodeToString(auth_payload)
	log.Print("Encoded auth payload: " + encoded_auth_payload)
	req.Header.Set("Authorization", "Basic " + encoded_auth_payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		log.Print("Retrieving Spotify access token failed!")
		return ""
	}

	body, err := ioutil.ReadAll(res.Body)
	auth_response := consumeAuthResponse(body)
	return auth_response.access_token
}

func consumeAuthResponse(body []byte) *spotifyAuthResponse{
	raw_response := string(body)
	log.Print("Raw Spotify Response: " + raw_response)
	newSpotifyAuthResponse := new(spotifyAuthResponse)
	err := json.Unmarshal(body, &newSpotifyAuthResponse)
	if err != nil {
		log.Print("Couldn't unmarshal json response")
	}

	return newSpotifyAuthResponse
}
/* func searchSpotifyForArtist (artist string) bool{
	
} */