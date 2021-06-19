package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type Config struct {
	MongoAddress string
	MongoDBName  string
}

var db *mongo.Database

func Init(parseConfig func(v interface{}) error) error {
	var cfg Config

	if err := parseConfig(&cfg); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoAddress))
	if err != nil {
		return err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	db = client.Database(cfg.MongoDBName)

	return nil
}
