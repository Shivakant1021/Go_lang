package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func scanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 500*time.Millisecond)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run scanner.go <hostname>")
		return
	}

	hostname := os.Args[1]
	fmt.Println("Starting port scan on", hostname)

	for port := 1; port <= 1024; port++ {
		open := scanPort("tcp", hostname, port)
		if open {
			fmt.Printf("Port %d is open\n", port)
		}
	}

}

//go run .\6_port_scanner\main.go scanme.nmap.org
// Output:
// Starting port scan on scanme.nmap.org
// Port 22 is open
// Port 80 is open

//go run .\6_port_scanner\main.go globalindia.com
