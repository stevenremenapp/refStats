package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Interaction struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Timestamp string `json:"time"`
}

type Interactions []Interaction

func allInteractions(w http.ResponseWriter, r *http.Request) {
	interactions := Interactions{
		Interaction{
			ID:        1,
			Type:      "tech",
			Timestamp: time.Now().Format(time.Kitchen),
		},
		Interaction{
			ID:        2,
			Type:      "reference",
			Timestamp: time.Now().Add(time.Hour * 5).Format(time.Kitchen),
		},
		Interaction{
			ID:        3,
			Type:      "tech",
			Timestamp: time.Now().Add(time.Hour * 2).Format(time.Kitchen),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Hit all interactions endpoint")
	json.NewEncoder(w).Encode(interactions)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "HomePage!")
}

func handleRequests() {

	// c := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"http://localhost:8080/"},
	// })

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/interactions", allInteractions).Methods("GET")
	handler := cors.Default().Handler(myRouter)
	http.ListenAndServe(":5000", handler)
}

func main() {
	handleRequests()
}
