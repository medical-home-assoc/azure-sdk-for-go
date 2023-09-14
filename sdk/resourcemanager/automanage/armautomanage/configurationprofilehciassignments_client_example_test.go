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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2dcad6d6e9a96882eb6d317e7500a94be007a9c6/specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/createOrUpdateConfigurationProfileHCIAssignment.json
func ExampleConfigurationProfileHCIAssignmentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomanage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationProfileHCIAssignmentsClient().CreateOrUpdate(ctx, "myResourceGroupName", "myClusterName", "default", armautomanage.ConfigurationProfileAssignment{
		Properties: &armautomanage.ConfigurationProfileAssignmentProperties{
			ConfigurationProfile: to.Ptr("/providers/Microsoft.Automanage/bestPractices/AzureBestPracticesProduction"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfigurationProfileAssignment = armautomanage.ConfigurationProfileAssignment{
	// 	Name: to.Ptr("default"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.AzureStackHci/clusters/myClusterName/providers/Microsoft.Automanage/AutomanageAssignments/default"),
	// 	Properties: &armautomanage.ConfigurationProfileAssignmentProperties{
	// 		ConfigurationProfile: to.Ptr("/providers/Microsoft.Automanage/bestPractices/AzureBestPracticesProduction"),
	// 		TargetID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/myResourceGroupName/providers/Microsoft.AzureStackHci/clusters/myClusterName"),
	// 	},
	// 	SystemData: &armautomanage.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.1075056Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1@outlook.com"),
	// 		CreatedByType: to.Ptr(armautomanage.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.1974346Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2@outlook.com"),
	// 		LastModifiedByType: to.Ptr(armautomanage.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2dcad6d6e9a96882eb6d317e7500a94be007a9c6/specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getConfigurationProfileHCIAssignment.json
func ExampleConfigurationProfileHCIAssignmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomanage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConfigurationProfileHCIAssignmentsClient().Get(ctx, "myResourceGroupName", "myClusterName", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConfigurationProfileAssignment = armautomanage.ConfigurationProfileAssignment{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Automanage/ConfigurationProfileAssignments"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/myResourceGroupName/Microsoft.AzureStackHci/clusters/machines/myClusterName/providers/Microsoft.Automanage/configurationProfileAssignments/default"),
	// 	Properties: &armautomanage.ConfigurationProfileAssignmentProperties{
	// 		ConfigurationProfile: to.Ptr("/providers/Microsoft.Automanage/bestPractices/AzureBestPracticesDevTest"),
	// 		Status: to.Ptr("Compliant"),
	// 		TargetID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/myResourceGroupName/Microsoft.AzureStackHci/clusters/machines/myClusterName"),
	// 	},
	// 	SystemData: &armautomanage.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.1075056Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1@outlook.com"),
	// 		CreatedByType: to.Ptr(armautomanage.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-04T02:03:01.1974346Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2@outlook.com"),
	// 		LastModifiedByType: to.Ptr(armautomanage.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2dcad6d6e9a96882eb6d317e7500a94be007a9c6/specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/deleteConfigurationProfileHCIAssignment.json
func ExampleConfigurationProfileHCIAssignmentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomanage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewConfigurationProfileHCIAssignmentsClient().Delete(ctx, "myResourceGroupName", "myClusterName", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
