//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armautomanage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2dcad6d6e9a96882eb6d317e7500a94be007a9c6/specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listServicePrincipalBySubscription.json
func ExampleServicePrincipalsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomanage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServicePrincipalsClient().NewListBySubscriptionPager(nil)
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
		// page.ServicePrincipalListResult = armautomanage.ServicePrincipalListResult{
		// 	Value: []*armautomanage.ServicePrincipal{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Automanage/ConfigurationProfileAssignments"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/providers/Microsoft.Automanage/servicePrincipals/default"),
		// 			Properties: &armautomanage.ServicePrincipalProperties{
		// 				AuthorizationSet: to.Ptr(true),
		// 				ServicePrincipalID: to.Ptr("<servicePrincipalId>"),
		// 			},
		// 			SystemData: &armautomanage.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T00:00:00.00Z"); return t}()),
		// 				CreatedBy: to.Ptr("SYSTEM"),
		// 				CreatedByType: to.Ptr(armautomanage.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T00:00:00.00Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("SYSTEM"),
		// 				LastModifiedByType: to.Ptr(armautomanage.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2dcad6d6e9a96882eb6d317e7500a94be007a9c6/specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getServicePrincipal.json
func ExampleServicePrincipalsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomanage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServicePrincipalsClient().Get(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServicePrincipal = armautomanage.ServicePrincipal{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Automanage/ConfigurationProfileAssignments"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/providers/Microsoft.Automanage/servicePrincipals/default"),
	// 	Properties: &armautomanage.ServicePrincipalProperties{
	// 		AuthorizationSet: to.Ptr(true),
	// 		ServicePrincipalID: to.Ptr("<servicePrincipalId>"),
	// 	},
	// 	SystemData: &armautomanage.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T00:00:00.00Z"); return t}()),
	// 		CreatedBy: to.Ptr("SYSTEM"),
	// 		CreatedByType: to.Ptr(armautomanage.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T00:00:00.00Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("SYSTEM"),
	// 		LastModifiedByType: to.Ptr(armautomanage.CreatedByTypeUser),
	// 	},
	// }
}
