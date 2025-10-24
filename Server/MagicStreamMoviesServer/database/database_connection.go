package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo" // to connect with the mongodb
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect() *mongo.Client {
	// function that connects with mongoDB and returs a pointer to a mongoDB client object
	// load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: unable to find .env file")
	}

	MongoDB := os.Getenv("MONGODB_URI")
	// check if MONGODB_URI var exists
	if MongoDB == "" {
		log.Fatal("MONGODB_URI not set!")
	}

	fmt.Println("MongoDB URI: ", MongoDB)

	// create configuration options for the connection (the mongoDB_URI)
	clientOptions := options.Client().ApplyURI(MongoDB)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil
	}

	return client
}

// this line runs once when the package is initialized
// it stores a global MongoDB Client that can be reused anywhere
// whenever the database package is imported this var can be used (we are connected with mongo)
var Client *mongo.Client = Connect()

func OpenCollection(collectionName string) *mongo.Collection {
	// function that returns a pointer to a mongoDB collection from the database
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: unable to find .env file")
	}

	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		log.Fatal("DatabaseName not set!")
	}
	fmt.Println("DATABASE_NAME", databaseName)

	collection := Client.Database(databaseName).Collection(collectionName)
	if collection == nil {
		return nil
	}

	return collection
}
