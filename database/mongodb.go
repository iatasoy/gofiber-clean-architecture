package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"gofiber-clean-architecture/configuration"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Mg MongoInstance

func Connect() error {
	configuration.LoadConfig()

	mongoURI := configuration.Get("MONGODB_URI")
	dbName := configuration.Get("DB_NAME")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(mongoURI).SetMaxPoolSize(50)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return fmt.Errorf("mongo connect failed: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("mongo ping failed: %w", err)
	}

	db := client.Database(dbName)
	Mg = MongoInstance{Client: client, Db: db}
	UserCollection = db.Collection("users")

	log.Println("✅ Connected to MongoDB")
	return nil
}
