# ginTodoApp
This is a simple RESTful To-Do application built with the Gin web framework in Go and uses MySQL as the database. It provides basic CRUD functionalities for managing to-do items.
# Project Structure
The project has a simple structure:
```azure
/todo
├── main.go # The main application file where the server and routes are defined.
├── go.mod # Go module file containing dependencies.
├── go.sum # Contains the expected cryptographic checksums of the content of specific module versions.
└── README.md # This file.
```
# Installing Dependencies
Open a terminal in the project directory and run:
``` go mod tidy  ```
# Running the Application
In the terminal, execute:
``` go run main.go ```