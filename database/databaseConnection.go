package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func dbInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error while loading env file")
		return nil
	}

	mongoUrl := os.Getenv("MONGO_URL")

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.second)
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	if err != nil{
		log.Fatal("error while creating mongo client")
		return nil
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("error while connecting to mongo client")
		return nil
	}

	return client
}

var Client *mongo.Client = dbInstance()
