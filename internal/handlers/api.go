package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	chimiddle "github.com/go-chi/chi/v5/middleware"
	"github.com/shanto-323/Go-Api/internal/middlewere"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/acount", func(router chi.Router) {

    //IN My Local File
    router.Use(middlewere.Auth)
    router.Get("/coins", GetCoinBalance)
   	})
}



