package main

import (
	"encoding/json"
	"net/http"

	"github.com/Schattenbrot/nos-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// createWeapon is the handler for the InsertWeapon method.
func (app *application) createWeapon(w http.ResponseWriter, r *http.Request) {
	var weapon models.Weapon

	err := json.NewDecoder(r.Body).Decode(&weapon)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	type jsonResp struct {
		OK bool                `json:"ok"`
		ID *primitive.ObjectID `json:"_id"`
	}

	id, err := app.models.DB.InsertWeapon(weapon)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	ok := jsonResp{
		OK: true,
		ID: id,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

// findAllWeapons is the handler for the FindAllWeapons method.
func (app *application) findAllWeapons(w http.ResponseWriter, r *http.Request) {
	weapons, err := app.models.DB.FindAllWeapons()
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, weapons, "weapons")
	if err != nil {
		app.logger.Println(err)
	}
}
