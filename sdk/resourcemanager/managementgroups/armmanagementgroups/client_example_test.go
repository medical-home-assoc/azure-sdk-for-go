//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagementgroups_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/ListManagementGroups.json
func ExampleClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListPager(&armmanagementgroups.ClientListOptions{CacheControl: to.Ptr("no-cache"),
		Skiptoken: nil,
	})
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
		// page.ManagementGroupListResult = armmanagementgroups.ManagementGroupListResult{
		// 	Value: []*armmanagementgroups.ManagementGroupInfo{
		// 		{
		// 			Name: to.Ptr("20000000-0001-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Management/managementGroups"),
		// 			ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0001-0000-0000-000000000000"),
		// 			Properties: &armmanagementgroups.ManagementGroupInfoProperties{
		// 				DisplayName: to.Ptr("Group 1 Tenant 2"),
		// 				TenantID: to.Ptr("20000000-0000-0000-0000-000000000000"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("20000000-0004-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Management/managementGroups"),
		// 			ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0004-0000-0000-000000000000"),
		// 			Properties: &armmanagementgroups.ManagementGroupInfoProperties{
		// 				DisplayName: to.Ptr("Group 4 Tenant 2"),
		// 				TenantID: to.Ptr("20000000-0000-0000-0000-000000000000"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetManagementGroup.json
