package az

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/spf13/cobra"
)

var subscriptionId = "5073fd4c-3a1b-4559-8371-21e034f70820"

func Info(cmd *cobra.Command, args []string) error {
	parent := *cmd.Parent()
	fmt.Println(parent.Name())

	return nil
}

func IsAzureAuthenticated() bool {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Authentication failure: %+v", err)
		return false
	} else {
		client, _ := armresources.NewClient(subscriptionId, cred, nil)
		log.Print("Authenticated to subscription", client)
		return true
	}
}

func GetRg(name string) []string {
	ctx := context.Background()
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	resourcesClientFactory, err := armresources.NewClientFactory(subscriptionId, cred, nil)
	if err != nil {
		log.Fatal(err)
	}
	resourceGroupClient := resourcesClientFactory.NewResourceGroupsClient()
}
