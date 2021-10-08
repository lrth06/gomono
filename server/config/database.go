package config

import(
	"os"
	"context"
    "fmt"
    "log"

    // "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func Database(){

	URI := os.Getenv("MONGO_URI")	
	
	fmt.Println(URI)

	mongoURI := "mongodb://mongodb:27017"
clientOptions := options.Client().ApplyURI(mongoURI)
client, err := mongo.Connect(context.TODO(), clientOptions)

if err != nil {
    log.Fatal(err)
}

err = client.Ping(context.TODO(), nil)

if err != nil {
    log.Fatal(err)
}

fmt.Println("Connected to MongoDB!")
}