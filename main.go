package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gocrud/controllers"
)

func main() {
	// Create a new Gin router instance
	r := gin.Default()

	// Create a new UserController instance with a MongoDB client
	uc := controllers.NewUserController(getClient())

	// Define routes for handling HTTP requests
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	r.PUT("/user/:id", uc.UpdateUser)

	// Run the server on port 9000
	r.Run(":9000")
}

// Function to create a MongoDB client
func getClient() *mongo.Client {
	// Define MongoDB connection options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Set a timeout for connecting to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping MongoDB to check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Return the MongoDB client
	return client
}
