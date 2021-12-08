package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routes handles all the URL-routing.
func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/weapons", app.findAllWeapons)
	router.HandlerFunc(http.MethodPost, "/v1/weapons", app.createWeapon)
	router.HandlerFunc(http.MethodGet, "/v1/weapons/profession/:profession", app.findAllWeaponsByProfession)
	router.HandlerFunc(http.MethodGet, "/v1/weapons/id/:id", app.findOneWeaponById)
	router.HandlerFunc(http.MethodPatch, "/v1/weapons/:id", app.updateWeaponById)
	router.HandlerFunc(http.MethodDelete, "/v1/weapons/:id", app.deleteWeaponById)

	router.HandlerFunc(http.MethodGet, "/v1/fairies", app.findAllFairies)
	router.HandlerFunc(http.MethodPost, "/v1/fairies", app.insertFairy)
	router.HandlerFunc(http.MethodGet, "/v1/fairies/element/:element", app.findAllFairiesByElement)
	router.HandlerFunc(http.MethodGet, "/v1/fairies/id/:id", app.findFairyById)
	router.HandlerFunc(http.MethodPatch, "/v1/fairies/:id", app.updateFairyById)
	router.HandlerFunc(http.MethodDelete, "/v1/fairies/:id", app.deleteFairyById)

	return router
}
