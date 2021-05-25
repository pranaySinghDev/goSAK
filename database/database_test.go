package database

import (
	"context"
	"log"
	"testing"

	guuid "github.com/google/uuid"
	"github.com/pranaySinghDev/goSAK/database/config"
)

type User struct {
	ID      string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string  `json:"name"`
	Product string  `json:"product"`
	Age     float64 `json:"age"`
}

const (
	url = "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"
)

func TestInsertDBEntity(t *testing.T) {
	db, err := Build(&config.DBConfig{
		Type: config.Mongodb,
		URL:  url,
	})
	if err != nil {
		log.Fatalf("Couldn't build database Factory: %v", err)
	}
	entity := &User{
		ID:      guuid.New().String(),
		Name:    "Jay",
		Product: "Item2",
		Age:     45,
	}
	err = db.Create(context.Background(), "awesomeApp", "users", entity)
	if err != nil {
		log.Fatalf("Couldn't insert into database: %v", err)
	}
}

func TestGetDBEntityByID(t *testing.T) {
	db, err := Build(&config.DBConfig{
		Type: config.Mongodb,
		URL:  url,
	})
	if err != nil {
		log.Fatalf("Couldn't build database Factory: %v", err)
	}
	entity := &User{}
	err = db.GetByID(context.Background(), "awesomeApp", "users", "2e0ce5b6-7b90-4bf6-9cde-8dd6734ec6c4", entity)
	if err != nil {
		log.Fatalf("Couldn't get by id: %v", err)
	}
}

func TestGetAllDBEntity(t *testing.T) {
	db, err := Build(&config.DBConfig{
		Type: config.Mongodb,
		URL:  url,
	})
	if err != nil {
		log.Fatalf("Couldn't build database Factory: %v", err)
	}
	var Users []User = make([]User, 0)
	err = db.GetAll(context.Background(), "awesomeApp", "users", &Users, 10, 1)
	if err != nil {
		log.Fatalf("Couldn't get all: %v", err)
	}
}
