package main

import "gocui-md-reader-list/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
