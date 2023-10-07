package handlers

import (
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joefitzy/golearn/goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
