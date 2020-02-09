package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// A function to generate a random number with a random seed.
func generateRanNum(min int, max int) int {
	// Start by getting a new random seed
	rand.Seed(time.Now().UnixNano())

	// Return a random number between min and max
	return rand.Intn(max-min) + min
}

// Define the logic to generate a random word.
func generate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Generate Called")

	// Declare vowels and consonants
	vowels := [5]string{"a", "e", "i", "o", "u"}
	consonants := [21]string{
		"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n",
		"p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}

	// Declare an empty word variable to hold the word
	word := ""

	// Get a random number for the length of our generated word
	maxLength := generateRanNum(4, 8)

	// Generate the word
	for i := 1; i < maxLength; i++ {
		if i%2 == 0 {
			word += consonants[generateRanNum(0, 21)]
		} else {
			word += vowels[generateRanNum(0, 5)]
		}
	}

	//Marshal or convert word to json and write to response
	wordJson, err := json.Marshal(word)
	if err != nil {
		panic(err)
	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(wordJson)
}

// Define a default response for the home route.
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! This is an API written in GO! Try navigating to the /generate url.")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api", home)
	router.HandleFunc("/api/generate", generate)

	// Accept CORS requests
	handler := cors.Default().Handler(router)

	fmt.Println("Server listening!")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
