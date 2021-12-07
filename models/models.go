package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// DBModel
type DBModel struct {
	DB *mongo.Database
}

// Models is the wrapper for the database.
type Models struct {
	DB DBModel
}

// New Models returns models with db pool
func NewModels(db *mongo.Database) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// Damage is the type for weapon damage
type WeaponDamage struct {
	Min int `json:"min,omitempty" bson:"min,omitempty"`
	Max int `json:"max,omitempty" bson:"max,omitempty"`
}

// WeaponPhysical is the type for physical weapon stat
type WeaponPhysical struct {
	HitRate    int `json:"hitRate,omitempty" bson:"hitRate,omitempty"`
	CritChance int `json:"critChance,omitempty" bson:"critChance,omitempty"`
	Crit       int `json:"crit,omitempty" bson:"crit,omitempty"`
}

// Weapon is the type for weapons
type Weapon struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Level         int                `json:"level,omitempty" bson:"level,omitempty"`
	ChampionLevel bool               `json:"championLevel,omitempty" bson:"championLevel,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	Image         string             `json:"image,omitempty" bson:"image,omitempty"`
	Damage        *WeaponDamage      `json:"damage,omitempty" bson:"damage,omitempty"`
	Physical      *WeaponPhysical    `json:"physical,omitempty" bson:"physical,omitempty"`
	Concentration int                `json:"concentration,omitempty" bson:"concentration, omitempty"`
	Effects       []string           `json:"effects,omitempty" bson:"effects,omitempty"`
	HowToGet      []string           `json:"howToGet,omitempty" bson:"howToGet,omitempty"`
}

// FairyLevel is the type for possible fairy levels
type FairyLevel struct {
	Min int `json:"min,omitempty" bson:"min,omitempty"`
	Max int `json:"max,omitempty" bson:"max,omitempty"`
}

// Fairy is the type for fairies
type Fairy struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Level    *FairyLevel        `json:"level,omitempty" bson:"level,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Element  string             `json:"element,omitempty" bson:"element,omitempty"`
	Effects  []string           `json:"effects,omitempty" bson:"effects,omitempty"`
	HowToGet []string           `json:"howToGet,omitempty" bson:"howToGet,omitempty"`
}
