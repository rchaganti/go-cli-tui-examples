package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	// define flags
	hostname := flag.String("hostname", "", "Specify a hostname to check connectivity")
	timeout := flag.Duration("timeout", time.Second*10, "Specify timeout for the connectivity check")

	var port int
	flag.IntVar(&port, "port", 80, "Specify a port number to check connectivity")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <-hostname> [-port] [-timeout]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	// Parse the collected arguments into required values
	if *hostname == "" {
		fmt.Printf("--hostname argument must be provided with a non-empty value\n")
		os.Exit(1)
	}
	h := *hostname + ":" + strconv.Itoa(port)

	fmt.Printf("Trying to connect to: %s\n", h)
	c, err := net.DialTimeout("tcp", h, *timeout)
	if err != nil {
		fmt.Println("Error connecting to", h)
		os.Exit(1)
	}

	defer c.Close()

	fmt.Printf("Connected to %s with no errors\n", h)
}
