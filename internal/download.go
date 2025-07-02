package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func RunDownlaodTest() {

	start := time.Now()
	resp, err := http.Get("https://speed.hetzner.de/10MB.bin")
	if err != nil {
		fmt.Println("Error making testing download speed :", err)
	} else {
		n, _ := io.Copy(io.Discard, resp.Body)
		elapsed := time.Since(start)
		speedMbps := (float64(n) * 8) / (elapsed.Seconds() * 1_000_000)
		fmt.Printf("Download Speed: %.2f Mbps\n", speedMbps)
		defer resp.Body.Close()
	}

}
