package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func RunUploadTest() {
	//Generate 10MB dummy data
	data := make([]byte, 10*1024*1024) // 10 MB

	//Wrap data in a reader
	bodyReader := bytes.NewReader(data)

	start := time.Now()

	//Send POST request
	resp, err := http.Post("http://localhost:8080/upload", "application/octet-stream", bodyReader)
	if err != nil {
		fmt.Println("Error uploading data:", err)
		return
	}
	defer resp.Body.Close()

	// Discard response body
	io.Copy(io.Discard, resp.Body)

	elapsed := time.Since(start)

	// Calculate upload speed in Mbps
	uploadSpeed := (float64(len(data)) * 8) / (elapsed.Seconds() * 1_000_000)

	fmt.Printf("Upload Speed: %.2f Mbps\n", uploadSpeed)
}
