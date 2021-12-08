package main

import (
	"github.com/Schattenbrot/nos-api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func chiRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH"},
		MaxAge:         300,
	}))

	r.Get("/", handlers.FindAllWeapons)

	r.Get("/status", handlers.StatusHandler)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/weapons", func(r chi.Router) {
			r.Get("/", handlers.FindAllWeapons)

			r.Post("/", handlers.InsertWeapon)
			r.Get("/{id}", handlers.FindOneWeaponById)
			r.Patch("/{id}", handlers.UpdateWeaponById)
			r.Delete("/{id}", handlers.DeleteWeaponById)

			r.Route("/profession", func(r chi.Router) {
				r.Get("/{profession}", handlers.FindAllWeaponsByProfession)
			})
		})

		r.Route("/fairies", func(r chi.Router) {
			r.Get("/", handlers.FindAllFairies)

			r.Post("/", handlers.InsertFairy)
			r.Get("/{id}", handlers.FindFairyById)
			r.Patch("/{id}", handlers.UpdateFairyById)
			r.Delete("/{id}", handlers.DeleteFairyById)

			r.Route("/element", func(r chi.Router) {
				r.Get("/{element}", handlers.FindAllFairiesByElement)
			})
		})
	})

	return r
}
