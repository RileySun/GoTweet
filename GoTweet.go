package main

import (
	"fmt"
	"time"
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"

//	"net/http/httputil"
)

//Global
var bearer string

//Structs
type UserIDReq struct {
	Data []UserIDData `json:"data"`
}
type UserIDData struct {
	ID 					string `json:"id"`
	ProfileImageURL 	string `json:"profile_image_url"`
	Username 			string `json:"username"`
	Name 				string `json:"name"`
}

type UserTweetReq struct {
	Data []UserTweetData `json:"data"`
	Meta UserTweetMeta `json:"meta"`
}
type UserTweetData struct {
	ID string `json:"id"`
	EditHistoryTweetIDs []string `json:"edit_history_tweet_ids"`
	Text string `json:"text"`
	CreatedAt string `json:"created_at"`
}
type UserTweetMeta struct {
	ResultCount int `json:"result_count"`
	OldestID string `json:"oldest_id"`
	NewestID string `json:"newest_id"`
}

//Main
func main() {
	bearer = "INSERT BEARER TOKEN HERE"
}

//Util
func twitterAPI(url string) []byte {
	//Generic GET request to Twitter API
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	
	req, err := http.NewRequest("GET", url, nil)
	
	if err != nil {
		log.Fatal(err)
	}
	
	req.Header.Set("Authorization", "Bearer " + bearer)
	
	resp, err := client.Do(req)
	
	if err != nil {
		log.Fatal(err)
	}
	
	responseData, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		log.Fatal(err)
	}
	
	defer resp.Body.Close()
	
	return responseData
}

func getUserID(user string) string {
	//Returns Twitter ID of User based of Username
	url := "https://api.twitter.com/2/users/by?usernames=" + user + "&user.fields=profile_image_url"
	responseData := twitterAPI(url)
   
	data := UserIDReq{}
	err := json.Unmarshal(responseData, &data)
	
	if err != nil {
		 fmt.Println(err)
	}
	
	return data.Data[0].ID
}

func getUserTweets(user string) UserTweetReq {
	userID := getUserID(user)

	url := "https://api.twitter.com/2/users/" + userID +"/tweets?&tweet.fields=created_at&expansions=attachments.media_keys&media.fields=preview_image_url,url"
	responseData := twitterAPI(url)
	
	data := UserTweetReq{}
	err := json.Unmarshal(responseData, &data)
	
	if err != nil {
		 fmt.Println(err)
	}
	
	return data
}