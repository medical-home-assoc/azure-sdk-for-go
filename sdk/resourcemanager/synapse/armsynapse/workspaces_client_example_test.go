//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListWorkspacesInResourceGroup.json
func ExampleWorkspacesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListByResourceGroupPager("resourceGroup1", nil)
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
		// page.WorkspaceInfoListResult = armsynapse.WorkspaceInfoListResult{
		// 	Value: []*armsynapse.Workspace{
		// 		{
		// 			Name: to.Ptr("workspace1"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace1"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Properties: &armsynapse.WorkspaceProperties{
		// 				ConnectivityEndpoints: map[string]*string{
		// 					"dev": to.Ptr("workspace1.dev.projectarcadia.net"),
		// 					"sql": to.Ptr("workspace1.sql.projectarcadia.net"),
		// 				},
		// 				DefaultDataLakeStorage: &armsynapse.DataLakeStorageAccountDetails{
		// 					AccountURL: to.Ptr("https://accountname.dfs.core.windows.net"),
		// 					Filesystem: to.Ptr("default"),
		// 				},
		// 				ManagedResourceGroupName: to.Ptr("resourceGroup2"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SQLAdministratorLogin: to.Ptr("login"),
		// 				SQLAdministratorLoginPassword: to.Ptr("password"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("workspace2"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace2"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Properties: &armsynapse.WorkspaceProperties{
		// 				ConnectivityEndpoints: map[string]*string{
		// 					"dev": to.Ptr("workspace2.dev.projectarcadia.net"),
		// 					"sql": to.Ptr("workspace2.sql.projectarcadia.net"),
		// 				},
		// 				DefaultDataLakeStorage: &armsynapse.DataLakeStorageAccountDetails{
		// 					AccountURL: to.Ptr("https://accountname.dfs.core.windows.net"),
		// 					Filesystem: to.Ptr("default"),
		// 				},
		// 				ManagedResourceGroupName: to.Ptr("resourceGroup2"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SQLAdministratorLogin: to.Ptr("login"),
		// 				SQLAdministratorLoginPassword: to.Ptr("password"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspace.json
func ExampleWorkspacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().Get(ctx, "resourceGroup1", "workspace1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Workspace = armsynapse.Workspace{
	// 	Name: to.Ptr("workspace1"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace1"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armsynapse.WorkspaceProperties{
	// 		ConnectivityEndpoints: map[string]*string{
	// 			"dev": to.Ptr("workspace1.dev.projectarcadia.net"),
	// 			"sql": to.Ptr("workspace1.sql.projectarcadia.net"),
	// 		},
	// 		CspWorkspaceAdminProperties: &armsynapse.CspWorkspaceAdminProperties{
	// 			InitialWorkspaceAdminObjectID: to.Ptr("6c20646f-8050-49ec-b3b1-80a0e58e454d"),
	// 		},
	// 		DefaultDataLakeStorage: &armsynapse.DataLakeStorageAccountDetails{
	// 			AccountURL: to.Ptr("https://accountname.dfs.core.windows.net"),
	// 			Filesystem: to.Ptr("default"),
	// 		},
	// 		ExtraProperties: map[string]any{
	// 			"IsScopeEnabled": "false",
	// 			"WorkspaceType": "Normal",
	// 		},
	// 		ManagedResourceGroupName: to.Ptr("resourceGroup2"),
	// 		ManagedVirtualNetworkSettings: &armsynapse.ManagedVirtualNetworkSettings{
	// 			AllowedAADTenantIDsForLinking: []*string{
	// 				to.Ptr("740239CE-A25B-485B-86A0-262F29F6EBDB")},
	// 				LinkedAccessCheckOnTargetResource: to.Ptr(false),
	// 				PreventDataExfiltration: to.Ptr(false),
	// 			},
	// 			PrivateEndpointConnections: []*armsynapse.PrivateEndpointConnection{
	// 				{
	// 					Name: to.Ptr("sql"),
	// 					Type: to.Ptr("Microsoft.Synapse/workspaces/privateEndpointConnections"),
	// 					ID: to.Ptr("/subscriptions/01234567-89ab-4def-0123-456789abcdef/resourceGroups/ExampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/privateEndpointConnections/ExamplePrivateEndpointConnection"),
	// 					Properties: &armsynapse.PrivateEndpointConnectionProperties{
	// 						PrivateEndpoint: &armsynapse.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/01234567-89ab-4def-0123-456789abcdef/resourceGroups/ExampleResourceGroup/providers/Microsoft.Network/privateEndpoints/ExamplePrivateEndpoint"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armsynapse.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("Auto-approved"),
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr("Approved"),
	// 						},
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 					},
	// 			}},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			PurviewConfiguration: &armsynapse.PurviewConfiguration{
	// 				PurviewResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.ProjectPurview/accounts/accountname1"),
	// 			},
	// 			SQLAdministratorLogin: to.Ptr("login"),
	// 			WorkspaceRepositoryConfiguration: &armsynapse.WorkspaceRepositoryConfiguration{
	// 				Type: to.Ptr("FactoryGitHubConfiguration"),
	// 				AccountName: to.Ptr("myGithubAccount"),
	// 				CollaborationBranch: to.Ptr("master"),
	// 				HostName: to.Ptr(""),
	// 				ProjectName: to.Ptr("myProject"),
	// 				RepositoryName: to.Ptr("myRepository"),
	// 				RootFolder: to.Ptr("/"),
	// 				TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 			},
	// 			WorkspaceUID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/UpdateWorkspace.json
func ExampleWorkspacesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspacesClient().BeginUpdate(ctx, "resourceGroup1", "workspace1", armsynapse.WorkspacePatchInfo{
		Identity: &armsynapse.ManagedIdentity{
			Type: to.Ptr(armsynapse.ResourceIdentityTypeSystemAssigned),
		},
		Properties: &armsynapse.WorkspacePatchProperties{
			Encryption: &armsynapse.EncryptionDetails{
				Cmk: &armsynapse.CustomerManagedKeyDetails{
					Key: &armsynapse.WorkspaceKeyDetails{
						Name:        to.Ptr("default"),
						KeyVaultURL: to.Ptr("https://vault.azure.net/keys/key1"),
					},
				},
			},
			ManagedVirtualNetworkSettings: &armsynapse.ManagedVirtualNetworkSettings{
				AllowedAADTenantIDsForLinking: []*string{
					to.Ptr("740239CE-A25B-485B-86A0-262F29F6EBDB")},
				LinkedAccessCheckOnTargetResource: to.Ptr(false),
				PreventDataExfiltration:           to.Ptr(false),
			},
			PublicNetworkAccess: to.Ptr(armsynapse.WorkspacePublicNetworkAccessEnabled),
			PurviewConfiguration: &armsynapse.PurviewConfiguration{
				PurviewResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.ProjectPurview/accounts/accountname1"),
			},
			SQLAdministratorLoginPassword: to.Ptr("password"),
			WorkspaceRepositoryConfiguration: &armsynapse.WorkspaceRepositoryConfiguration{
				Type:                to.Ptr("FactoryGitHubConfiguration"),
				AccountName:         to.Ptr("adifferentacount"),
				CollaborationBranch: to.Ptr("master"),
				HostName:            to.Ptr(""),
				ProjectName:         to.Ptr("myproject"),
				RepositoryName:      to.Ptr("myrepository"),
				RootFolder:          to.Ptr("/"),
			},
		},
		Tags: map[string]*string{
			"key": to.Ptr("value"),
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
	// res.Workspace = armsynapse.Workspace{
	// 	Name: to.Ptr("workspace1"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace1"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armsynapse.WorkspaceProperties{
	// 		ConnectivityEndpoints: map[string]*string{
	// 			"dev": to.Ptr("workspace1.dev.projectarcadia.net"),
	// 			"sql": to.Ptr("workspace1.sql.projectarcadia.net"),
	// 		},
	// 		DefaultDataLakeStorage: &armsynapse.DataLakeStorageAccountDetails{
	// 			AccountURL: to.Ptr("https://accountname.dfs.core.windows.net"),
	// 			Filesystem: to.Ptr("default"),
	// 		},
	// 		Encryption: &armsynapse.EncryptionDetails{
	// 			Cmk: &armsynapse.CustomerManagedKeyDetails{
	// 				Key: &armsynapse.WorkspaceKeyDetails{
	// 					Name: to.Ptr("default"),
	// 					KeyVaultURL: to.Ptr("https://vault.azure.net/keys/key1"),
	// 				},
	// 				Status: to.Ptr("Consistent"),
	// 			},
	// 			DoubleEncryptionEnabled: to.Ptr(true),
	// 		},
	// 		ManagedResourceGroupName: to.Ptr("resourceGroup2"),
	// 		ManagedVirtualNetworkSettings: &armsynapse.ManagedVirtualNetworkSettings{
	// 			AllowedAADTenantIDsForLinking: []*string{
	// 				to.Ptr("740239CE-A25B-485B-86A0-262F29F6EBDB")},
	// 				LinkedAccessCheckOnTargetResource: to.Ptr(false),
	// 				PreventDataExfiltration: to.Ptr(false),
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			PublicNetworkAccess: to.Ptr(armsynapse.WorkspacePublicNetworkAccessEnabled),
	// 			PurviewConfiguration: &armsynapse.PurviewConfiguration{
	// 				PurviewResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.ProjectPurview/accounts/accountname1"),
	// 			},
	// 			SQLAdministratorLogin: to.Ptr("login"),
	// 			WorkspaceUID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateWorkspace.json
func ExampleWorkspacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspacesClient().BeginCreateOrUpdate(ctx, "resourceGroup1", "workspace1", armsynapse.Workspace{
		Location: to.Ptr("East US"),
		Tags: map[string]*string{
			"key": to.Ptr("value"),
		},
		Identity: &armsynapse.ManagedIdentity{
			Type: to.Ptr(armsynapse.ResourceIdentityTypeSystemAssignedUserAssigned),
			UserAssignedIdentities: map[string]*armsynapse.UserAssignedManagedIdentity{
				"/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1": {},
			},
		},
		Properties: &armsynapse.WorkspaceProperties{
			CspWorkspaceAdminProperties: &armsynapse.CspWorkspaceAdminProperties{
				InitialWorkspaceAdminObjectID: to.Ptr("6c20646f-8050-49ec-b3b1-80a0e58e454d"),
			},
			DefaultDataLakeStorage: &armsynapse.DataLakeStorageAccountDetails{
				AccountURL: to.Ptr("https://accountname.dfs.core.windows.net"),
				Filesystem: to.Ptr("default"),
			},
			Encryption: &armsynapse.EncryptionDetails{
				Cmk: &armsynapse.CustomerManagedKeyDetails{
					KekIdentity: &armsynapse.KekIdentityProperties{
						UseSystemAssignedIdentity: false,
						UserAssignedIdentity:      to.Ptr("/subscriptions/b64d7b94-73e7-4d36-94b2-7764ea3fd74a/resourcegroups/SynapseCI/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
					},
					Key: &armsynapse.WorkspaceKeyDetails{
						Name:        to.Ptr("default"),
						KeyVaultURL: to.Ptr("https://vault.azure.net/keys/key1"),
					},
				},
			},
			ManagedResourceGroupName: to.Ptr("workspaceManagedResourceGroupUnique"),
			ManagedVirtualNetwork:    to.Ptr("default"),
			ManagedVirtualNetworkSettings: &armsynapse.ManagedVirtualNetworkSettings{
				AllowedAADTenantIDsForLinking: []*string{
					to.Ptr("740239CE-A25B-485B-86A0-262F29F6EBDB")},
				LinkedAccessCheckOnTargetResource: to.Ptr(false),
				PreventDataExfiltration:           to.Ptr(false),
			},
			PublicNetworkAccess: to.Ptr(armsynapse.WorkspacePublicNetworkAccessEnabled),
			PurviewConfiguration: &armsynapse.PurviewConfiguration{
				PurviewResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.ProjectPurview/accounts/accountname1"),
			},
			SQLAdministratorLogin:         to.Ptr("login"),
			SQLAdministratorLoginPassword: to.Ptr("password"),
			WorkspaceRepositoryConfiguration: &armsynapse.WorkspaceRepositoryConfiguration{
				Type:                to.Ptr("FactoryGitHubConfiguration"),
				AccountName:         to.Ptr("mygithubaccount"),
				CollaborationBranch: to.Ptr("master"),
				HostName:            to.Ptr(""),
				ProjectName:         to.Ptr("myproject"),
				RepositoryName:      to.Ptr("myrepository"),
				RootFolder:          to.Ptr("/"),
			},
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
	// res.Workspace = armsynapse.Workspace{
	// 	Name: to.Ptr("workspace1"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace1"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Identity: &armsynapse.ManagedIdentity{
	// 		Type: to.Ptr(armsynapse.ResourceIdentityTypeSystemAssignedUserAssigned),
	// 		PrincipalID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		TenantID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		UserAssignedIdentities: map[string]*armsynapse.UserAssignedManagedIdentity{
	// 			"/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1": &armsynapse.UserAssignedManagedIdentity{
	// 				ClientID: to.Ptr("ffffffff-8888-4444-8888-333333333333"),
	// 				PrincipalID: to.Ptr("eeeeeeee-9999-4444-8888-333333333333"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armsynapse.WorkspaceProperties{
	// 		ConnectivityEndpoints: map[string]*string{
	// 			"dev": to.Ptr("workspace1.dev.projectarcadia.net"),
	// 			"sql": to.Ptr("workspace1.sql.projectarcadia.net"),
	// 		},
	// 		DefaultDataLakeStorage: &armsynapse.DataLakeStorageAccountDetails{
	// 			AccountURL: to.Ptr("https://accountname.dfs.core.windows.net"),
	// 			Filesystem: to.Ptr("default"),
	// 		},
	// 		Encryption: &armsynapse.EncryptionDetails{
	// 			Cmk: &armsynapse.CustomerManagedKeyDetails{
	// 				KekIdentity: &armsynapse.KekIdentityProperties{
	// 					UseSystemAssignedIdentity: false,
	// 					UserAssignedIdentity: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/resourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uami1"),
	// 				},
	// 				Key: &armsynapse.WorkspaceKeyDetails{
	// 					Name: to.Ptr("default"),
	// 					KeyVaultURL: to.Ptr("https://vault.azure.net/keys/key1"),
	// 				},
	// 				Status: to.Ptr("Consistent"),
	// 			},
	// 			DoubleEncryptionEnabled: to.Ptr(true),
	// 		},
	// 		ManagedResourceGroupName: to.Ptr("workspaceManagedResourceGroupUnique"),
	// 		ManagedVirtualNetwork: to.Ptr("default"),
	// 		ManagedVirtualNetworkSettings: &armsynapse.ManagedVirtualNetworkSettings{
	// 			AllowedAADTenantIDsForLinking: []*string{
	// 				to.Ptr("740239CE-A25B-485B-86A0-262F29F6EBDB")},
	// 				LinkedAccessCheckOnTargetResource: to.Ptr(false),
	// 				PreventDataExfiltration: to.Ptr(false),
	// 			},
	// 			PrivateEndpointConnections: []*armsynapse.PrivateEndpointConnection{
	// 				{
	// 					Name: to.Ptr("sql"),
	// 					Type: to.Ptr("Microsoft.Synapse/workspaces/privateEndpointConnections"),
	// 					ID: to.Ptr("/subscriptions/01234567-89ab-4def-0123-456789abcdef/resourceGroups/ExampleResourceGroup/providers/Microsoft.Synapse/workspaces/ExampleWorkspace/privateEndpointConnections/ExamplePrivateEndpointConnection"),
	// 					Properties: &armsynapse.PrivateEndpointConnectionProperties{
	// 						PrivateEndpoint: &armsynapse.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/01234567-89ab-4def-0123-456789abcdef/resourceGroups/ExampleResourceGroup/providers/Microsoft.Network/privateEndpoints/ExamplePrivateEndpoint"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armsynapse.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("Auto-approved"),
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr("Approved"),
	// 						},
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 					},
	// 			}},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			PublicNetworkAccess: to.Ptr(armsynapse.WorkspacePublicNetworkAccessEnabled),
	// 			PurviewConfiguration: &armsynapse.PurviewConfiguration{
	// 				PurviewResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.ProjectPurview/accounts/accountname1"),
	// 			},
	// 			SQLAdministratorLogin: to.Ptr("login"),
	// 			WorkspaceUID: to.Ptr("00000000-1111-2222-3333-444444444444"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteWorkspace.json
func ExampleWorkspacesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspacesClient().BeginDelete(ctx, "resourceGroup1", "workspace1", nil)
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
	// res.Workspace = armsynapse.Workspace{
	// 	Name: to.Ptr("workspace1"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace1"),
	// 	Location: to.Ptr("East US"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armsynapse.WorkspaceProperties{
	// 		ConnectivityEndpoints: map[string]*string{
	// 			"dev": to.Ptr("workspace1.dev.projectarcadia.net"),
	// 			"sql": to.Ptr("workspace1.sql.projectarcadia.net"),
	// 		},
	// 		DefaultDataLakeStorage: &armsynapse.DataLakeStorageAccountDetails{
	// 			AccountURL: to.Ptr("https://accountname.dfs.core.windows.net"),
	// 			Filesystem: to.Ptr("default"),
	// 		},
	// 		ManagedResourceGroupName: to.Ptr("resourceGroup2"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SQLAdministratorLogin: to.Ptr("login"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListWorkspacesInSubscription.json
func ExampleWorkspacesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListPager(nil)
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
		// page.WorkspaceInfoListResult = armsynapse.WorkspaceInfoListResult{
		// 	Value: []*armsynapse.Workspace{
		// 		{
		// 			Name: to.Ptr("workspace1"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace1"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Properties: &armsynapse.WorkspaceProperties{
		// 				ConnectivityEndpoints: map[string]*string{
		// 					"dev": to.Ptr("workspace1.dev.projectarcadia.net"),
		// 					"sql": to.Ptr("workspace1.sql.projectarcadia.net"),
		// 				},
		// 				DefaultDataLakeStorage: &armsynapse.DataLakeStorageAccountDetails{
		// 					AccountURL: to.Ptr("https://accountname.dfs.core.windows.net"),
		// 					Filesystem: to.Ptr("default"),
		// 				},
		// 				ManagedResourceGroupName: to.Ptr("resourceGroup2"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SQLAdministratorLogin: to.Ptr("login"),
		// 				WorkspaceUID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("workspace2"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace2"),
		// 			Location: to.Ptr("East US"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Properties: &armsynapse.WorkspaceProperties{
		// 				ConnectivityEndpoints: map[string]*string{
		// 					"dev": to.Ptr("workspace2.dev.projectarcadia.net"),
		// 					"sql": to.Ptr("workspace2.sql.projectarcadia.net"),
		// 				},
		// 				DefaultDataLakeStorage: &armsynapse.DataLakeStorageAccountDetails{
		// 					AccountURL: to.Ptr("https://accountname.dfs.core.windows.net"),
		// 					Filesystem: to.Ptr("default"),
		// 				},
		// 				ManagedResourceGroupName: to.Ptr("resourceGroup2"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SQLAdministratorLogin: to.Ptr("login"),
		// 				WorkspaceUID: to.Ptr("00000000-1111-2222-3333-444444444444"),
		// 			},
		// 	}},
		// }
	}
}
