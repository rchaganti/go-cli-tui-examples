package az

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// private functions
func newConnection() (cred azcore.TokenCredential, err error) {
	/*
		azlog.SetListener(func(event azlog.Event, s string) {
			fmt.Println(s)
		})

		azlog.SetEvents(azidentity.EventAuthentication)
	*/
	cred, err = azidentity.NewAzureCLICredential(nil)

	if err != nil {
		return nil, err
	}

	return cred, nil
}

func newRgClient(subscriptionId string) (rgClient *armresources.ResourceGroupsClient, err error) {
	cred, err := newConnection()
	if err != nil {
		log.Fatal("Unable to connect to Azure: %+v", err)
	}

	clientFactory, err := armresources.NewClientFactory(subscriptionId, cred, nil)
	if err != nil {
		return nil, err
	} else {
		rgClient = clientFactory.NewResourceGroupsClient()
		return rgClient, nil
	}
}

func newVmClient(subscriptionId string) (vmClient *armcompute.VirtualMachinesClient, err error) {
	cred, err := newConnection()
	if err != nil {
		log.Fatal("Unable to connect to Azure: %+v", err)
	}

	clientFactory, err := armcompute.NewClientFactory(subscriptionId, cred, nil)
	if err != nil {
		return nil, err
	} else {
		vmClient = clientFactory.NewVirtualMachinesClient()
		return vmClient, nil
	}
}

// exported functions
func ListResourceGroup(subscriptionId string) {
	rgClient, err := newRgClient(subscriptionId)

	if err != nil {
		log.Fatal(err)
	}
	resultPager := rgClient.NewListPager(nil)

	ctx := context.Background()
	resourceGroups := make([]*armresources.ResourceGroup, 0)
	for resultPager.More() {
		pageResp, err := resultPager.NextPage(ctx)
		if err != nil {
			log.Fatal(err)
		}
		resourceGroups = append(resourceGroups, pageResp.ResourceGroupListResult.Value...)
	}

	jsonData, err := json.MarshalIndent(resourceGroups, "\t", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}

func GetResourceGroup(subscriptionId, name string) {
	var resourceGroup *armresources.ResourceGroup
	rgClient, err := newRgClient(subscriptionId)

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	resourceGroupResp, err := rgClient.Get(ctx, name, nil)
	if err != nil {
		log.Fatal(err)
	}

	resourceGroup = &resourceGroupResp.ResourceGroup
	jsonData, err := json.MarshalIndent(*resourceGroup, "\t", "\t")
	if err != nil {
		log.Fatalf("Failed to marshal resource group to JSON: %v", err)
	}
	fmt.Println(string(jsonData))
}

func ListVirtualMachine(subscriptionId, resourceGroupName string) {
	vmClient, err := newVmClient(subscriptionId)

	if err != nil {
		log.Fatal(err)
	}

	vmPager := vmClient.NewListPager(resourceGroupName, nil)
	ctx := context.Background()
	virtualMachines := make([]*armcompute.VirtualMachine, 0)

	for vmPager.More() {
		page, err := vmPager.NextPage(ctx)
		if err != nil {
			log.Fatal(err)
		}

		virtualMachines = append(virtualMachines, page.VirtualMachineListResult.Value...)
	}

	jsonData, err := json.MarshalIndent(virtualMachines, "\t", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}

func GetVirtualMachine(subscriptionId, resourceGroupName, name string) {
	var virtualMachine *armcompute.VirtualMachine
	vmClient, err := newVmClient(subscriptionId)

	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	virtualMachineGetResp, err := vmClient.Get(ctx, resourceGroupName, name, nil)

	if err != nil {
		log.Fatal(err)
	}

	virtualMachine = &virtualMachineGetResp.VirtualMachine
	jsonData, err := json.MarshalIndent(*virtualMachine, "\t", "\t")
	if err != nil {
		log.Fatalf("Failed to marshal resource group to JSON: %v", err)
	}
	fmt.Println(string(jsonData))
}
