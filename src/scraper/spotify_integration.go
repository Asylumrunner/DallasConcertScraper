package main

import (
	"net/http"
	"strings"
	"encoding/base64"
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
	"regexp"
)

type spotifyAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type spotifySearchArtistsResponse struct {
	Artists struct {
		Href  string `json:"href"`
		Items []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Genres []interface{} `json:"genres"`
			Href   string        `json:"href"`
			ID     string        `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name       string `json:"name"`
			Popularity int    `json:"popularity"`
			Type       string `json:"type"`
			URI        string `json:"uri"`
		} `json:"items"`
		Limit    int         `json:"limit"`
		Next     interface{} `json:"next"`
		Offset   int         `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int         `json:"total"`
	} `json:"artists"`
}

func getSpotifyAuth() string {
	auth_request_body := strings.NewReader("grant_type=client_credentials")
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", auth_request_body)
	auth_payload := []byte(os.Getenv("client_id") + ":" + os.Getenv("client_secret"))
	encoded_auth_payload := base64.StdEncoding.EncodeToString(auth_payload)
	req.Header.Set("Authorization", "Basic " + encoded_auth_payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		log.Print("Retrieving Spotify access token failed!")
		return ""
	}

	body, err := ioutil.ReadAll(res.Body)
	auth_response := consumeAuthResponse(body)
	log.Print(*auth_response)
	return auth_response.AccessToken
}

func consumeAuthResponse(body []byte) *spotifyAuthResponse{
	newSpotifyAuthResponse := spotifyAuthResponse{}
	err := json.Unmarshal(body, &newSpotifyAuthResponse)
	if err != nil {
		log.Print("Couldn't unmarshal json response")
	}

	return &newSpotifyAuthResponse
}

func searchSpotifyForArtist(artist string, spotify_api_cred string) string{
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.spotify.com/v1/search", nil)
	req.Header.Set("Authorization", "Bearer "+ spotify_api_cred)

	query := req.URL.Query()
	query.Add("type", "artist")
	cleanedArtistName := cleanEventName(artist)
	query.Add("q", strings.Replace(cleanedArtistName, " ", "+", -1))
	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	if err != nil {
		log.Print("An error occured while trying to retrieve the artist from Spotify")
		return ""
	}
	body, _ := ioutil.ReadAll(res.Body)
	search_response := consumeSearchResponse(body)
	if len(search_response.Artists.Items) > 0 {
		return search_response.Artists.Items[0].ExternalUrls.Spotify
	} else {
		return ""
	}
	
}

func consumeSearchResponse(body []byte) *spotifySearchArtistsResponse{
	newSpotifySearchResponse := spotifySearchArtistsResponse{}
	err := json.Unmarshal(body, &newSpotifySearchResponse)
	if err != nil {
		log.Print("Couldn't unmarshal JSON response")
	}
	log.Print(newSpotifySearchResponse)
	return &newSpotifySearchResponse
}

func cleanEventName(originalName string) string{
	compiledExpression, _ := regexp.Compile("^[^:]*")
	//Grabs the substring occuring before the first colon, or the whole string if no colon is there
	cleanedName := compiledExpression.Find([]byte(originalName))

	finalizedName := strings.Replace(string(cleanedName), " On Tour", "", -1)
	return finalizedName
}