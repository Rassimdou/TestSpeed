package internal

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func RunUploadTest() {
	//Generate 10MB dummy data
	fmt.Println("Starting upload speed test...")
	data := make([]byte, 50*1024*1024) // 50 MB

	//Wrap data in a reader
	bodyReader := bytes.NewReader(data)

	start := time.Now()
	//Send POST request
	resp, err := http.Post("https://speed.cloudflare.com/__up", "application/octet-stream", bodyReader)
	if err != nil {
		fmt.Println("Error uploading data:", err)
		return
	}
	defer resp.Body.Close()

	// Discard response body
	

	io.Copy(io.Discard, resp.Body)


	elapsed := time.Since(start)

	// Calculate upload speed in Mbps
	uploadSpeed := (float64(len(data)*8) ) / (elapsed.Seconds() * 1_000_000)

	fmt.Printf("Upload Speed: %.2f Mbps\n", uploadSpeed)
}
