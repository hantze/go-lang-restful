package model

import (
	"database/sql"
	"fmt"
)

type TodoModel struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	db        *sql.DB
}

func (tdm *TodoModel) GetTodo(id string) TodoModel {

	query := fmt.Sprintf("SELECT * FROM todo WHERE id = '%s'", id)

	row := tdm.db.QueryRow(query)

	todo := TodoModel{}
	row.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Completed)

	return todo
}

func (tdm *TodoModel) AllTodos() []TodoModel {

	listTodos := []TodoModel{}

	const query = "SELECT * FROM todo"
	rows, _ := tdm.db.Query(query)

	for rows.Next() {
		todo := TodoModel{}
		rows.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Completed)
		listTodos = append(listTodos, todo)
	}

	return listTodos
}

func (tdm *TodoModel) CreateTodo(tm *TodoModel) bool {

	commited := false
	tx, err := tdm.db.Begin()
	if err != nil {
		return false
	}

	{
		const query = "INSERT INTO todo VALUES($1, $2, $3, $4)"
		_, err := tdm.db.Exec(query, tm.ID, tm.UserID, tm.Title, tm.Completed)

		if !commited {
			defer tx.Rollback()
		}

		if err != nil {
			return false
		}

	}

	commited = true
	return true
}

func (tdm *TodoModel) UpdateTodo(id string, tm *TodoModel) bool {

	tx, err := tdm.db.Begin()
	commited := false

	if err != nil {
		return false
	}

	{
		const query = "UPDATE todo SET user_id = $1, title = $2, completed = $3 WHERE id = $4"
		_, err := tdm.db.Exec(query, tm.UserID, tm.Title, tm.Completed, id)

		if !commited {
			defer tx.Rollback()
		}

		if err != nil {
			return false
		}
	}

	commited = true
	return true
}

func (tdm *TodoModel) DeleteTodo(id string) bool {

	commited := false
	tx, err := tdm.db.Begin()
	if err != nil {
		return false
	}

	{
		const query = "DELETE FROM todo WHERE id = $1"
		_, err := tdm.db.Exec(query, id)

		if !commited {
			defer tx.Rollback()
		}

		if err != nil {
			return false
		}

	}

	commited = true
	return true
}

// NewTodoModel creates a new todo model
func NewTodoModel(db *sql.DB) *TodoModel {

	return &TodoModel{db: db}
}
