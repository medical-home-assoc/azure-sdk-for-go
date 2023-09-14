//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/ScopeMapList.json
func ExampleScopeMapsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScopeMapsClient().NewListPager("myResourceGroup", "myRegistry", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ScopeMapListResult = armcontainerregistry.ScopeMapListResult{
		// 	Value: []*armcontainerregistry.ScopeMap{
		// 		{
		// 			Name: to.Ptr("myScopeMap"),
		// 			Type: to.Ptr("Microsoft.ContainerRegistry/registries/scopeMaps"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/scopeMaps/myScopeMap"),
		// 			Properties: &armcontainerregistry.ScopeMapProperties{
		// 				Type: to.Ptr("IsUserDefined"),
		// 				Actions: []*string{
		// 					to.Ptr("repositories/myrepository/contentWrite"),
		// 					to.Ptr("repositories/myrepository/delete")},
		// 					CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T23:14:37.0707808Z"); return t}()),
		// 					ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/ScopeMapGet.json
func ExampleScopeMapsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScopeMapsClient().Get(ctx, "myResourceGroup", "myRegistry", "myScopeMap", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScopeMap = armcontainerregistry.ScopeMap{
	// 	Name: to.Ptr("myScopeMap"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/scopeMaps"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/scopeMaps/myScopeMap"),
	// 	Properties: &armcontainerregistry.ScopeMapProperties{
	// 		Type: to.Ptr("IsUserDefined"),
	// 		Actions: []*string{
	// 			to.Ptr("repositories/myrepository/contentWrite"),
	// 			to.Ptr("repositories/myrepository/delete")},
	// 			CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T23:14:37.0707808Z"); return t}()),
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/ScopeMapCreate.json
func ExampleScopeMapsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewScopeMapsClient().BeginCreate(ctx, "myResourceGroup", "myRegistry", "myScopeMap", armcontainerregistry.ScopeMap{
		Properties: &armcontainerregistry.ScopeMapProperties{
			Description: to.Ptr("Developer Scopes"),
			Actions: []*string{
				to.Ptr("repositories/myrepository/contentWrite"),
				to.Ptr("repositories/myrepository/delete")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScopeMap = armcontainerregistry.ScopeMap{
	// 	Name: to.Ptr("myScopeMap"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/scopeMaps"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/scopeMaps/myScopeMap"),
	// 	Properties: &armcontainerregistry.ScopeMapProperties{
	// 		Type: to.Ptr("IsUserDefined"),
	// 		Actions: []*string{
	// 			to.Ptr("repositories/myrepository/contentWrite"),
	// 			to.Ptr("repositories/myrepository/delete")},
	// 			CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T23:14:37.0707808Z"); return t}()),
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/ScopeMapDelete.json
func ExampleScopeMapsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewScopeMapsClient().BeginDelete(ctx, "myResourceGroup", "myRegistry", "myScopeMap", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-01-01-preview/examples/ScopeMapUpdate.json
func ExampleScopeMapsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewScopeMapsClient().BeginUpdate(ctx, "myResourceGroup", "myRegistry", "myScopeMap", armcontainerregistry.ScopeMapUpdateParameters{
		Properties: &armcontainerregistry.ScopeMapPropertiesUpdateParameters{
			Description: to.Ptr("Developer Scopes"),
			Actions: []*string{
				to.Ptr("repositories/myrepository/contentWrite"),
				to.Ptr("repositories/myrepository/contentRead")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScopeMap = armcontainerregistry.ScopeMap{
	// 	Name: to.Ptr("myScopeMap"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/scopeMaps"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/scopeMaps/myScopeMap"),
	// 	Properties: &armcontainerregistry.ScopeMapProperties{
	// 		Type: to.Ptr("IsUserDefined"),
	// 		Actions: []*string{
	// 			to.Ptr("repositories/myrepository/contentWrite"),
	// 			to.Ptr("repositories/myrepository/contentRead")},
	// 			CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T23:14:37.0707808Z"); return t}()),
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		},
	// 	}
}
