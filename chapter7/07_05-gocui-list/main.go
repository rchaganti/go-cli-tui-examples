package main

import "gocui-list/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
