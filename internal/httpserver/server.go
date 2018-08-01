package httpserver

import (
    "log"
    "net/http"

    "github.com/rickywinata/todos/internal/httpserver/router"
)

// HttpServer represents the http server that handles the requests
type HttpServer struct{}

// Serve serves http requests
func (hs *HttpServer) Serve() {
    r := router.NewV1Router()
    log.Printf("About to listen on 3333. Go to http://127.0.0.1:3333")
    http.ListenAndServe(":3333", r)
}

// NewServer creates a new http server
func NewServer() *HttpServer {
    return &HttpServer{}
}
