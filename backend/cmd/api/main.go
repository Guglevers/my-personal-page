package main

import (
	"net/http"
	handler "my-personal-page/backend/internal/handler"
)

func main() {

	http.HandleFunc("/post", handler.NewPostHandler(http.ResponseWriter, &http.Request))

	http.ListenAndServe(":80", nil)
}