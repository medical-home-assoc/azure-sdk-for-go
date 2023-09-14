//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/CancelDatabaseOperation.json
func ExampleDatabaseOperationsClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDatabaseOperationsClient().Cancel(ctx, "sqlcrudtest-7398", "sqlcrudtest-6661", "testdb", "f779414b-e748-4925-8cfe-c8598f7660ae", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/ListDatabaseOperations.json
func ExampleDatabaseOperationsClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabaseOperationsClient().NewListByDatabasePager("sqlcrudtest-7398", "sqlcrudtest-4645", "testdb", nil)
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
		// page.DatabaseOperationListResult = armsql.DatabaseOperationListResult{
		// 	Value: []*armsql.DatabaseOperation{
		// 		{
		// 			Name: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/servers/sqlcrudtest-4645/databases/testdb/operations/11111111-1111-1111-1111-111111111111"),
		// 			Properties: &armsql.DatabaseOperationProperties{
		// 				DatabaseName: to.Ptr("testdb"),
		// 				Operation: to.Ptr("UpdateLogicalDatabase"),
		// 				OperationFriendlyName: to.Ptr("ALTER DATABASE"),
		// 				PercentComplete: to.Ptr[int32](100),
		// 				ServerName: to.Ptr("sqlcrudtest-4645"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-01T09:10:08.1Z"); return t}()),
		// 				State: to.Ptr(armsql.ManagementOperationStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("55555555-5555-5555-5555-555555555555"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/servers/sqlcrudtest-4645/databases/testdb/operations/55555555-5555-5555-5555-555555555555"),
		// 			Properties: &armsql.DatabaseOperationProperties{
		// 				DatabaseName: to.Ptr("testdb"),
		// 				Operation: to.Ptr("UpdateLogicalDatabase"),
		// 				OperationFriendlyName: to.Ptr("ALTER DATABASE"),
		// 				PercentComplete: to.Ptr[int32](19),
		// 				ServerName: to.Ptr("sqlcrudtest-4645"),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-01T10:10:08.1Z"); return t}()),
		// 				State: to.Ptr(armsql.ManagementOperationStateInProgress),
		// 			},
		// 	}},
		// }
	}
}
