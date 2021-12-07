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
