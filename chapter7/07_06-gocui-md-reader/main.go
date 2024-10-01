package main

import "gocui-md-reader/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
