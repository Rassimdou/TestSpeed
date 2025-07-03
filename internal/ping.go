package internal

import (
	"fmt"
	
	"time"

	"github.com/go-ping/ping"
)


// ICMP ping using github.com/go-ping/ping
func RunICMPPingTest() {
	fmt.Println("Starting ICMP Ping (like speedtest.net)...")

	pinger, err := ping.NewPinger("8.8.8.8") // Use a reliable server, like Google DNS
	if err != nil {
		fmt.Println("Error creating pinger:", err)
		return
	}
	pinger.Count = 3 // Send 3 packets
	pinger.Timeout = time.Second * 5
	pinger.SetPrivileged(true) // Required for Windows
	fmt.Println("Pinging 8.8.8.8 with ICMP...")
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		fmt.Println("Ping error:", err)
		return
	}
	stats := pinger.Statistics()
	fmt.Printf("ICMP Ping ms: min=%v avg=%v max=%v\n", stats.MinRtt.Milliseconds(), stats.AvgRtt.Milliseconds(), stats.MaxRtt.Milliseconds())

}
