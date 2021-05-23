package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	client *mongo.Client
}

//Return db connection
func Connect(url string) (*Mongodb, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &Mongodb{client}, nil

}

func (db *Mongodb) Create(ctx context.Context, database, table string, entity interface{}) error {
	_, err := db.client.Database(database).Collection(table).InsertOne(ctx, entity)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
func (db *Mongodb) Get(ctx context.Context, entityId string) (interface{}, error) { return nil, nil }
func (db *Mongodb) Update(ctx context.Context, enitty string) error               { return nil }
func (db *Mongodb) Delete(ctx context.Context, entityId string) error             { return nil }
