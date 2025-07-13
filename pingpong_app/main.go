package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func pongHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++
	fmt.Fprintf(w, "pong %d\n", counter-1)
	mu.Unlock()
}

func main() {
	http.HandleFunc("/", pongHandler)
	log.Println("PingPong app listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
