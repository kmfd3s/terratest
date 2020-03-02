package azure

import (
	"context"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
	"github.com/gruntwork-io/terratest/modules/shell"
)

// CreateResourceGroup create resource group
func CreateResourceGroup(t *testing.T, name string, location string) {
	groupsClient := getResourceGroupClient(t)
	groupsClient.CreateOrUpdate(context.Background(), name, resources.Group{
		Location: &location,
	})
}

// DeleteResourceGroup delete resource group
func DeleteResourceGroup(t *testing.T, name string) {
	groupsClient := getResourceGroupClient(t)
	groupsClient.Delete(context.Background(), name)
}

func getResourceGroupClient(t *testing.T) *resources.GroupsClient {
	subscriptionID := getSubscriptionID(t)
	groupsClient := resources.NewGroupsClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		t.Fatal(err)
	}

	groupsClient.Authorizer = *authorizer
	return &groupsClient
}

func getSubscriptionID(t *testing.T) string {
	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")

	if len(subscriptionID) == 0 {
		subscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")
	}

	if len(subscriptionID) == 0 {
		subscriptionID = shell.RunCommandAndGetOutput(t, shell.Command{
			Command: "az",
			Args: []string{
				"account",
				"show",
				"--query", "id",
				"-o", "tsv",
			}})
	}

	return subscriptionID
}
