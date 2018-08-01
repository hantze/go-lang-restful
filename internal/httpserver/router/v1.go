package router

import (
	"net/http"

	"../controller"

	"github.com/gin-gonic/gin"

	"database/sql"
)

// NewV1Router creates a new v1 router
func NewV1Router(db *sql.DB) http.Handler {
	c := controller.NewTodoController(db)
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/todos", c.List)    // GET /todos - read a list of todos
		v1.POST("/todos", c.Create) // POST /todos - create a new todo and persist it
		v1.PUT("/todos", c.Delete)
		v1.GET("/todos/:id", c.Get)
		v1.PUT("/todos/:id", c.Update)    // PUT /todos/{id} - update a single todo by :id
		v1.DELETE("/todos/:id", c.Delete) // DELETE /todos/{id} - delete a single todo by :id
	}
	return r
}
