package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo/models"
	"todo/repositories"
)

type TodoHandler struct {
	repository repositories.TodoRepository
}

type TodoHandlerImpl interface {
	GetTodos(c *gin.Context)
	CreateTodo(c *gin.Context)
	GetTodoByID(c *gin.Context)
	UpdateTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.repository.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo models.Todo
	// Bind the incoming JSON to the todostruct
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use the repository to save the new todo item
	createdTodo, err := h.repository.Create(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdTodo})
}

func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.repository.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {

	// Get model if exist
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := h.repository.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTodo, err := h.repository.Update(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedTodo})
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {

	// Check if the todo item exists

	id, _ := strconv.Atoi(c.Param("id"))
	_, err := h.repository.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	err = h.repository.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func NewTodoHandler(repository repositories.TodoRepository) TodoHandlerImpl {
	return &TodoHandler{repository: repository}
}
