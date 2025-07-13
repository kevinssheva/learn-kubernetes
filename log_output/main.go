package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func generateRandomString(n int, r *rand.Rand) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func randomStringHandler(w http.ResponseWriter, r *http.Request) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomString := generateRandomString(16, s)
	timestamp := time.Now().Format(time.RFC3339)
	fmt.Fprintf(w, "[%s] %s\n", timestamp, randomString)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/random", randomStringHandler)

	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
