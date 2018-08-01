package controller

import (
	"net/http"

	"../model"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// TodoController represents the todo controller
type TodoController struct {
	db *sql.DB
}

// List lists all todo
func (tc *TodoController) List(c *gin.Context) {

	todoModel := model.NewTodoModel(tc.db)
	listTodos := todoModel.AllTodos()

	c.JSON(http.StatusOK, listTodos)
}

// Create creates a new todo
func (tc *TodoController) Create(c *gin.Context) {

	todoModel := model.NewTodoModel(tc.db)
	c.BindJSON(&todoModel)

	result := todoModel.CreateTodo(todoModel)

	if result == true {
		c.JSON(http.StatusOK, &model.TodoModel{
			ID:        todoModel.ID,
			UserID:    todoModel.UserID,
			Title:     todoModel.Title,
			Completed: todoModel.Completed,
		})
	}

	{
		c.JSON(http.StatusBadRequest, "Cannot Process the request")
	}

}

// Get gets a todo
func (tc *TodoController) Get(c *gin.Context) {

	todoModel := model.NewTodoModel(tc.db)
	result := todoModel.GetTodo(c.Params.ByName("id"))

	if result.ID != "" {
		c.JSON(http.StatusOK, result)
	}

	{
		c.JSON(http.StatusBadRequest, "Cannot Get the Data")
	}
}

// Update updates a todo
func (tc *TodoController) Update(c *gin.Context) {

	todoModel := model.NewTodoModel(tc.db)
	c.BindJSON(&todoModel)
	result := todoModel.UpdateTodo(c.Params.ByName("id"), todoModel)

	if result == true {
		c.JSON(http.StatusOK, "Success Update the Data")
	}

	{
		c.JSON(http.StatusBadRequest, "Cannot Update the Data")
	}
}

// Delete deletes a todo
func (tc *TodoController) Delete(c *gin.Context) {

	todoModel := model.NewTodoModel(tc.db)
	result := todoModel.DeleteTodo(c.Params.ByName("id"))

	if result == true {
		c.JSON(http.StatusOK, "Success Delete the Data")
	}

	{
		c.JSON(http.StatusBadRequest, "Cannot Delete the Data")
	}
}

// NewTodoController creates a new todo controller
func NewTodoController(db *sql.DB) *TodoController {

	return &TodoController{db: db}
}
