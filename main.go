package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Todo struct{
	ID int `json:"_id" bson:"_id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}
var collection *mongo.Collection

func main() {
	fmt.Println("Hello, World!")

	const MONGODB_URI string= "mongodb+srv://zukorlicumejr_db_user:KjvH4xHBi1nx8Lpi@cluster0.qz0jxb4.mongodb.net/golang_db?retryWrites=true&w=majority&appName=Cluster0"
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	collection = client.Database("golang_db").Collection("todos")
	app := fiber.New()

	app.Get("/api/todos", getTodos)
	// app.Post("/api/todos", createTodos)
	// app.Patch("/api/todos/:id", updateTodos)
	// app.Delete("/api/todos/:id", deleteTodos)

	app.Listen(":5000")
}
func getTodos(c *fiber.Ctx) error {
	var todos []Todo
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}
	return c.JSON(todos)
	
}
// func createTodos(c *fiber.Ctx) error {}
// func updateTodos(c *fiber.Ctx) error {}
// func deleteTodos(c *fiber.Ctx) error {}
