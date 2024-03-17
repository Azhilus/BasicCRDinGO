# Go CRUD Application with Gin and MongoDB

This is a simple CRUD (Create, Read, Update, Delete) application built using Go (Golang), Gin web framework, and MongoDB. It provides basic functionality to manage user data including creating, retrieving, updating, and deleting users.

## Features

- **Create User**: Add a new user to the database with name, gender, and age.
- **Retrieve User**: Fetch user details by ID.
- **Update User**: Update user details such as name, gender, or age.
- **Delete User**: Remove a user from the database by ID.

## Prerequisites

Before running the application, make sure you have the following installed on your system:

- Go programming language (version 1.16 or higher)
- MongoDB database server
- `go.mongodb.org/mongo-driver` package for MongoDB Go driver
- `github.com/gin-gonic/gin` package for Gin web framework
- `github.com/julienschmidt/httprouter` package for HTTP router (only needed for previous versions)

## Installation

1. Clone this repository to your local machine:

    ```bash
    git clone <repository-url>
    ```

2. Navigate to the project directory:

    ```bash
    cd gocrud
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Start the MongoDB server on your local machine.

## Configuration

- The MongoDB connection URI is set to `mongodb://localhost:27017` by default. Update the URI in the code if your MongoDB server is running on a different host or port.

## Usage

1. Run the application:

    ```bash
    go run main.go
    ```

2. The application will start and listen for incoming HTTP requests on `localhost:9000` by default.

## API Endpoints

- **GET /user/:id**: Retrieve user details by ID.
- **POST /user**: Create a new user.
- **PUT /user/:id**: Update user details by ID.
- **DELETE /user/:id**: Delete a user by ID.

## Examples

### Create User

```bash
curl -X POST \
  http://localhost:9000/user \
  -H 'Content-Type: application/json' \
  -d '{
	"name": "John Doe",
	"gender": "male",
	"age": 30
}'
```

### Retrieve User

```bash
curl -X GET http://localhost:9000/user/<user-id>
```

### Update User

```bash
curl -X PUT \
  http://localhost:9000/user/<user-id> \
  -H 'Content-Type: application/json' \
  -d '{
	"name": "Updated Name",
	"gender": "female",
	"age": 35
}'
```

### Delete User

```bash
curl -X DELETE http://localhost:9000/user/<user-id>
```

Replace `<user-id>` with the actual ID of the user.
