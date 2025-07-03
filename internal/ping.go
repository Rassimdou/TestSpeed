package internal

import (
	"fmt"
	"net/http"
	"time"
)

func RunPingTest() {
	start := time.Now()
	_, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error making GET request:", err)
	}
	elapsed := time.Since(start)
	fmt.Println("Ping :", elapsed.Milliseconds(), "ms")
}
