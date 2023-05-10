package main

import (
	"github.com/devnandito/imageapi/api"
	"github.com/devnandito/imageapi/server"
)

func main() {
	http := server.NewServer(":8080")
	// API clients
	http.Handle("GET", "/api/images", api.HandleApiShowImage)
	http.Handle("POST", "/api/upload", api.HandleUploadImage)
	// http.Handle("POST", "/api/clients", api.HandleApiCreateClient)
	// http.Handle("GET", "/api/clients/:id", api.HandleApiPutClient)
	// http.Handle("POST", "/api/users", handlers.HandleUserPostRequest)
	
	http.Listen()
}