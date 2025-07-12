package main

import (
	"fmt"
	"math/rand"
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

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomString := generateRandomString(16, r)
	for {
		timestamp := time.Now().Format(time.RFC3339)
		fmt.Printf("[%s] %s\n", timestamp, randomString)
		time.Sleep(5 * time.Second)
	}
}
