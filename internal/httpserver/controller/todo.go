package httpserver/controller 

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// TodoController represents the todo controller
type TodoController struct{}

// TodoResponse represents the renderer response
type TodoResponse struct {
    ID        string `json:"id"`
    UserID    string `json:"userId"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

// List lists all todo
func (tc *TodoController) List(c *gin.Context) {
    c.JSON(http.StatusOK, []*TodoResponse{
        &TodoResponse{
            ID:        "todo001",
            UserID:    "user001",
            Title:     "Hello",
            Completed: false,
        },
    })
}

// Create creates a new todo
func (tc *TodoController) Create(c *gin.Context) {
    c.JSON(http.StatusOK, &TodoResponse{
        ID:        "todo002",
        UserID:    "user001",
        Title:     "Hello",
        Completed: false,
    })
}

// Get gets a todo
func (tc *TodoController) Get(c *gin.Context) {
    c.JSON(http.StatusOK, &TodoResponse{
        ID:        "todo001",
        UserID:    "user001",
        Title:     "Hello",
        Completed: false,
    })
}

// Update updates a todo
func (tc *TodoController) Update(c *gin.Context) {
    c.JSON(http.StatusOK, &TodoResponse{
        ID:        "todo001",
        UserID:    "user001",
        Title:     "Hello",
        Completed: false,
    })
}

// Delete deletes a todo
func (tc *TodoController) Delete(c *gin.Context) {
    c.String(http.StatusNoContent, "")
}

// NewTodoController creates a new todo controller
func NewTodoController() *TodoController {
    return &TodoController{}
}
