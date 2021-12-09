package handlers

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"

	"github.com/Schattenbrot/nos-api/config"
	"github.com/Schattenbrot/nos-api/models"
	"github.com/Schattenbrot/nos-api/validator"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertFairy(w http.ResponseWriter, r *http.Request) {
	var fairy models.Fairy

	err := json.NewDecoder(r.Body).Decode(&fairy)
	if err != nil {
		errorJSON(w, err)
		return
	}

	// Validation
	err = validator.FairyValidation(fairy)
	if err != nil {
		errorJSON(w, err)
		return
	}
	// Escape strings
	for i := 0; i < len(fairy.Effects); i++ {
		fairy.Effects[i] = template.HTMLEscapeString(fairy.Effects[i])
	}
	for i := 0; i < len(fairy.HowToGet); i++ {
		fairy.HowToGet[i] = template.HTMLEscapeString(fairy.HowToGet[i])
	}

	type jsonResp struct {
		OK bool                `json:"ok"`
		ID *primitive.ObjectID `json:"_id"`
	}

	id, err := config.App.Models.DB.InsertFairy(fairy)
	if err != nil {
		errorJSON(w, err)
		return
	}

	ok := jsonResp{OK: true, ID: id}

	err = writeJSON(w, http.StatusCreated, ok, "response")
	if err != nil {
		errorJSON(w, err)
		return
	}
}

// findAllFairies is the handler for the FindAllFairies method.
func FindAllFairies(w http.ResponseWriter, r *http.Request) {
	fairies, err := config.App.Models.DB.FindAllFairies()
	if err != nil {
		config.App.Logger.Println(err)
	}

	if fairies == nil {
		err = errors.New("the items do not exist")
		errorJSON(w, err, http.StatusNotFound)
		config.App.Logger.Println(err)
		return
	}

	err = writeJSON(w, http.StatusOK, fairies, "fairies")
	if err != nil {
		config.App.Logger.Println(err)
	}
}

func FindAllFairiesByElement(w http.ResponseWriter, r *http.Request) {
	element := chi.URLParam(r, "element")
	if element != "fire" && element != "water" && element != "light" && element != "shadow" {
		err := errors.New("invalid element parameter")
		errorJSON(w, err, http.StatusBadRequest)
		config.App.Logger.Println(err)
		return
	}

	fairies, err := config.App.Models.DB.FindAllFairiesByElement(element)
	if err != nil {
		config.App.Logger.Println(err)
	}

	if fairies == nil {
		err = errors.New("the items do not exist")
		errorJSON(w, err, http.StatusNotFound)
		config.App.Logger.Println(err)
		return
	}

	err = writeJSON(w, http.StatusOK, fairies, "fairies")
	if err != nil {
		config.App.Logger.Println(err)
	}
}

func FindFairyById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		config.App.Logger.Println(errors.New("invalid id parameter"))
		errorJSON(w, err)
		return
	}

	fairy, err := config.App.Models.DB.FindFairyById(id)
	if err != nil {
		errorJSON(w, err)
		config.App.Logger.Println(err)
		return
	}

	err = writeJSON(w, http.StatusOK, fairy, "fairy")
	if err != nil {
		config.App.Logger.Println(err)
	}
}

func UpdateFairyById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		config.App.Logger.Println(errors.New("invalid id parameter"))
		errorJSON(w, err)
		return
	}

	var updateFairy models.Fairy
	err = json.NewDecoder(r.Body).Decode(&updateFairy)
	if err != nil {
		errorJSON(w, err)
		return
	}

	result, err := config.App.Models.DB.UpdateFairyById(id, updateFairy)
	if err != nil {
		if err.Error() == "not found" {
			errorJSON(w, err, http.StatusNotFound)
			return
		}
		errorJSON(w, err)
		return
	}

	err = writeJSON(w, http.StatusNoContent, result, "updated")
	if err != nil {
		config.App.Logger.Println(err)
	}
}

func DeleteFairyById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		config.App.Logger.Println(errors.New("invalid id parameter"))
		errorJSON(w, err)
		return
	}

	result, err := config.App.Models.DB.DeleteFairyById(id)
	if err != nil {
		if err.Error() == "not found" {
			errorJSON(w, err, http.StatusNotFound)
			return
		}
		errorJSON(w, err)
		return
	}

	err = writeJSON(w, http.StatusNoContent, result, "deleted")
	if err != nil {
		config.App.Logger.Println(err)
	}
}
