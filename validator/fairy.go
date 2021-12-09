package validator

import (
	"errors"

	"github.com/Schattenbrot/nos-api/models"
)

func FairyValidation(fairy models.Fairy) error {
	if fairy.Level != nil {
		if fairy.Level.Min < 1 || fairy.Level.Min > 99 {
			return errors.New("min level must be between 1 and 99")
		}
		if fairy.Level.Max < 1 || fairy.Level.Max > 99 {
			return errors.New("max level must be between 1 and 99")
		}
		if fairy.Level.Min > fairy.Level.Max {
			return errors.New("min level must be higher than max level")
		}
	}
	if fairy.Name == "" {
		return errors.New("name must not be empty")
	}
	if fairy.Element != "fire" && fairy.Element != "water" && fairy.Element != "light" && fairy.Element != "shadow" {
		return errors.New("element must be either fire, water, light, or shadow")
	}

	return nil
}
