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

	// Connect to MongoDB with retry
	client, err := connectWithRetry(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
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
		_, err := collection.InsertOne(context.Background(), user)
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

		cursor, err := collection.Find(context.Background(), filter)
		if err != nil {
			return err
		}
		defer cursor.Close(context.Background())

		var users []User
		if err := cursor.All(context.Background(), &users); err != nil {
			return err
		}

		return c.JSON(users)
	})

	// Start the server
	log.Fatal(app.Listen(":3000"))
}

// connectWithRetry establishes connection to MongoDB with retry
func connectWithRetry(clientOptions *options.ClientOptions) (*mongo.Client, error) {
	var client *mongo.Client
	var err error
	for i := 0; i < 3; i++ { // Retry 3 times
		client, err = mongo.NewClient(clientOptions)
		if err != nil {
			log.Println("Failed to create MongoDB client:", err)
			time.Sleep(2 * time.Second) // Wait before retrying
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Increase timeout to 10 seconds
		defer cancel()

		err = client.Connect(ctx)
		if err != nil {
			log.Println("Failed to connect to MongoDB:", err)
			time.Sleep(2 * time.Second) // Wait before retrying
			continue
		}

		// Check the connection
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Println("Failed to ping MongoDB:", err)
			time.Sleep(2 * time.Second) // Wait before retrying
			continue
		}

		log.Println("Connected to MongoDB!")
		return client, nil
	}
	return nil, err
}
