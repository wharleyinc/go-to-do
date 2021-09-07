package dataLayer

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"wharleyinc.com/to-do/models"
)

// collection object/instance
var Collection *mongo.Collection

/* func InitDataWale() {

	// create a new context with a 10 second timeout
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// mongoDbLocal := "mongodb://localhost:27017/"
	mongoDbAtlas := "mongodb+srv://wharley01:Sanctity12!@@wharleycluster01.j4stb.mongodb.net/ToDoList?retryWrites=true&w=majority"

	ctx := context.Background()

	// define the database
	database := "ToDoList"

	// define/set the collection/table
	todosCollection := "todoList"
	// log.Fatal(todosCollection)

	// create a mongo client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDbAtlas))
	if err != nil {
		panic(err)
	} else {
		log.Println("Connected to MongoDB DATA")
	}
	// disconnects from mongo when there's error
	// defer Client.Disconnect(ctx)

	Collection = client.Database(database).Collection(todosCollection)

	fmt.Println("Collection instance created!")

} */

func UseMongoDbAtlas() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://wharley01:Sanctity1%40@wharleycluster01.j4stb.mongodb.net/ToDoList?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	
	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database("ToDoList").Collection("todoList")
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
		panic(err)
	}

	fmt.Println("Inserted a Single TODO Record ", insertResult.InsertedID)
}

// get one todo from db
func GetTodo(id primitive.ObjectID) (models.ToDo) {
	var todo models.ToDo
	
	err := Collection.
		FindOne(context.Background(), bson.D{{Key: "_id", Value: id}}).
		Decode(&todo)
	if err != nil {
		panic(err)
	}
	return todo
}

// update status of todo from db
func UpdateTodo(id primitive.ObjectID, status string) error {
	// to get rid remove the primitive.E composite literal uses unkeyed fields error? https://stackoverflow.com/a/67651664
	filter := bson.D{{Key: "_id", Value: id}}
	// to get rid remove the primitive.E composite literal uses unkeyed fields error? https://stackoverflow.com/a/67651664
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: status}}}}
	_, err := Collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	return err
}


// delete todo  by id from db
func DeleteTodo(id primitive.ObjectID) error {
	_, err := Collection.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return err
	}
	return nil
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
