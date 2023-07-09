package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// send the GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP GET request failed: ", err)
		return
	}
	// Ensure that the response body is closed after we're done with it
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body: ", err)
		return
	}
	// Parse the JSON response into a Post struct
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println("Error parsing the JSON response: ", err)
		return
	}

	// Print the parsed post
	fmt.Printf("User Id: %d\n", post.UserId)
	fmt.Printf("Id: %d\n", post.Id)
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("Body: %s\n", post.Body)
}
