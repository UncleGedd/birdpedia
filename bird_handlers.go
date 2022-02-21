package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	db := connectToDb()
	rows, err := db.Query("SELECT bird, description FROM birds LIMIT 10")
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		return
	}
	var birds []Bird
	for rows.Next() {
		bird := Bird{}
		if err := rows.Scan(&bird.Species, &bird.Description); err != nil {
			log.Fatalf("could not scan row: %v", err)
		}
		birds = append(birds, bird)
	}
	birdListBytes, err := json.Marshal(birds)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(birdListBytes)
}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	db := connectToDb()
	newBird := Bird{}
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newBird.Species = r.Form.Get("species")
	newBird.Description = r.Form.Get("description")

	result, err := db.Exec("INSERT INTO birds (bird, description) VALUES ($1, $2)",
		newBird.Species, newBird.Description)

	if err != nil {
		log.Fatalf("could not insert row: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("error retrieving affected rows: %v", err)
	}

	fmt.Println("inserted", rowsAffected, "rows")
	http.Redirect(w, r, "/", http.StatusFound)
}
