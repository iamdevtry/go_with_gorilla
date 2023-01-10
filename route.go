package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var posts []Post

func init() {
	posts = []Post{
		{Id: 1, Title: "Hello World", Text: "This is a sample post"},
		{Id: 2, Title: "Hello World 2", Text: "This is a sample post 2"},
		{Id: 3, Title: "Hello World 3", Text: "This is a sample post 3"},
	}
}

func getPosts(r http.ResponseWriter, req *http.Request) {
	r.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		r.WriteHeader(http.StatusInternalServerError)
		r.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}
	r.WriteHeader(http.StatusOK)
	r.Write(result)
}

func addPost(r http.ResponseWriter, req *http.Request) {
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		r.WriteHeader(http.StatusInternalServerError)
		r.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}

	post.Id = len(posts) + 1
	// post.Title = "Hello World " + string(rune(post.Id))
	// post.Text = "This is a sample post " + string(rune(post.Id))
	posts = append(posts, post)

	r.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	if err != nil {
		r.WriteHeader(http.StatusBadRequest)
		r.Write([]byte(`{"error": "Error marshalling the post"}`))
	}

	r.Write(result)

}
