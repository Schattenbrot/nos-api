package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertWeapon is the method for inserting weapons.
func (m *DBModel) InsertWeapon(weapon Weapon) (*primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.DB.Collection("weapons")

	oid, err := collection.InsertOne(ctx, weapon)
	if err != nil {
		return nil, err
	}

	result := oid.InsertedID.(primitive.ObjectID)
	return &result, nil
}

// FindAllWeapons is the method for finding all weapons.
func (m *DBModel) FindAllWeapons() ([]*Weapon, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := m.DB.Collection("weapons")

	findOptions := *options.Find()

	cursor, err := collection.Find(ctx, bson.D{}, &findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var weapons []*Weapon

	for cursor.Next(ctx) {
		var weapon Weapon
		cursor.Decode(&weapon)
		weapons = append(weapons, &weapon)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return weapons, nil
}
