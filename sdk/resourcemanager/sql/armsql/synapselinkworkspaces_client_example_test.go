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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2022-05-01-preview/examples/SynapseLinkWorkspaceListByDatabase.json
func ExampleSynapseLinkWorkspacesClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSynapseLinkWorkspacesClient().NewListByDatabasePager("Default-SQL-SouthEastAsia", "testsvr", "dbSynapse", nil)
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
		// page.SynapseLinkWorkspaceListResult = armsql.SynapseLinkWorkspaceListResult{
		// 	Value: []*armsql.SynapseLinkWorkspace{
		// 		{
		// 			Name: to.Ptr("current"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/linkWorkspaces"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/dbSynapse/linkWorkspaces/current"),
		// 			Properties: &armsql.SynapseLinkWorkspaceProperties{
		// 				Workspaces: []*armsql.SynapseLinkWorkspaceInfoProperties{
		// 					{
		// 						LinkConnectionName: to.Ptr("buildlink-cloud1"),
		// 						WorkspaceID: to.Ptr("/subcriptions/00000000-1111-2222-3333-444444444444/resourcegroups/Default-SQL-SouthEastAsia/providers/Microsoft.Synapse/workspaces/workspace1"),
		// 					},
		// 					{
		// 						LinkConnectionName: to.Ptr("buildlink-cloud2"),
		// 						WorkspaceID: to.Ptr("/subcriptions/00000000-1111-2222-3333-444444444444/resourcegroups/Default-SQL-SouthEastAsia/providers/Microsoft.Synapse/workspaces/workspace2"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
