package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Interaction struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"time"`
}

type Interactions []Interaction

func allInteractions(w http.ResponseWriter, r *http.Request) {
	interactions := Interactions{
		Interaction{
			ID:        1,
			Type:      "tech",
			Timestamp: time.Now(),
		},
		Interaction{
			ID:        2,
			Type:      "reference",
			Timestamp: time.Now().Add(time.Hour * 5),
		},
		Interaction{
			ID:        3,
			Type:      "tech",
			Timestamp: time.Now().Add(time.Hour * 2),
		},
	}

	fmt.Println("Hit all interactions endpoint")
	json.NewEncoder(w).Encode(interactions)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage!")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/interactions", allInteractions).Methods("GET")
	http.ListenAndServe(":5000", myRouter)
}

func main() {
	handleRequests()
}
