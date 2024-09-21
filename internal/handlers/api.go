package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/marcelowf/First-Go-API/internal/middleware"
)

func Handler(router *chi.Mux) {
	//Middleware Global
	router.Use(chimiddle.StripSlashes)

	router.Route("/account", func(r chi.Router){
		//Middleware para rota /account 
		r.Use(middleware.Authorization)
		r.Get("/coins", GetCoinBalance)
	})
}
