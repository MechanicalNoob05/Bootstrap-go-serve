package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client{
	// Loading in .env Variable file
	err := godotenv.Load(".env")
	// Logging if there is any error while Loading it
	if err!= nil{
		log.Fatal("Error loading env Variables",err)
	}
	// getting connection url for mongo db url 
	MongoDb := os.Getenv("MONGODB_URL")
	// applying url and option to create a new session
	client , err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	// Again checking for err 
	if err!= nil{
		log.Fatal("Error with DB",err)
	}
	// This Basically says that try to get connection for 10 sec or cancel shut down the whole operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// this is the function to bail out
	defer cancel()
	// Basically creating bunch of fail safe 
	err = client.Connect(ctx)
	if err!= nil{
		log.Fatal("Error with Client",err)
	}
	// Finally telling that the sweet use
	fmt.Println("Connected to DB")
	// Returning what you asked for the sweet sweet db instance
	return client
}
// Basically we are just capuring the function in the variable
var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	// So basically creating data base cluster in mongo db
	var collection *mongo.Collection = client.Database("clusterduck").Collection(collectionName)

	// And we now return the collection as per the promise by the candy man
	return collection
}
