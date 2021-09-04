package services

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"net/http"

	"github.com/gin-gonic/gin"

	dataLayer "wharleyinc.com/to-do/data"
	"wharleyinc.com/to-do/models"
)

// create connection with mongo db
/* func init() {
	createDBInstance()
} */

/* func createDBInstance() {

// create a new context with a 10 second timeout
ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// ctx := context.Background()

// create a mongo client
client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/"))
if err != nil {
	panic(err)
}
// disconnects from mongo when there's error
defer client.Disconnect(ctx)


// define the database
database = client.Database("ToDoList")

// define/set the collection/table
TodosCollection := Database.Collection("todoList")
*/

/* // DB connection string
connectionString := "mongodb://localhost:27017"

// Database Name
dbName := "toDoWale"

// Collection name
collName := "toDOList"


// Set client options
clientOptions := options.Client().ApplyURI(connectionString)

// Connect to MongoDB
client, err := mongo.Connect(context.TODO(), clientOptions)

if err != nil {
	log.Fatal(err)
}

// Check the connection
err = client.Ping(context.TODO(), nil)

if err != nil {
	log.Fatal(err)
}

fmt.Println("Connected to MongoDB!")

// collection := client.Database("TodoList").Collection("todolist")

collection = client.Database(dbName).Collection(collName)

fmt.Printf("Collection instance created!") */

/* } */

var todos = []models.ToDo{}

type mongoClient struct {
	client *mongo.Client
}

func (mC *mongoClient) FindByID(ctx context.Context, id primitive.ObjectID) (*models.ToDo, error) {
	var todo models.ToDo
	err := mC.client.Database("TodoList").Collection("todoList").FindOne(ctx, bson.M{"_id": id}).Decode((&todo))
	if err != nil {
		return nil, err
	}
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &todo, nil
}

// getAllTodosWale, responds with the list of all todos as JSON.
func GetAllTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dataLayer.GetAllTodos())
}

// getAllTodosWale, responds with the list of all todos as JSON.
func CreateTodo(c *gin.Context) {
	var newTodo models.ToDo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	dataLayer.CreateTodo(newTodo)
}

/* func getAllTodos(c *gin.Context) {
	var newTodo models.ToDo

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	// Add the new album to the slice.
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)

	// insertResult, err := Client.Database("TodoList").Collection("todoList").InsertOne(context.TODO(), newTodo)
	insertResult, err := dataLayer.InitDataWale().Database("TodoList").Collection("todoList").InsertOne(context.TODO(), newTodo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
} */

// Insert a single record to the db
/* func CreateTodo(ctx, c *gin.Context) {
	var newTodo models.ToDo

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	// Add the new album to the slice.
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)

	createResult, err := todosCollection.InsertOne(ctx, todo)
	if err != nil {
		panic(err)
	}
	fmt.Println(createResult.InsertedID)
} */

// GetTodos responds with the list of all toDos as JSON.
func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

// CreateTodos adds an album from JSON received in the request body.
func CreateTodos(c *gin.Context) {
	var newTodo models.ToDo

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	// Add the new album to the slice.
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
/* func GetTodoByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range todos {
		if a.ID == id. {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
} */
