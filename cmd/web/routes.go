package main

import (
	"net/http"

	"github.com/ArtyomHov/go-course/pkg/config"
	"github.com/ArtyomHov/go-course/pkg/handlers"
	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
