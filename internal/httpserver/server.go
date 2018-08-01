package httpserver

import (
	"log"
	"net/http"

	"./router"

	"database/sql"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "hz081ganteng"
	DB_NAME     = "todos"
)

// HttpServer represents the http server that handles the requests
type HttpServer struct{}

// Serve serves http requests
func (hs *HttpServer) Serve(db *sql.DB) {
	r := router.NewV1Router(db)

	log.Printf("About to listen on 3333. Go to http://127.0.0.1:3333")
	http.ListenAndServe(":3333", r)
}

// NewServer creates a new http server
func NewServer() *HttpServer {
	return &HttpServer{}
}
