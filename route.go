package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/iamdevtry/go_with_mux/entity"
	"github.com/iamdevtry/go_with_mux/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPosts(r http.ResponseWriter, req *http.Request) {
	r.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		r.WriteHeader(http.StatusInternalServerError)
		r.Write([]byte(`{"error": "Error getting the posts"}`))
		return
	}

	r.WriteHeader(http.StatusOK)
	json.NewEncoder(r).Encode(posts)
}

func addPost(r http.ResponseWriter, req *http.Request) {
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		r.WriteHeader(http.StatusInternalServerError)
		r.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}

	post.Id = rand.Int()
	repo.Save(&post)

	r.WriteHeader(http.StatusOK)
	json.NewEncoder(r).Encode(post)
}
