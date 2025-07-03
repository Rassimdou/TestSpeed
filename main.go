package main

import (
	"fmt"
)

func main() {

	fmt.Println("the server is running on port 8080")

	RunPingTest()
	RunDownlaodTest()
	RunUploadTest()

	server := NewServer()
	server.Start(":5000")

}
