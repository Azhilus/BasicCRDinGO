package controllers

import (
	"gocrud/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// UserController handles user-related HTTP requests
type UserController struct {
	client *mongo.Client
}

// NewUserController creates a new UserController instance
func NewUserController(client *mongo.Client) *UserController {
	return &UserController{client}
}

// GetUser retrieves a user by ID
func (uc UserController) GetUser(c *gin.Context) {
	// Get user ID from request parameters
	id := c.Param("id")

	// Convert ID string to ObjectID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid ID"})
		return
	}

	// Initialize a new User instance
	u := models.User{}

	// Get the MongoDB collection for users
	collection := uc.client.Database("mongo-golang").Collection("users")

	// Find the user by ID
	err = collection.FindOne(c.Request.Context(), bson.M{"_id": oid}).Decode(&u)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding user"})
		return
	}

	// Return the user as JSON response
	c.JSON(http.StatusOK, u)
}

// CreateUser creates a new user
func (uc UserController) CreateUser(c *gin.Context) {
	// Initialize a new User instance
	var u models.User

	// Bind JSON request body to User struct
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a new ObjectID for the user
	u.ID = primitive.NewObjectID()

	// Get the MongoDB collection for users
	collection := uc.client.Database("mongo-golang").Collection("users")

	// Insert the new user into the database
	if _, err := collection.InsertOne(c.Request.Context(), u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	// Return the created user as JSON response
	c.JSON(http.StatusCreated, u)
}

// DeleteUser deletes a user by ID
func (uc UserController) DeleteUser(c *gin.Context) {
	// Get user ID from request parameters
	id := c.Param("id")

	// Convert ID string to ObjectID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid ID"})
		return
	}

	// Get the MongoDB collection for users
	collection := uc.client.Database("mongo-golang").Collection("users")

	// Delete the user from the database
	if _, err := collection.DeleteOne(c.Request.Context(), bson.M{"_id": oid}); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// UpdateUser updates a user by ID
func (uc UserController) UpdateUser(c *gin.Context) {
	// Get user ID from request parameters
	id := c.Param("id")

	// Convert ID string to ObjectID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid ID"})
		return
	}

	// Initialize a new User instance
	var u models.User

	// Bind JSON request body to User struct
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Get the MongoDB collection for users
	collection := uc.client.Database("mongo-golang").Collection("users")

	// Define filter and update for updating the user
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": bson.M{
		"name":   u.Name,
		"gender": u.Gender,
		"age":    u.Age,
	}}

	// Perform the update operation
	_, err = collection.UpdateOne(c.Request.Context(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}
