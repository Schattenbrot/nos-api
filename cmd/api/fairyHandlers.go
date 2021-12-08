package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Schattenbrot/nos-api/models"
	"github.com/go-chi/chi"
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

func (app *application) findAllFairiesByElement(w http.ResponseWriter, r *http.Request) {
	element := chi.URLParam(r, "element")

	fairies, err := app.models.DB.FindAllFairiesByElement(element)
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, fairies, "fairies")
	if err != nil {
		app.logger.Println(err)
	}
}

func (app *application) findFairyById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	fairy, err := app.models.DB.FindFairyById(id)
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, fairy, "fairy")
	if err != nil {
		app.logger.Println(err)
	}
}

func (app *application) updateFairyById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	var updateFairy models.Fairy
	err = json.NewDecoder(r.Body).Decode(&updateFairy)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	result, err := app.models.DB.UpdateFairyById(id, updateFairy)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, result, "updated")
	if err != nil {
		app.logger.Println(err)
	}
}

func (app *application) deleteFairyById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	result, err := app.models.DB.DeleteFairyById(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, result, "deleted")
	if err != nil {
		app.logger.Println(err)
	}
}
