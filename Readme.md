# CAR API

This is a basic CRUD API that has been created in Golang using the Gofr framework. This API has an extra layer of authentication using JWT.



## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Testing](#testing)
- [Usage](#usage)
- [Diagrams](#UML/SEQUENCE Diagrams)
- [Contact](#contact)


## Introduction

The Car API is designed to provide CRUD (Create, Read, Update, Delete) functionalities for managing cars. It includes authentication using JWT for enhanced security.

## Features

- **Login & Signup**: Endpoints for user login and signup with JWT authentication.
- **Get All Cars**: Retrieve a list of all cars in the database.
- **Create a Car**: Add a new car to the database.
- **Get Car by ID**: Retrieve a specific car by its unique carno.
- **Update Car**: Update details of a specific car.
- **Delete Car**: Remove a car from the database.

## Installation

To set up the Car API on your local machine, follow these steps:

1. Clone this repository.
2. Install dependencies using `go mod tidy`.
3. Ensure you have a local MongoDB instance running for regular car CRUD operations.
4. Set up an Atlas MongoDB cluster and obtain the connection string.
5. Add the Atlas MongoDB connection string to the project configuration.
6. Run the application using `go run main.go`.
7. Ensure that Atlas MongoDB is Connected for authorization 
8. Else you can remove midlleware from main.go routes to run CRUD on local instance

### Configuring MongoDB

For local car CRUD operations, ensure your local MongoDB instance is running and accessible. You may need to configure the MongoDB connection details in the project.

For CRUD CAR using Atlas MongoDB:

1. Create an Atlas MongoDB cluster and obtain the connection string.
2. Modify the respective files to connect to Atlas MongoDB for authentication and CRUD .
3. Update `main.go` to set up the connection using the Atlas MongoDB client.

### Example Changes

```go
// Add these lines to main.go to set up the connection
store := cars.New(client)
h := handlers.New(store)
stores.SetMongoClient(client)
handlers.SetMongoClient(client)
```

```go
// Add these lines to stores.go
type car struct {
	client *mongo.Client
}

func New(client *mongo.Client) car {
	return car{client: client}
}
```
```go
// Add these updates in handlers.go in the handlers folder
type handler struct {
	store stores.Car
}

func New(store stores.Car) handler {
	return handler{
		store: store,
	}
}
```
## Testing

You can run the test files for different files and handlers

For testing main.go 
make sure that main.go is running as for authenticatio API uses ATLAS and for CAR it uses Local DB and 

Before testing Change the username and password that you are going to send in main_test.go
```go
{"register success", http.MethodPost, "signup", http.StatusCreated, []byte(`{"username": "changethis","changethis": "1234567890"}`)},
{"login success", http.MethodPost, "login", http.StatusCreated, []byte(`{"username": "changethis","password": "changethis"}`)},
```
Type this i your CLI 

```CLI
go test
```
you can do same for other test file just go to the package thatyou are going to test 

Example for package handlers

```CLI
cd handlers
go test
```



## Usage

The Car API provides various endpoints for managing cars. Here are some key endpoints:

- **POST /login**: Login with user credentials.
- **POST /signup**: Sign up with user credentials.
- **GET /cars**: Retrieve a list of all cars.
- **POST /cars**: Add a new car.
- **GET /cars/{carno}**: Get details of a specific car by carno.
- **PUT /cars/{carno}**: Update details of a specific car by carno.
- **DELETE /cars/{carno}**: Delete a car by carno.

### API Specifications

[Swagger File](./swagger.yaml)

## UML/SEQUENCE DIAGRAM

### SEQUENCE DIAGRAM
![Alt text](https://github.com/Winkingroad/GofrProject/blob/final/Sequence%20diagram.png)

### UML DIAGRAM
![Alt text](https://github.com/Winkingroad/GofrProject/blob/final/UML%20diagram.png)


## Contact

For any questions or feedback, please reach out to yash.srivastava1102@gmail.com.
