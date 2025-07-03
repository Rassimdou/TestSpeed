package main

import (
	"fmt"
	"net/http"

	"github.com/Rassimdou/TestSpeed/handlers"
	"github.com/Rassimdou/TestSpeed/internal"
)

func main() {
	// Start the server in a goroutine so it can handle uploads during the tests
	http.HandleFunc("/upload", handlers.UploadHandler)
	go func() {
		fmt.Println("Server is running on port 8080")
		http.ListenAndServe(":8080", nil)
	}()

	// Run tests
	internal.RunPingTest()
	internal.RunICMPPingTest()
	internal.RunDownlaodTest()
	internal.RunUploadTest()

	// Wait for user input to keep the program running (so server stays alive)
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
