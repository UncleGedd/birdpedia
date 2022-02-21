package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/http"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	staticFileDir := http.Dir("./assets")
	staticFileHandler := http.FileServer(staticFileDir)
	r.Path("/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return r
}

func connectToDb() *sql.DB {
	db, err := sql.Open("pgx", "postgresql://postgres@localhost:5432/bird_encyclopedia")
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}
	fmt.Println("database reachable")
	return db
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}
