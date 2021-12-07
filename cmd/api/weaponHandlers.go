package main

import (
	"encoding/json"
	"net/http"

	"github.com/Schattenbrot/nos-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Damage is the type for weapon damage
type WeaponDamagePayload struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}

// WeaponPhysical is the type for physical weapon stat
type WeaponPhysicalPayload struct {
	HitRate    int `json:"hitRate,omitempty"`
	CritChance int `json:"critChance,omitempty"`
	Crit       int `json:"crit,omitempty"`
}

// Weapon is the type for weapons
type WeaponPayload struct {
	ID            primitive.ObjectID    `json:"_id,omitempty"`
	Level         int                   `json:"level,omitempty"`
	Name          string                `json:"name,omitempty"`
	Image         string                `json:"image,omitempty"`
	Damage        WeaponDamagePayload   `json:"damage,omitempty"`
	Physical      WeaponPhysicalPayload `json:"physical,omitempty"`
	Concentration int                   `json:"concentration,omitempty"`
	Effects       []string              `json:"effects,omitempty"`
	HowToGet      []string              `json:"howToGet,omitempty"`
}

// createWeapon is the handler for the InsertWeapon method.
func (app *application) createWeapon(w http.ResponseWriter, r *http.Request) {
	var payload WeaponPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var weaponDamage models.WeaponDamage
	weaponDamage.Max = payload.Damage.Max
	weaponDamage.Min = payload.Damage.Min

	var weaponPhysical models.WeaponPhysical
	weaponPhysical.HitRate = payload.Physical.HitRate
	weaponPhysical.CritChance = payload.Physical.CritChance
	weaponPhysical.Crit = payload.Physical.Crit

	var weapon models.Weapon
	weapon.Level = payload.Level
	weapon.Name = payload.Name
	weapon.Image = payload.Image
	weapon.Damage = &weaponDamage
	weapon.Physical = &weaponPhysical
	weapon.Concentration = payload.Concentration
	weapon.Effects = payload.Effects
	weapon.HowToGet = payload.HowToGet

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
