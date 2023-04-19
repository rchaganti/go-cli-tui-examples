package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	// Get Args supplied at the command-line
	args := os.Args[1:]

	if len(args) < 3 {
		log.Fatal(`
"You must supply three arguments.\n
Usage: ./webchk hostname portnumber timeout.\n
Example: ./webchk google.com 80 15s"
		`)
	}

	// Parse the collected arguments into required values
	hostname := args[0] + ":" + args[1]

	timeout, err := time.ParseDuration(args[2])
	if err != nil {
		fmt.Println("Error converting timeout argument to time.Duration", err)
		os.Exit(1)
	}

	fmt.Printf("Trying to connect to: %s\n", hostname)
	c, err := net.DialTimeout("tcp", hostname, timeout)
	if err != nil {
		fmt.Println("Error connecting to", hostname)
		os.Exit(1)
	}

	defer c.Close()

	fmt.Printf("Connected to %s with no errors\n", hostname)
}
