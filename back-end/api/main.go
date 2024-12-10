package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	// connect to database
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	//create router
	router := mux.NewRouter()
	router.HandleFunc("/sync", syncHandler(db)).Methods("GET")
	router.HandleFunc("/push", pushHandler(db)).Methods("POST")
	router.HandleFunc("/history", historyHandler(db)).Methods("GET")
	router.HandleFunc("clone", cloneHandler(db)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", jsonContentTypeMiddleware(router)))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func syncHandler(_ *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// sync data
	}
}

func pushHandler(_ *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// push data
	}
}

func historyHandler(_ *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get history
	}
}

func cloneHandler(_ *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// clone data
	}
}
