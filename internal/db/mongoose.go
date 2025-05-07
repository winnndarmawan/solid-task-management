package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func ConnectMongo() *mongo.Database {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	viper.AutomaticEnv()

	uri := viper.GetString("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI is not set in the environment")
	}
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.Database("taskapp")
}
