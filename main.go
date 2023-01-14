package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type RandomNumber struct {
	Number int `json:"number"`
}

func generateRandomNumber(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(100)
	randomNumber := RandomNumber{Number: number}

	json.NewEncoder(w).Encode(randomNumber)
}

func main() {
	http.HandleFunc("/random", generateRandomNumber)
	fmt.Println("Random number microservice started on port 3000")
	http.ListenAndServe(":3000", nil)
}
