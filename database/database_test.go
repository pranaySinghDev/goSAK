package database

import (
	"context"
	"log"
	"testing"

	"github.com/pranaySinghDev/goSAK/database/config"
)

func TestMongoDBFactory(t *testing.T) {
	db, err := Build(&config.DBConfig{
		Type: config.Mongodb,
		URL:  "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb",
	})
	if err != nil {
		log.Fatalf("Couldn't build database Factory: %v", err)
	}
	data := map[string]string{"name": "pranay"}
	err = db.Create(context.Background(), "awesomeApp", "users", data)
	if err != nil {
		log.Fatalf("Couldn't insert into database: %v", err)
	}
}
