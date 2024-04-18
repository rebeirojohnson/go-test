package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define a model struct
type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Height int    `json:"height"`
}

func main() {
	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI("mongodb://199.241.138.96:27017")

	// Connect to MongoDBs
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // Increase timeout to 30 seconds

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	// Close the connection when the application exits
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Create Fiber app
	app := fiber.New()

	// Define routes
	app.Post("/users", func(c *fiber.Ctx) error {
		// Access MongoDB collections
		collection := client.Database("gotest").Collection("users")

		// Parse request body into User struct
		var user User
		if err := c.BodyParser(&user); err != nil {
			return err
		}

		// Insert user document into collection
		_, err := collection.InsertOne(ctx, user)
		if err != nil {
			return err
		}

		return c.JSON(user)
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		// Access MongoDB collections
		collection := client.Database("gotest").Collection("users")

		// Example query: find all users
		filter := bson.M{}

		cursor, err := collection.Find(ctx, filter)
		if err != nil {
			return err
		}
		defer cursor.Close(ctx)

		var users []User
		if err := cursor.All(ctx, &users); err != nil {
			return err
		}

		return c.JSON(users)
	})

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
