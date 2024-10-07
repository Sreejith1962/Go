package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ = godotenv.Load()

var dbUrl = os.Getenv("MongoDB")

func ConnectDb() *mongo.Collection {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = mongo.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("demo_db").Collection("books")
}