func ExampleClient_Get_getManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Get(ctx, "20000000-0001-0000-0000-000000000000", &armmanagementgroups.ClientGetOptions{Expand: nil,
		Recurse:      nil,
		Filter:       nil,
		CacheControl: to.Ptr("no-cache"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagementGroup = armmanagementgroups.ManagementGroup{
	// 	Name: to.Ptr("20000000-0001-0000-0000-000000000000"),
	// 	Type: to.Ptr("Microsoft.Management/managementGroups"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0001-0000-0000-000000000000"),
	// 	Properties: &armmanagementgroups.ManagementGroupProperties{
	// 		DisplayName: to.Ptr("Group 1 Tenant 2"),
	// 		TenantID: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 		Details: &armmanagementgroups.ManagementGroupDetails{
	// 			Parent: &armmanagementgroups.ParentGroupInfo{
	// 				Name: to.Ptr("RootGroup"),
	// 				DisplayName: to.Ptr("RootGroup"),
	// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/RootGroup"),
	// 			},
	// 			UpdatedBy: to.Ptr("16b8ef21-5c9f-420c-bcc9-e4f8c9f30b4b"),
	// 			UpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00.00Z"); return t}()),
	// 			Version: to.Ptr[int32](1),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetManagementGroupWithAncestors.json
func ExampleClient_Get_getManagementGroupWithAncestors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Get(ctx, "20000000-0001-0000-0000-00000000000", &armmanagementgroups.ClientGetOptions{Expand: to.Ptr(armmanagementgroups.ManagementGroupExpandTypeAncestors),
		Recurse:      nil,
		Filter:       nil,
		CacheControl: to.Ptr("no-cache"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagementGroup = armmanagementgroups.ManagementGroup{
	// 	Name: to.Ptr("20000000-0001-0000-0000-000000000000"),
	// 	Type: to.Ptr("Microsoft.Management/managementGroups"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0001-0000-0000-000000000000"),
	// 	Properties: &armmanagementgroups.ManagementGroupProperties{
	// 		DisplayName: to.Ptr("Group 1 Tenant 2"),
	// 		TenantID: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 		Details: &armmanagementgroups.ManagementGroupDetails{
	// 			ManagementGroupAncestorsChain: []*armmanagementgroups.ManagementGroupPathElement{
	// 				{
	// 					Name: to.Ptr("20000000-0000-0000-0000-000000000001"),
	// 					DisplayName: to.Ptr("Parent display name"),
	// 				},
	// 				{
	// 					Name: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 					DisplayName: to.Ptr("Grandparent display name"),
	// 			}},
	// 			Parent: &armmanagementgroups.ParentGroupInfo{
	// 				Name: to.Ptr("20000000-0000-0000-0000-000000000001"),
	// 				DisplayName: to.Ptr("Parent display name"),
	// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0000-0000-0000-000000000001"),
	// 			},
	// 			UpdatedBy: to.Ptr("Test"),
	// 			UpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00.00Z"); return t}()),
	// 			Version: to.Ptr[int32](1),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetManagementGroupWithExpand.json
func ExampleClient_Get_getManagementGroupWithExpand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Get(ctx, "20000000-0001-0000-0000-000000000000", &armmanagementgroups.ClientGetOptions{Expand: to.Ptr(armmanagementgroups.ManagementGroupExpandTypeChildren),
		Recurse:      nil,
		Filter:       nil,
		CacheControl: to.Ptr("no-cache"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagementGroup = armmanagementgroups.ManagementGroup{
	// 	Name: to.Ptr("20000000-0001-0000-0000-000000000000"),
	// 	Type: to.Ptr("Microsoft.Management/managementGroups"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0001-0000-0000-000000000000"),
	// 	Properties: &armmanagementgroups.ManagementGroupProperties{
	// 		Children: []*armmanagementgroups.ManagementGroupChildInfo{
	// 			{
	// 				Name: to.Ptr("20000000-0002-0000-0000-000000000000"),
	// 				Type: to.Ptr(armmanagementgroups.ManagementGroupChildTypeMicrosoftManagementManagementGroups),
	// 				DisplayName: to.Ptr("Group 2 Tenant 2"),
	// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0002-0000-0000-000000000000"),
	// 			},
	// 			{
	// 				Name: to.Ptr("20000000-0003-0000-0000-000000000000"),
	// 				Type: to.Ptr(armmanagementgroups.ManagementGroupChildTypeMicrosoftManagementManagementGroups),
	// 				DisplayName: to.Ptr("Group 3 Tenant 2"),
	// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0003-0000-0000-000000000000"),
	// 			},
	// 			{
	// 				Name: to.Ptr("10000000-F004-0000-0000-000000000000"),
	// 				Type: to.Ptr(armmanagementgroups.ManagementGroupChildTypeSubscriptions),
	// 				DisplayName: to.Ptr("Subscription 4 Tenant 1"),
	// 				ID: to.Ptr("/subscriptions/10000000-F004-0000-0000-000000000000"),
	// 			},
	// 			{
	// 				Name: to.Ptr("20000000-F005-0000-0000-000000000000"),
	// 				Type: to.Ptr(armmanagementgroups.ManagementGroupChildTypeSubscriptions),
	// 				DisplayName: to.Ptr("Subscription 5 Tenant 2"),
	// 				ID: to.Ptr("/subscriptions/20000000-F005-0000-0000-000000000000"),
	// 			},
	// 			{
	// 				Name: to.Ptr("30000000-F003-0000-0000-000000000000"),
	// 				Type: to.Ptr(armmanagementgroups.ManagementGroupChildTypeSubscriptions),
	// 				DisplayName: to.Ptr("Subscription 3 Tenant 3"),
	// 				ID: to.Ptr("/subscriptions/30000000-F003-0000-0000-000000000000"),
	// 		}},
	// 		DisplayName: to.Ptr("Group 1 Tenant 2"),
	// 		TenantID: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 		Details: &armmanagementgroups.ManagementGroupDetails{
	// 			Parent: &armmanagementgroups.ParentGroupInfo{
	// 				Name: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 				DisplayName: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0000-0000-0000-000000000000"),
	// 			},
	// 			UpdatedBy: to.Ptr("Test"),
	// 			UpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00.00Z"); return t}()),
	// 			Version: to.Ptr[int32](1),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetManagementGroupWithPath.json
func ExampleClient_Get_getManagementGroupWithPath() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Get(ctx, "20000000-0001-0000-0000-000000000000", &armmanagementgroups.ClientGetOptions{Expand: to.Ptr(armmanagementgroups.ManagementGroupExpandTypePath),
		Recurse:      nil,
		Filter:       nil,
		CacheControl: to.Ptr("no-cache"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagementGroup = armmanagementgroups.ManagementGroup{
	// 	Name: to.Ptr("20000000-0001-0000-0000-000000000000"),
	// 	Type: to.Ptr("Microsoft.Management/managementGroups"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0001-0000-0000-000000000000"),
	// 	Properties: &armmanagementgroups.ManagementGroupProperties{
	// 		DisplayName: to.Ptr("Group 1 Tenant 2"),
	// 		TenantID: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 		Details: &armmanagementgroups.ManagementGroupDetails{
	// 			Path: []*armmanagementgroups.ManagementGroupPathElement{
	// 				{
	// 					Name: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 					DisplayName: to.Ptr("Grandparent display name"),
	// 				},
	// 				{
	// 					Name: to.Ptr("20000000-0000-0000-0000-000000000001"),
	// 					DisplayName: to.Ptr("Parent display name"),
	// 			}},
	// 			Parent: &armmanagementgroups.ParentGroupInfo{
	// 				Name: to.Ptr("20000000-0000-0000-0000-000000000001"),
	// 				DisplayName: to.Ptr("Parent display name"),
	// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0000-0000-0000-000000000001"),
	// 			},
	// 			UpdatedBy: to.Ptr("Test"),
	// 			UpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00.00Z"); return t}()),
	// 			Version: to.Ptr[int32](1),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetManagementGroupWithExpandAndRecurse.json
func ExampleClient_Get_getManagementGroupsWithExpandAndRecurse() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Get(ctx, "20000000-0001-0000-0000-000000000000", &armmanagementgroups.ClientGetOptions{Expand: to.Ptr(armmanagementgroups.ManagementGroupExpandTypeChildren),
		Recurse:      to.Ptr(true),
		Filter:       nil,
		CacheControl: to.Ptr("no-cache"),
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagementGroup = armmanagementgroups.ManagementGroup{
	// 	Name: to.Ptr("RootGroup"),
	// 	Type: to.Ptr("Microsoft.Management/managementGroups"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/RootGroup"),
	// 	Properties: &armmanagementgroups.ManagementGroupProperties{
	// 		Children: []*armmanagementgroups.ManagementGroupChildInfo{
	// 			{
	// 				Name: to.Ptr("Child"),
	// 				Type: to.Ptr(armmanagementgroups.ManagementGroupChildTypeMicrosoftManagementManagementGroups),
	// 				Children: []*armmanagementgroups.ManagementGroupChildInfo{
	// 					{
	// 						Name: to.Ptr("Leaf"),
	// 						Type: to.Ptr(armmanagementgroups.ManagementGroupChildTypeMicrosoftManagementManagementGroups),
	// 						Children: []*armmanagementgroups.ManagementGroupChildInfo{
	// 							{
	// 								Name: to.Ptr("728bcbe4-8d56-4510-86c2-4921b8beefbc"),
	// 								Type: to.Ptr(armmanagementgroups.ManagementGroupChildTypeSubscriptions),
	// 								DisplayName: to.Ptr("Pay-As-You-Go"),
	// 								ID: to.Ptr("/subscriptions/728bcbe4-8d56-4510-86c2-4921b8beefbc"),
	// 						}},
	// 						DisplayName: to.Ptr("Leaf"),
	// 						ID: to.Ptr("/providers/Microsoft.Management/managementGroups/Leaf"),
	// 				}},
	// 				DisplayName: to.Ptr("Child"),
	// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/Child"),
	// 			},
	// 			{
	// 				Name: to.Ptr("AnotherChild"),
	// 				Type: to.Ptr(armmanagementgroups.ManagementGroupChildTypeMicrosoftManagementManagementGroups),
	// 				DisplayName: to.Ptr("Leaf"),
	// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/AnotherChild"),
	// 		}},
	// 		DisplayName: to.Ptr("RootGroup"),
	// 		TenantID: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 		Details: &armmanagementgroups.ManagementGroupDetails{
	// 			Parent: &armmanagementgroups.ParentGroupInfo{
	// 				Name: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 				DisplayName: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0000-0000-0000-000000000000"),
	// 			},
	// 			UpdatedBy: to.Ptr("bd490e30-04cb-433e-b8c8-6066959a8bab"),
	// 			UpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T02:26:49.0022093Z"); return t}()),
	// 			Version: to.Ptr[int32](2),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/PutManagementGroup.json
func ExampleClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginCreateOrUpdate(ctx, "ChildGroup", armmanagementgroups.CreateManagementGroupRequest{
		Properties: &armmanagementgroups.CreateManagementGroupProperties{
			DisplayName: to.Ptr("ChildGroup"),
			Details: &armmanagementgroups.CreateManagementGroupDetails{
				Parent: &armmanagementgroups.CreateParentGroupInfo{
					ID: to.Ptr("/providers/Microsoft.Management/managementGroups/RootGroup"),
				},
			},
		},
	}, &armmanagementgroups.ClientBeginCreateOrUpdateOptions{CacheControl: to.Ptr("no-cache")})
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
	// res.ManagementGroup = armmanagementgroups.ManagementGroup{
	// 	Name: to.Ptr("ChildGroup"),
	// 	Type: to.Ptr("Microsoft.Management/managementGroups"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/ChildGroup"),
	// 	Properties: &armmanagementgroups.ManagementGroupProperties{
	// 		DisplayName: to.Ptr("ChildGroup"),
	// 		TenantID: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 		Details: &armmanagementgroups.ManagementGroupDetails{
	// 			Parent: &armmanagementgroups.ParentGroupInfo{
	// 				Name: to.Ptr("RootGroup"),
	// 				DisplayName: to.Ptr("RootGroup"),
	// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/RootGroup"),
	// 			},
	// 			UpdatedBy: to.Ptr("16b8ef21-5c9f-420c-bcc9-e4f8c9f30b4b"),
	// 			UpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00.00Z"); return t}()),
	// 			Version: to.Ptr[int32](1),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/PatchManagementGroup.json
func ExampleClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Update(ctx, "ChildGroup", armmanagementgroups.PatchManagementGroupRequest{
		DisplayName:   to.Ptr("AlternateDisplayName"),
		ParentGroupID: to.Ptr("/providers/Microsoft.Management/managementGroups/AlternateRootGroup"),
	}, &armmanagementgroups.ClientUpdateOptions{CacheControl: to.Ptr("no-cache")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagementGroup = armmanagementgroups.ManagementGroup{
	// 	Name: to.Ptr("ChildGroup"),
	// 	Type: to.Ptr("Microsoft.Management/managementGroups"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementGroups/ChildGroup"),
	// 	Properties: &armmanagementgroups.ManagementGroupProperties{
	// 		DisplayName: to.Ptr("AlternateDisplayName"),
	// 		TenantID: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 		Details: &armmanagementgroups.ManagementGroupDetails{
	// 			Parent: &armmanagementgroups.ParentGroupInfo{
	// 				Name: to.Ptr("AlternateRootGroup"),
	// 				DisplayName: to.Ptr("AlternateRootGroup"),
	// 				ID: to.Ptr("/providers/Microsoft.Management/managementGroups/AlternateRootGroup"),
	// 			},
	// 			UpdatedBy: to.Ptr("bd490e30-04cb-433e-b8c8-6066959a8bab"),
	// 			UpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T02:46:59.0545645Z"); return t}()),
	// 			Version: to.Ptr[int32](2),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/DeleteManagementGroup.json
func ExampleClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginDelete(ctx, "GroupToDelete", &armmanagementgroups.ClientBeginDeleteOptions{CacheControl: to.Ptr("no-cache")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/GetDescendants.json
func ExampleClient_NewGetDescendantsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagementgroups.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewGetDescendantsPager("20000000-0000-0000-0000-000000000000", &armmanagementgroups.ClientGetDescendantsOptions{Skiptoken: nil,
		Top: nil,
	})
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
		// page.DescendantListResult = armmanagementgroups.DescendantListResult{
		// 	Value: []*armmanagementgroups.DescendantInfo{
		// 		{
		// 			Name: to.Ptr("20000000-0001-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Management/managementGroups"),
		// 			ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0001-0000-0000-000000000000"),
		// 			Properties: &armmanagementgroups.DescendantInfoProperties{
		// 				DisplayName: to.Ptr("Group 1"),
		// 				Parent: &armmanagementgroups.DescendantParentGroupInfo{
		// 					ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0000-0000-0000-000000000000"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("20000000-0004-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Management/managementGroups/subscriptions"),
		// 			ID: to.Ptr("/subscriptions/20000000-0004-0000-0000-000000000000"),
		// 			Properties: &armmanagementgroups.DescendantInfoProperties{
		// 				DisplayName: to.Ptr("Subscription 4"),
		// 				Parent: &armmanagementgroups.DescendantParentGroupInfo{
		// 					ID: to.Ptr("/providers/Microsoft.Management/managementGroups/20000000-0000-0000-0000-000000000000"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
