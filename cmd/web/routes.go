package main

import (
	"GoTel/pkg/config"
	"GoTel/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	//	mux := pat.New()
	//
	//	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	//	return mux
	//

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	//mux.Use(WriteToConsole)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	//accessing static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
