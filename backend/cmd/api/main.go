package main

import (
	"log"
	handler "my-personal-page/backend/internal/handler"
	repository "my-personal-page/backend/internal/repository"
	service "my-personal-page/backend/internal/service"
	"net/http"
)

func main() {

	repo        := repository.NewPostMemoryRepo()
	postService := service.NewPostService(repo)
	postHandler := handler.NewPostHandler(postService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/posts", postHandler.CreatePost)
	mux.HandleFunc("GET /api/posts", postHandler.GetPosts)

	log.Println("Iniciando servidor na porta 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil{
		log.Fatal(err)
	}
}