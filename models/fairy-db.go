package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *DBModel) InsertFairy(fairy Fairy) (*primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.DB.Collection("fairies")

	oid, err := collection.InsertOne(ctx, fairy)
	if err != nil {
		return nil, err
	}

	result := oid.InsertedID.(primitive.ObjectID)

	return &result, nil
}

func (m *DBModel) FindAllFairies() ([]*Fairy, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := m.DB.Collection("fairies")

	findOptions := *options.Find()

	cursor, err := collection.Find(ctx, bson.D{}, &findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var fairies []*Fairy

	for cursor.Next(ctx) {
		var fairy Fairy
		cursor.Decode(&fairy)
		fairies = append(fairies, &fairy)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return fairies, nil
}

func (m *DBModel) FindAllFairiesByElement(element string) ([]*Fairy, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	collection := m.DB.Collection("fairies")

	filter := Fairy{Element: element}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var fairies []*Fairy

	for cursor.Next(ctx) {
		var fairy Fairy
		cursor.Decode(&fairy)
		fairies = append(fairies, &fairy)
	}

	return fairies, nil
}

func (m *DBModel) FindFairyById(id primitive.ObjectID) (*Fairy, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	collection := m.DB.Collection("fairies")

	filter := Fairy{ID: id}

	var fairy Fairy

	err := collection.FindOne(ctx, filter).Decode(&fairy)
	if err != nil {
		return nil, err
	}

	return &fairy, nil
}

func (m *DBModel) UpdateFairyById(id primitive.ObjectID, fairy Fairy) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := m.DB.Collection("fairies")

	update := bson.M{"$set": fairy}
	result, err := collection.UpdateByID(ctx, id, update)
	if err != nil {
		return 0, err
	}

	return int(result.ModifiedCount), nil
}

func (m *DBModel) DeleteFairyById(id primitive.ObjectID) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	collection := m.DB.Collection("fairies")

	filter := Fairy{ID: id}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}
