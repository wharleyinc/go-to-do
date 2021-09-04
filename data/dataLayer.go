package dataLayer

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDataWale() *mongo.Client {

	// create a new context with a 10 second timeout
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ctx := context.Background()

	// create a mongo client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		panic(err)
	} else {
		log.Println("Connected to MongoDB DATA")
	}
	// disconnects from mongo when there's error
	// defer Client.Disconnect(ctx)

	// define the database
	database := client.Database("ToDoList")

	// define/set the collection/table
	todosCollection := database.Collection("todoList")
	log.Fatal(todosCollection)

	return client
}
