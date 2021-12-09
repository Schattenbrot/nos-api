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

// createWeapon is the handler for the InsertWeapon method.
func InsertWeapon(w http.ResponseWriter, r *http.Request) {
	var weapon models.Weapon

	err := json.NewDecoder(r.Body).Decode(&weapon)
	if err != nil {
		errorJSON(w, err)
		return
	}

	// Validation
	// Validation
	err = validator.WeaponValidation(weapon)
	if err != nil {
		errorJSON(w, err)
		return
	}
	// Escape strings
	for i := 0; i < len(weapon.Effects); i++ {
		weapon.Effects[i] = template.HTMLEscapeString(weapon.Effects[i])
	}
	for i := 0; i < len(weapon.HowToGet); i++ {
		weapon.HowToGet[i] = template.HTMLEscapeString(weapon.HowToGet[i])
	}

	type jsonResp struct {
		OK bool                `json:"ok"`
		ID *primitive.ObjectID `json:"_id"`
	}

	id, err := config.App.Models.DB.InsertWeapon(weapon)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	ok := jsonResp{
		OK: true,
		ID: id,
	}

	err = writeJSON(w, http.StatusCreated, ok, "response")
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}
}

// findAllWeapons is the handler for the FindAllWeapons method.
func FindAllWeapons(w http.ResponseWriter, r *http.Request) {
	weapons, err := config.App.Models.DB.FindAllWeapons()
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		config.App.Logger.Println(err)
	}

	if weapons == nil {
		err = errors.New("the items do not exist")
		errorJSON(w, err, http.StatusNotFound)
		config.App.Logger.Println(err)
		return
	}

	err = writeJSON(w, http.StatusOK, weapons, "weapons")
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		config.App.Logger.Println(err)
	}
}

func FindAllWeaponsByProfession(w http.ResponseWriter, r *http.Request) {
	profession := chi.URLParam(r, "profession")
	if profession != "adventurer" && profession != "mage" && profession != "bowman" && profession != "swordsman" && profession != "martial-artist" {
		err := errors.New("invalid profession parameter")
		errorJSON(w, err, http.StatusBadRequest)
		config.App.Logger.Println(err)
		return
	}

	weapons, err := config.App.Models.DB.FindAllWeaponsByProfession(profession)
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		config.App.Logger.Println(err)
	}

	if weapons == nil {
		err = errors.New("the items do not exist")
		errorJSON(w, err, http.StatusNotFound)
		config.App.Logger.Println(err)
		return
	}

	err = writeJSON(w, http.StatusOK, weapons, "weapons")
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		config.App.Logger.Println(err)
	}
}

func FindOneWeaponById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		config.App.Logger.Println(errors.New("invalid id parameter"))
		errorJSON(w, err)
		return
	}

	weapon, err := config.App.Models.DB.FindOneWeaponById(id)
	if err != nil {
		errorJSON(w, err, http.StatusNotFound)
		config.App.Logger.Println(err)
		return
	}

	err = writeJSON(w, http.StatusOK, weapon, "weapon")
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		config.App.Logger.Println(err)
	}
}

func UpdateWeaponById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		config.App.Logger.Println(errors.New("invalid id parameter"))
		errorJSON(w, err)
		return
	}

	var updateWeapon models.Weapon
	err = json.NewDecoder(r.Body).Decode(&updateWeapon)
	if err != nil {
		errorJSON(w, err)
		return
	}

	result, err := config.App.Models.DB.UpdateOneWeaponById(id, updateWeapon)
	if err != nil {
		if err.Error() == "not found" {
			errorJSON(w, err, http.StatusNotFound)
			return
		}
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = writeJSON(w, http.StatusNoContent, result, "updated")
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		config.App.Logger.Println(err)
	}
}

func DeleteWeaponById(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		config.App.Logger.Println(errors.New("invalid id parameter"))
		errorJSON(w, err)
		return
	}

	result, err := config.App.Models.DB.DeleteWeaponById(id)
	if err != nil {
		if err.Error() == "not found" {
			errorJSON(w, err, http.StatusNotFound)
			return
		}
		errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = writeJSON(w, http.StatusNoContent, result, "deleted")
	if err != nil {
		errorJSON(w, err, http.StatusInternalServerError)
		config.App.Logger.Println(err)
	}
}
