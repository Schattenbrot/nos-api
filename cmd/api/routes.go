package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func (app *application) chiRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH"},
		MaxAge:         300,
	}))

	r.Get("/", app.findAllWeapons)

	r.Get("/status", app.statusHandler)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/weapons", func(r chi.Router) {
			r.Get("/", app.findAllWeapons)

			r.Post("/", app.createWeapon)
			r.Get("/{id}", app.findOneWeaponById)
			r.Patch("/{id}", app.updateWeaponById)
			r.Delete("/{id}", app.deleteWeaponById)

			r.Route("/profession", func(r chi.Router) {
				r.Get("/{profession}", app.findAllWeaponsByProfession)
			})
		})

		r.Route("/fairies", func(r chi.Router) {
			r.Get("/", app.findAllFairies)

			r.Post("/", app.insertFairy)
			r.Get("/{id}", app.findFairyById)
			r.Patch("/{id}", app.updateFairyById)
			r.Delete("/{id}", app.deleteFairyById)

			r.Route("/element", func(r chi.Router) {
				r.Get("/{element}", app.findAllFairiesByElement)
			})
		})
	})

	return r
}
