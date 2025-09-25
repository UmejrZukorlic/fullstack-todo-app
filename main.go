package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Todo struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
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
	app.Post("/api/todos", createTodos)
	app.Patch("/api/todos/:id", updateTodos)
	app.Delete("/api/todos/:id", deleteTodos)

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
func createTodos(c *fiber.Ctx) error {
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Body is required",
		})
	}
	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}
	todo.ID = insertResult.InsertedID.(primitive.ObjectID)
	return c.Status(201).JSON(todo)
}
func updateTodos(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID , err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}
	_ , err = collection.UpdateOne(context.Background(), filter , update)
	if err != nil {
		return err;
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Todo updated successfully",
	})
}
func deleteTodos(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID , err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	filter := bson.M{"_id": objectID}
	_ ,err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Todo deleted successfully",
	}) 

}
