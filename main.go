package main

import (
	"fmt"
	"net/http"

	"github.com/Rassimdou/TestSpeed/handlers"
)

func main() {

	http.HandleFunc("/upload", handlers.UploadHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
	RunPingTest()
	RunDownlaodTest()
	RunUploadTest()

}
