package api

import (
	db "importer-api/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Define routes here
	router.POST("/user", server.createUser)
	router.GET("/user/id/:id", server.getUserById)

	router.GET("/get-fornecedores-by-id/:id", server.getFornecedoresById)

	server.router = router
	return server
}

// Start runs the HTTP server on the specified address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse formats the error message for the API response
func errorResponse(err error) gin.H {
	return gin.H{"api has error": err.Error()}
}
