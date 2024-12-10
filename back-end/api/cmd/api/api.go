package api

import (
	"database/sql"
	"file-sync-tool-api/service/project"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	address string
	db      *sql.DB
}

func NewServer(address string, db *sql.DB) *Server {
	return &Server{
		address: address,
		db:      db,
	}
}

func (server *Server) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	projectStore := project.NewStore(server.db)
	projectHandler := project.NewHandler(projectStore)
	projectHandler.RegisterRoutes(subrouter)

	log.Println("Server is running on", server.address)

	return http.ListenAndServe(server.address, router)
}
