package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "World", "Specify a name to greet")
	flag.Parse()

	fmt.Printf("Hello, %s\n", name)
}
