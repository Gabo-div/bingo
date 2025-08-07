package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Gabo-div/bingo/apps/backend-main/internal/echo"
	"github.com/Gabo-div/bingo/apps/backend-main/internal/user"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	echo.Register(r)
	user.Register(r)

	log.Printf("Running on localhost:3001")

	http.ListenAndServe(
		"localhost:3001",
		r,
	)
}
