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

	router.HandlerFunc(http.MethodGet, "/v1/fairies", app.findAllFairies)
	router.HandlerFunc(http.MethodPost, "/v1/fairies", app.insertFairy)

	return router
}
