package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	db, err := sql.Open("pgx", "postgresql://postgres:postgres@localhost:5432/bird_encyclopedia")
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}
	fmt.Println("database reachable")
	return db
}

func migrateDb() {
	db, err := sql.Open("pgx", "postgresql://postgres:postgres@localhost:5432/bird_encyclopedia")
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}
	instance, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	migrations, err := migrate.NewWithDatabaseInstance("file://./db/migrations", "postgres", instance)
	if err != nil {
		log.Fatal(err)
	}
	if err := migrations.Up(); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	migrateDb()
	r := newRouter()
	http.ListenAndServe(":8080", r)
}
