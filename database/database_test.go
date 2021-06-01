package database

import (
	"context"
	"log"
	"testing"

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
	ctx := context.Background()
	db, err := Build(&config.DBConfig{
		Type: config.Mongodb,
		URL:  url,
	})
	if err != nil {
		log.Fatalf("Couldn't build database Factory: %v", err)
	}
	entity := &User{
		Name:    "Jay",
		Product: "60b5e270e787e76051de9626",
		Age:     45,
	}
	id, err := db.Create(ctx, "awesomeApp", "users", entity)
	if err != nil {
		log.Fatalf("Couldn't insert into database: %v", err)
	}
	log.Printf("id %s", id)
	db.Disconnect(ctx)
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
	err = db.GetByID(context.Background(), "awesomeApp", "users", "60af6f6c35dc24af5e3897ec", entity)
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
