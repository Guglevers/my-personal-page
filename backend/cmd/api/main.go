package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"my-personal-page/backend/internal/db"
	postsDB "my-personal-page/backend/internal/db/generated"
	handler "my-personal-page/backend/internal/handler"
	repository "my-personal-page/backend/internal/repository"
	service "my-personal-page/backend/internal/service"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

func run() (*postsDB.Queries, error) {
	ctx := context.Background()

	dbConn, err := sql.Open("sqlite3", ":memory:?_loc=auto")
	if err != nil {
		return nil, err
	}

	if _, err := dbConn.ExecContext(ctx, db.DDL); err != nil {
		return nil, err
	}

	return postsDB.New(dbConn), nil
}

func main() {
	queries, err := run()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewPostRepo(queries)
	postService := service.NewPostService(repo)
	postHandler := handler.NewPostHandler(postService)

	mux := chi.NewMux()

	mux.Post("/api/posts", postHandler.Create)
	mux.Get("/api/posts", postHandler.GetAll)
	mux.Get("/api/posts/{id}", postHandler.Get)
	mux.Delete("/api/posts/{id}", postHandler.Delete)
	mux.Put("/api/posts", postHandler.Update)

	log.Println("Iniciando servidor na porta 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
