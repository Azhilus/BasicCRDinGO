### UserController Struct

```go
// UserController handles user-related HTTP requests
type UserController struct {
	client *mongo.Client
}
```

- `UserController`: This struct represents a controller responsible for handling HTTP requests related to users.
- `client *mongo.Client`: This field holds a reference to the MongoDB client for database operations.

### NewUserController Function

```go
// NewUserController creates a new UserController instance
func NewUserController(client *mongo.Client) *UserController {
	return &UserController{client}
}
```

- `NewUserController`: This function creates a new instance of UserController with the provided MongoDB client.
- `client *mongo.Client`: The MongoDB client is passed as a parameter to initialize the UserController.

### GetUser Method

```go
// GetUser retrieves a user by ID
func (uc UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid ID"})
		return
	}
	u := models.User{}
	collection := uc.client.Database("mongo-golang").Collection("users")
	err = collection.FindOne(c.Request.Context(), bson.M{"_id": oid}).Decode(&u)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding user"})
		return
	}
	c.JSON(http.StatusOK, u)
}
```

- `GetUser`: This method retrieves a user by ID from the MongoDB collection.
- `id := c.Param("id")`: Extracts the user ID from the request URL parameters.
- `oid, err := primitive.ObjectIDFromHex(id)`: Converts the ID string to an ObjectID, which MongoDB uses for querying documents.
- `u := models.User{}`: Initializes an empty User struct to hold the retrieved user data.
- `collection := uc.client.Database("mongo-golang").Collection("users")`: Retrieves the MongoDB collection named "users" from the client's database.
- `err = collection.FindOne(c.Request.Context(), bson.M{"_id": oid}).Decode(&u)`: Finds a document in the collection with the specified ID and decodes it into the User struct.
- Error handling is included for cases where the ID is invalid or the user is not found in the database.

### CreateUser Method

```go
// CreateUser creates a new user
func (uc UserController) CreateUser(c *gin.Context) {
	var u models.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u.ID = primitive.NewObjectID()
	collection := uc.client.Database("mongo-golang").Collection("users")
	if _, err := collection.InsertOne(c.Request.Context(), u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}
	c.JSON(http.StatusCreated, u)
}
```

- `CreateUser`: This method creates a new user using the data provided in the request body.
- `var u models.User`: Initializes a new User struct to hold the user data.
- `if err := c.BindJSON(&u); err != nil { ... }`: Parses the JSON request body into the User struct.
- `u.ID = primitive.NewObjectID()`: Generates a new ObjectID for the user.
- `collection := uc.client.Database("mongo-golang").Collection("users")`: Retrieves the "users" collection from the MongoDB client's database.
- `_, err := collection.InsertOne(c.Request.Context(), u)`: Inserts the new user document into the collection.
- Error handling is included for cases where the request body is invalid or there is an error inserting the user into the database.

### DeleteUser Method

```go
// DeleteUser deletes a user by ID
func (uc UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid ID"})
		return
	}
	collection := uc.client.Database("mongo-golang").Collection("users")
	if _, err := collection.DeleteOne(c.Request.Context(), bson.M{"_id": oid}); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
```

- `DeleteUser`: This method deletes a user by ID from the MongoDB collection.
- `id := c.Param("id")`: Extracts the user ID from the request URL parameters.
- `oid, err := primitive.ObjectIDFromHex(id)`: Converts the ID string to an ObjectID.
- `collection := uc.client.Database("mongo-golang").Collection("users")`: Retrieves the "users" collection from the MongoDB client's database.
- `_, err := collection.DeleteOne(c.Request.Context(), bson.M{"_id": oid})`: Deletes the user document from the collection based on the provided ID.
- Error handling is included for cases where the ID is invalid or the user is not found in the database.

### UpdateUser Method

```go
// UpdateUser updates a user by ID
func (uc UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid ID"})
		return
	}
	var u models.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	collection := uc.client.Database("mongo-golang").Collection("users")
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": bson.M{
		"name":   u.Name,
		"gender": u.Gender,
		"age":    u.Age,
	}}
	_, err = collection.UpdateOne(c.Request.Context(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}
```

- `UpdateUser`: This method updates a user by ID in the MongoDB collection.
- `id := c.Param("id")`: Extracts the user ID from the request URL parameters.
- `oid, err := primitive.ObjectID

FromHex(id)`: Converts the ID string to an ObjectID.
- `var u models.User`: Initializes a new User struct to hold the updated user data.
- `if err := c.BindJSON(&u); err != nil { ... }`: Parses the JSON request body into the User struct.
- `collection := uc.client.Database("mongo-golang").Collection("users")`: Retrieves the "users" collection from the MongoDB client's database.
- Constructs a filter and update document for the MongoDB update operation.
- `_, err = collection.UpdateOne(c.Request.Context(), filter, update)`: Performs the update operation on the user document.
- Error handling is included for cases where the ID is invalid, the request payload is invalid, or there is an error updating the user.