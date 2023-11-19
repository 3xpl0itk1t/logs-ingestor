package handlers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var client *mongo.Client

func ConnectToDB() {
	// Set your MongoDB Atlas connection string
	connectionURI := os.Getenv("MONGODB_URI")
	// Set options for the MongoDB Go driver
	clientOptions := options.Client().ApplyURI(connectionURI)

	// Create a MongoDB client
	client, ERR := mongo.Connect(context.Background(), clientOptions)
	if ERR != nil {
		log.Fatal(ERR)
	}

	// Set a timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ping the MongoDB server to check the connection
	ERR = client.Ping(ctx, nil)
	if ERR != nil {
		log.Fatal(ERR)
	}

	fmt.Println("Connected to MongoDB Atlas!")

	// Perform any database operations here...
	collection = client.Database("log-ingestor").Collection("logs")

}

func DisconnectFromDB() {
	// Disconnect from MongoDB Atlas
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Disconnected from MongoDB Atlas!")
}
