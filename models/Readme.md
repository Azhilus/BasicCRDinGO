This `models` package defines the structure of the User model, which represents the data schema for users in the application.

```go
package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user model
type User struct {
    ID     primitive.ObjectID `json:"id" bson:"_id"`
    Name   string             `json:"name" bson:"name"`
    Gender string             `json:"gender" bson:"gender"`
    Age    int                `json:"age" bson:"age"`
}
```

- `User`: This struct represents a user in the application.
- `ID primitive.ObjectID`: The unique identifier for the user, represented as an ObjectID. This field is tagged with `bson:"_id"` to specify its mapping to the MongoDB document's "_id" field. The `json:"id"` tag specifies the JSON field name.
- `Name string`: The name of the user. This field is tagged with `bson:"name"` to specify its mapping to the "name" field in the MongoDB document. The `json:"name"` tag specifies the JSON field name.
- `Gender string`: The gender of the user. This field is tagged similarly to the "Name" field.
- `Age int`: The age of the user. This field is tagged similarly to the "Name" field.

**Notes:**
- This model follows the convention of using BSON tags (`bson`) for mapping struct fields to MongoDB document fields and JSON tags (`json`) for specifying field names in JSON serialization/deserialization.
- The `primitive.ObjectID` type from the `go.mongodb.org/mongo-driver/bson/primitive` package is used for representing MongoDB's ObjectIDs. It ensures compatibility with MongoDB's document structure and indexing requirements.
- Each field in the struct corresponds to a specific attribute of a user, such as their name, gender, and age.
- By defining a User struct, this package provides a structured way to represent user data, making it easier to work with and maintain throughout the application.