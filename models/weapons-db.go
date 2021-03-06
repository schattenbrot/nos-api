package models

import (
	"context"
	"errors"
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

	update := bson.M{"$set": weapon}

	result, err := collection.UpdateByID(ctx, id, update)
	if err != nil {
		return 0, err
	}

	if int(result.ModifiedCount) == 0 {
		err = errors.New("not found")
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

	if result.DeletedCount == 0 {
		return 0, errors.New("not found")
	}

	return int(result.DeletedCount), nil
}
