package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've reached the todo app!")
	})

	fmt.Printf("Server started in port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
