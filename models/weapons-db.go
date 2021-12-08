package models

import (
	"context"
	"log"
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

// FindAllWeaponsByProfession returns all weapons by their profession
func (m *DBModel) FindAllWeaponsByProfession(profession string) ([]*Weapon, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	collection := m.DB.Collection("weapons")

	// filter := *options.Find(Weapon{Profession: profession})
	// filter := bson.D{{"profession", profession}}
	filter := Weapon{Profession: profession}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var weapons []*Weapon

	for cursor.Next(ctx) {
		var weapon Weapon
		cursor.Decode(&weapon)
		weapons = append(weapons, &weapon)
	}

	return weapons, nil
}

// FindOneWeaponById returns one weapon based on its id
func (m *DBModel) FindOneWeaponById(id primitive.ObjectID) (*Weapon, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	collection := m.DB.Collection("weapons")

	// filter := *options.Find(Weapon{Profession: profession})
	// filter := bson.D{{"profession", profession}}
	filter := Weapon{ID: id}

	var weapon Weapon

	err := collection.FindOne(ctx, filter).Decode(&weapon)
	if err != nil {
		return nil, err
	}

	return &weapon, nil
}

// UpdateOneWeaponById updates the weapon based on its id and returns the number of updated weapons.
func (m *DBModel) UpdateOneWeaponById(id primitive.ObjectID, weapon Weapon) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := m.DB.Collection("weapons")
	log.Println("Test:", id)

	filter := Weapon{ID: id}

	var oldWeapon Weapon

	err := collection.FindOne(ctx, filter).Decode(&oldWeapon)
	if err != nil {
		log.Println("uwuff")
		return 0, err
	}

	if oldWeapon.Level != weapon.Level {
		oldWeapon.Level = weapon.Level
	}
	if weapon.ChampionLevel != oldWeapon.ChampionLevel {
		oldWeapon.ChampionLevel = weapon.ChampionLevel
	}
	if weapon.Name != oldWeapon.Name {
		oldWeapon.Name = weapon.Name
	}
	if weapon.Profession != oldWeapon.Profession {
		oldWeapon.Profession = weapon.Profession
	}
	if weapon.Image != oldWeapon.Image {
		oldWeapon.Image = weapon.Image
	}
	if weapon.Damage != nil {
		if weapon.Damage.Max != oldWeapon.Damage.Max {
			oldWeapon.Damage.Max = weapon.Damage.Max
		}
		if weapon.Damage.Min != oldWeapon.Damage.Min {
			oldWeapon.Damage.Min = weapon.Damage.Min
		}
	}
	if weapon.Physical != nil {
		if weapon.Physical.HitRate != oldWeapon.Physical.HitRate {
			oldWeapon.Physical.HitRate = weapon.Physical.HitRate
		}
		if weapon.Physical.CritChance != oldWeapon.Physical.CritChance {
			oldWeapon.Physical.CritChance = weapon.Physical.CritChance
		}
		if weapon.Physical.Crit != oldWeapon.Physical.Crit {
			oldWeapon.Physical.Crit = weapon.Physical.Crit
		}
	}
	if weapon.Concentration != oldWeapon.Concentration {
		oldWeapon.Concentration = weapon.Concentration
	}
	if len(weapon.HowToGet) != len(oldWeapon.HowToGet) {
		oldWeapon.HowToGet = weapon.HowToGet
	}

	update := bson.M{"$set": oldWeapon}
	result, err := collection.UpdateByID(ctx, id, update)
	if err != nil {
		return 0, err
	}

	return int(result.ModifiedCount), nil
}

func (m *DBModel) DeleteWeaponById(id primitive.ObjectID) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	collection := m.DB.Collection("weapons")

	filter := Weapon{ID: id}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil
}
