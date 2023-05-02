package main

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/spf13/pflag"
)

func main() {
	var ip net.IP
	var port int

	pflag.IPVarP(&ip, "ipAddress", "i", net.IPv4(192, 168, 1, 1), "Specify an IP address to validate")
	pflag.IntVarP(&port, "port", "p", 80, "Specify a port to check connectivity")
	pflag.Parse()

	address := fmt.Sprintf("%s:%s", ip.String(), strconv.Itoa(port))
	timeout, _ := time.ParseDuration("10s")
	_, err := net.DialTimeout("ipv4::Tcp", address, timeout)

	if err != nil {
		fmt.Printf("Network address %s cannot be reached!\n", address)
	} else {
		fmt.Printf("Network address %s is available!\n", address)
	}
}
