package main

import (
	"context"
	"fmt"
	"net"

	"github.com/spf13/pflag"
)

func main() {
	var ns net.IP
	var domainName []string

	pflag.IPVarP(&ns, "nameserver", "n", net.IP{}, "Specify a name server IP address")
	pflag.StringArrayVarP(&domainName, "domainname", "d", []string{}, "Specify one or more domain names to resolve")

	pflag.Parse()

	dns := net.JoinHostPort(ns.String(), "53")

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return net.Dial(network, dns)
		},
	}

	for _, v := range domainName {
		ip, err := resolver.LookupHost(context.Background(), v)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Domain name %s resolved to %s using %s as the name server\n", v, ip[0], ns)
	}
}
