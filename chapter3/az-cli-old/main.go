package main

import (
	_ "az-cli/cmd/get"
	_ "az-cli/cmd/list"
	"az-cli/pkg/az"
	"fmt"
)

func main() {
	auth := az.IsAzureAuthenticated()
	if auth {
		fmt.Println("Authenticated")
	} else {
		fmt.Println("Not authenticated")
	}

}
