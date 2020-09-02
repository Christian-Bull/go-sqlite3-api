package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// Status trail status struct
type Status struct {
	ID     string `json:"Id"`
	Status int    `json:"Status"`
	Text   string `json:"Text"`
}

// Tweet struct
type Tweet struct {
	ID        string
	createdat string
	text      string
	status    int
}

// Statuses list of status
var status Status

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Base page")
	fmt.Println("Endpoint hit: home")
}

// returns the most recent trail status from the database
func trailStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: trail status")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func handleRequests() {
	// create new mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/currentstatus", trailStatus)
	log.Fatal(http.ListenAndServe(":2000", myRouter))
}

func getStatus(db *sql.DB) Status {
	t := Tweet{}

	err := db.QueryRow("SELECT ID, MAX(created_at), tweet_text, status FROM tweets", 1).Scan(
		&t.ID, &t.createdat, &t.text, &t.status,
	)
	if err != nil {
		fmt.Println(err)
	}
	return Status{ID: t.ID, Status: t.status, Text: t.text}
}

func main() {
	// open db and defer close until execute
	sqliteDatabase, err := sql.Open("sqlite3", os.Getenv("sqldatabase"))
	if err != nil {
		fmt.Println(err)
	}
	defer sqliteDatabase.Close()
	status = getStatus(sqliteDatabase)

	// run server
	handleRequests()
}
