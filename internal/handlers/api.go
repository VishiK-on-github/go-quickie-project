package handlers

import (
	"goapi/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// global middleware - applies to all endpoints
	// removes slashes at end of the url path
	// to avoid 404s
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		// middleware to check auth for /account
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})

}
