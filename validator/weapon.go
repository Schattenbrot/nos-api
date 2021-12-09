package validator

import (
	"errors"

	"github.com/Schattenbrot/nos-api/models"
)

func WeaponValidation(weapon models.Weapon) error {
	if weapon.Level < 1 || weapon.Level > 99 {
		return errors.New("level must be between 1 and 99")
	}

	if weapon.Profession != "adventurer" && weapon.Profession != "mage" && weapon.Profession != "bowman" && weapon.Profession != "swordsman" {
		return errors.New("profession must be either adventurer, mage, bowman, or swordsman")
	}

	if weapon.Name == "" {
		return errors.New("name must not be empty")
	}

	if weapon.Damage != nil {
		if weapon.Damage.Min < 1 {
			return errors.New("min damage must be greater than zero")
		}
		if weapon.Damage.Max < 1 {
			return errors.New("max damage must be greater than zero")
		}
		if weapon.Damage.Max < weapon.Damage.Min {
			return errors.New("max damage must be greater than min damage")
		}
	}

	if weapon.Physical != nil {
		if weapon.Physical.HitRate < 1 {
			return errors.New("hit rate must be greater than zero")
		}
		if weapon.Physical.CritChance < 0 {
			return errors.New("crit chance must be greater or equal zero")
		}
		if weapon.Physical.Crit < 0 {
			return errors.New("crit must be greater or equal zero")
		}
	}

	if weapon.Concentration < 0 {
		return errors.New("concentration must be greater than zero")
	}

	return nil
}
