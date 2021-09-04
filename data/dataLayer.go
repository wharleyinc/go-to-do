package dataLayer

import (
	"context"
	"log"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"wharleyinc.com/to-do/models"
	
)

// collection object/instance
var Collection *mongo.Collection

func InitDataWale() {

	// create a new context with a 10 second timeout
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ctx := context.Background()

	// define the database
	database := "ToDoList"

	// define/set the collection/table
	todosCollection := "todoList"
	// log.Fatal(todosCollection)

	// create a mongo client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		panic(err)
	} else {
		log.Println("Connected to MongoDB DATA")
	}
	// disconnects from mongo when there's error
	// defer Client.Disconnect(ctx)

	Collection = client.Database(database).Collection(todosCollection)

	fmt.Println("Collection instance created!")

}

func GetAllTodos() []primitive.M {
	cur, err := Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		panic(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			panic(e)
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		panic(err)
	}

	cur.Close(context.Background())
	return results
}

// Insert one todo in the DB
func CreateTodo(todo models.ToDo) {
	insertResult, err := Collection.InsertOne(context.Background(), todo)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single TODO Record ", insertResult.InsertedID)
}


// todo complete method, update todo's status to true
func todoComplete(todo string) {
	fmt.Println(todo)
	id, _ := primitive.ObjectIDFromHex(todo)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// todo undo method, update todo's status to false
func undoTodo(todo string) {
	fmt.Println(todo)
	id, _ := primitive.ObjectIDFromHex(todo)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete one todo from the DB, delete by ID
func deleteOneTodo(todo string) {
	fmt.Println(todo)
	id, _ := primitive.ObjectIDFromHex(todo)
	filter := bson.M{"_id": id}
	d, err := Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}

// delete all the todos from the DB
func deleteAllTodo() int64 {
	d, err := Collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
	return d.DeletedCount
}
