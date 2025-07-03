package internal

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func RunDownlaodTest() {

	fmt.Println("Starting download speed test...")
	start := time.Now()
	
	resp, err := http.Get("https://speed.cloudflare.com/__down?bytes=100000000") // Valid test file
	if err != nil {
		fmt.Println("Error making testing download speed :", err)
	} 
	defer resp.Body.Close()

	//copy the body to discard it but count the bytes
	bytesDownloaded, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	elapsed:= time.Since(start)
	fmt.Printf("Downloaded %d bytes in %.2f seconds\n", bytesDownloaded, elapsed.Seconds())
	speedMbps := (float64(bytesDownloaded) * 8) / (elapsed.Seconds() * 1_000_000)
	fmt.Printf("Download Speed: %.2f Mbps\n", speedMbps)
}
