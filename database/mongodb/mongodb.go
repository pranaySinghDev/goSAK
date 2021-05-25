package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
		return err
	}
	return nil
}

func (db *Mongodb) GetByID(ctx context.Context, database, table string, entityId string, entity interface{}) error {
	query := bson.D{{Key: "_id", Value: entityId}}
	cursor := db.client.Database(database).Collection(table).FindOne(ctx, query)
	err := cursor.Decode(entity)
	if err != nil {
		return err
	}
	return nil
}

func (db *Mongodb) GetAll(ctx context.Context, database, table string, entities interface{}, limit int64, index int64) error {
	query := bson.D{{}}
	var findoptions *options.FindOptions
	skip := limit * index
	if limit > 0 {
		findoptions = &options.FindOptions{
			Limit: &limit,
			Skip:  &skip,
		}
	}

	cursor, err := db.client.Database(database).Collection(table).Find(ctx, query, findoptions)
	if err != nil {
		return err
	}
	if err := cursor.All(ctx, entities); err != nil {
		return err
	}
	return nil
}

func (db *Mongodb) Update(ctx context.Context, enitty string) error   { return nil }
func (db *Mongodb) Delete(ctx context.Context, entityId string) error { return nil }
