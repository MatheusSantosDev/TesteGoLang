package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Post struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Body  string `json:"body"`
}

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/comments")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Fatalln(err)
	}

	channel := make(chan Post)

	go func() {
		for i := 0; i < 50; i++ {
			go worker(channel)
		}
	}()

	for _, post := range posts {
		channel <- post
	}

}

func worker(channel chan Post) {
	for p := range channel {
		fmt.Println("From: " + p.Email + "\n" + p.Name + " Say " + p.Body + "\n")
		time.Sleep(time.Second * 1)
	}

}
