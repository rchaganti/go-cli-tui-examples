package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "netutil COMMAND ARGUMENTS [OPTIONS]",
		Short: "Perform network and connection checks",
		Long:  "A CLI tool for performing network connection and DNS checks",
		Args:  cobra.NoArgs,
	}

	checkCmd = &cobra.Command{
		Use:   "check",
		Short: "Validate if an IP socket can be connected",
		Long:  "Validate if a provided IP address and port can be connected from the local machine",
		Args:  cobra.ExactArgs(1),
		RunE:  checkSocket,
	}

	resolveCmd = &cobra.Command{
		Use:   "resolve",
		Short: "Resolve a domain name to an IPv4 address",
		Long:  "Resolve a domain name to an IPv4 address and, optionally, using a custom DNS server IP address",
		Args:  cobra.ExactArgs(1),
		RunE:  resolveName,
	}
)

var port int
var nameserver net.IP

func checkSocket(cmd *cobra.Command, args []string) error {
	port, _ := cmd.Flags().GetInt("port")
	address := fmt.Sprintf("%s:%s", args[0], strconv.Itoa(port))
	timeout, _ := time.ParseDuration("10s")
	_, err := net.DialTimeout("tcp", "address", timeout)

	if err != nil {
		fmt.Printf("Network address %s cannot be reached!\n", address)
		return err
	} else {
		fmt.Printf("Network address %s can be reached!\n", address)
		return nil
	}
}

func resolveName(cmd *cobra.Command, args []string) error {
	nameserver, _ := cmd.Flags().GetIP("nameserver")
	dns := net.JoinHostPort(nameserver.String(), "53")

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return net.Dial(network, dns)
		},
	}

	ip, err := resolver.LookupHost(context.Background(), args[0])
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Printf("Domain name %s resolved to %s using %s as the name server\n", args[0], ip[0], nameserver)
		return nil
	}
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	checkCmd.Flags().IntVarP(&port, "port", "p", 80, "Port number to check")
	rootCmd.AddCommand(checkCmd)

	resolveCmd.Flags().IPVarP(&nameserver, "nameserver", "n", net.IPv4(8, 8, 8, 8), "Custom nameserver IP address")
	rootCmd.AddCommand(resolveCmd)
}
