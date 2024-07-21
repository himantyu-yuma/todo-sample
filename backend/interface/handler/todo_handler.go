package handler

import (
	"net/http"

	"todo-sample/domain/entity"
	"todo-sample/usecase"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoUsecase usecase.TodoUsecase
}

type TodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// @summary Create a new todo
// @description Create a new todo
// @tags todos
// @Router /todos [post]
// @Param body body TodoRequest true "Todo object that needs to be added"
// @Success 201
func NewTodoHandler(todoUsecase usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		todoUsecase: todoUsecase,
	}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo entity.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTodo, err := h.todoUsecase.CreateTodo(c.Request.Context(), todo.Title, todo.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTodo)
}
