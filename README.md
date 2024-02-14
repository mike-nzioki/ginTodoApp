# ginTodoApp
This is a simple RESTful To-Do application built with the Gin web framework in Go and uses MySQL as the database. It provides basic CRUD functionalities for managing to-do items.
# Project Structure
The project has a simple structure:
```azure
/todo
├── db.go                  # Database connection setup and configuration.
├── main.go                # The entry point of the application where the server is started.
├── models                 # Data models for the application.
│   └── todo.go            # The To-Do model.
├── repositories           # Repository layer for encapsulating the logic that interacts with the database.
└── todorepository.go      # Repository for To-Do operations.
└── handlers               # Handlers for HTTP requests.
    └── todohandler.go     # HTTP handlers for To-Do operations.
├── go.mod                 # Go module file containing dependencies.
├── go.sum                 # Contains the expected cryptographic checksums of the content of specific module versions.
└── README.md              # This file.
```
# Installing Dependencies
Open a terminal in the project directory and run:
```
 go mod tidy  
```
# Running the Application
In the terminal, execute:
``` 
go run main.go 
```