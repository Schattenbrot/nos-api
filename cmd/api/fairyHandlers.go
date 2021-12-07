package main

import (
	"encoding/json"
	"net/http"

	"github.com/Schattenbrot/nos-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FairyLevel is the type for possible fairy levels
type FairyLevelPayload struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}

// Fairy is the type for fairies
type FairyPayload struct {
	ID       primitive.ObjectID `json:"_id,omitempty"`
	Level    FairyLevelPayload  `json:"level,omitempty"`
	Name     string             `json:"name,omitempty"`
	Element  string             `json:"element,omitempty"`
	Effects  []string           `json:"effects,omitempty"`
	HowToGet []string           `json:"howToGet,omitempty"`
}

func (app *application) insertFairy(w http.ResponseWriter, r *http.Request) {
	var payload FairyPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorJSON(w, err)
	}

	var fairyLevel models.FairyLevel
	fairyLevel.Max = payload.Level.Max
	fairyLevel.Min = payload.Level.Min

	var fairy models.Fairy
	fairy.Level = &fairyLevel
	fairy.Name = payload.Name
	fairy.Element = payload.Element
	fairy.Effects = payload.Effects
	fairy.HowToGet = payload.HowToGet

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
