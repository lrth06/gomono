package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnDB(collection string) *mongo.Collection {
	mongoUri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
	collectionReference := client.Database(dbName).Collection(collection)
	return collectionReference

}
