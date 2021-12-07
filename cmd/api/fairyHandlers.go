package main

import (
	"encoding/json"
	"net/http"

	"github.com/Schattenbrot/nos-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *application) insertFairy(w http.ResponseWriter, r *http.Request) {
	var fairy models.Fairy

	err := json.NewDecoder(r.Body).Decode(&fairy)
	if err != nil {
		app.errorJSON(w, err)
	}

	type jsonResp struct {
		OK bool                `json:"ok"`
		ID *primitive.ObjectID `json:"_id"`
	}

	id, err := app.models.DB.InsertFairy(fairy)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	ok := jsonResp{OK: true, ID: id}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

// findAllFairies is the handler for the FindAllFairies method.
func (app *application) findAllFairies(w http.ResponseWriter, r *http.Request) {
	fairies, err := app.models.DB.FindAllFairies()
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, fairies, "fairies")
	if err != nil {
		app.logger.Println(err)
	}
}
