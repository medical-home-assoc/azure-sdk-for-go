//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armm365securityandcompliance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/SecurityCenterServiceGet.json
func ExamplePrivateLinkServicesForM365SecurityCenterClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armm365securityandcompliance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkServicesForM365SecurityCenterClient().Get(ctx, "rg1", "service1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkServicesForM365SecurityCenterDescription = armm365securityandcompliance.PrivateLinkServicesForM365SecurityCenterDescription{
	// 	Name: to.Ptr("service1"),
	// 	Type: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter"),
	// 	Etag: to.Ptr("etagvalue"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/service1"),
	// 	Kind: to.Ptr(armm365securityandcompliance.KindFhirR4),
	// 	Location: to.Ptr("West US"),
	// 	SystemData: &armm365securityandcompliance.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-24T13:30:28.958Z"); return t}()),
	// 		CreatedBy: to.Ptr("sove"),
	// 		CreatedByType: to.Ptr(armm365securityandcompliance.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-24T13:30:28.958Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sove"),
	// 		LastModifiedByType: to.Ptr(armm365securityandcompliance.CreatedByTypeUser),
	// 	},
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armm365securityandcompliance.ServicesProperties{
	// 		AccessPolicies: []*armm365securityandcompliance.ServiceAccessPolicyEntry{
	// 			{
	// 				ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
	// 			},
	// 			{
	// 				ObjectID: to.Ptr("5b307da8-43d4-492b-8b66-b0294ade872f"),
	// 		}},
	// 		AuthenticationConfiguration: &armm365securityandcompliance.ServiceAuthenticationConfigurationInfo{
	// 			Audience: to.Ptr("https://azurehealthcareapis.com"),
	// 			Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
	// 			SmartProxyEnabled: to.Ptr(true),
	// 		},
	// 		CorsConfiguration: &armm365securityandcompliance.ServiceCorsConfigurationInfo{
	// 			AllowCredentials: to.Ptr(false),
	// 			Headers: []*string{
	// 				to.Ptr("*")},
	// 				MaxAge: to.Ptr[int64](1440),
	// 				Methods: []*string{
	// 					to.Ptr("DELETE"),
	// 					to.Ptr("GET"),
	// 					to.Ptr("OPTIONS"),
	// 					to.Ptr("PATCH"),
	// 					to.Ptr("POST"),
	// 					to.Ptr("PUT")},
	// 					Origins: []*string{
	// 						to.Ptr("*")},
	// 					},
	// 					CosmosDbConfiguration: &armm365securityandcompliance.ServiceCosmosDbConfigurationInfo{
	// 						KeyVaultKeyURI: to.Ptr("https://my-vault.vault.azure.net/keys/my-key"),
	// 						OfferThroughput: to.Ptr[int64](1000),
	// 					},
	// 					PrivateEndpointConnections: []*armm365securityandcompliance.PrivateEndpointConnection{
	// 					},
	// 					ProvisioningState: to.Ptr(armm365securityandcompliance.ProvisioningStateSucceeded),
	// 					PublicNetworkAccess: to.Ptr(armm365securityandcompliance.PublicNetworkAccessDisabled),
	// 				},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/SecurityCenterServiceCreate.json
func ExamplePrivateLinkServicesForM365SecurityCenterClient_BeginCreateOrUpdate_createOrUpdateAServiceWithAllParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armm365securityandcompliance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateLinkServicesForM365SecurityCenterClient().BeginCreateOrUpdate(ctx, "rg1", "service1", armm365securityandcompliance.PrivateLinkServicesForM365SecurityCenterDescription{
		Identity: &armm365securityandcompliance.ServicesResourceIdentity{
			Type: to.Ptr(armm365securityandcompliance.ManagedServiceIdentityTypeSystemAssigned),
		},
		Kind:     to.Ptr(armm365securityandcompliance.KindFhirR4),
		Location: to.Ptr("westus2"),
		Tags:     map[string]*string{},
		Properties: &armm365securityandcompliance.ServicesProperties{
			AccessPolicies: []*armm365securityandcompliance.ServiceAccessPolicyEntry{
				{
					ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
				},
				{
					ObjectID: to.Ptr("5b307da8-43d4-492b-8b66-b0294ade872f"),
				}},
			AuthenticationConfiguration: &armm365securityandcompliance.ServiceAuthenticationConfigurationInfo{
				Audience:          to.Ptr("https://azurehealthcareapis.com"),
				Authority:         to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
				SmartProxyEnabled: to.Ptr(true),
			},
			CorsConfiguration: &armm365securityandcompliance.ServiceCorsConfigurationInfo{
				AllowCredentials: to.Ptr(false),
				Headers: []*string{
					to.Ptr("*")},
				MaxAge: to.Ptr[int64](1440),
				Methods: []*string{
					to.Ptr("DELETE"),
					to.Ptr("GET"),
					to.Ptr("OPTIONS"),
					to.Ptr("PATCH"),
					to.Ptr("POST"),
					to.Ptr("PUT")},
				Origins: []*string{
					to.Ptr("*")},
			},
			CosmosDbConfiguration: &armm365securityandcompliance.ServiceCosmosDbConfigurationInfo{
				KeyVaultKeyURI:  to.Ptr("https://my-vault.vault.azure.net/keys/my-key"),
				OfferThroughput: to.Ptr[int64](1000),
			},
			ExportConfiguration: &armm365securityandcompliance.ServiceExportConfigurationInfo{
				StorageAccountName: to.Ptr("existingStorageAccount"),
			},
			PrivateEndpointConnections: []*armm365securityandcompliance.PrivateEndpointConnection{},
			PublicNetworkAccess:        to.Ptr(armm365securityandcompliance.PublicNetworkAccessDisabled),
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
	// res.PrivateLinkServicesForM365SecurityCenterDescription = armm365securityandcompliance.PrivateLinkServicesForM365SecurityCenterDescription{
	// 	Name: to.Ptr("service1"),
	// 	Type: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter"),
	// 	Etag: to.Ptr("etagvalue"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/service1"),
	// 	Identity: &armm365securityandcompliance.ServicesResourceIdentity{
	// 		Type: to.Ptr(armm365securityandcompliance.ManagedServiceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("03fe6ae0-952c-4e4b-954b-cc0364dd252e"),
	// 		TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d8cd011db47"),
	// 	},
	// 	Kind: to.Ptr(armm365securityandcompliance.KindFhirR4),
	// 	Location: to.Ptr("West US 2"),
	// 	SystemData: &armm365securityandcompliance.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-24T13:30:28.958Z"); return t}()),
	// 		CreatedBy: to.Ptr("sove"),
	// 		CreatedByType: to.Ptr(armm365securityandcompliance.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-24T13:30:28.958Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sove"),
	// 		LastModifiedByType: to.Ptr(armm365securityandcompliance.CreatedByTypeUser),
	// 	},
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armm365securityandcompliance.ServicesProperties{
	// 		AccessPolicies: []*armm365securityandcompliance.ServiceAccessPolicyEntry{
	// 			{
	// 				ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
	// 			},
	// 			{
	// 				ObjectID: to.Ptr("5b307da8-43d4-492b-8b66-b0294ade872f"),
	// 		}},
	// 		AuthenticationConfiguration: &armm365securityandcompliance.ServiceAuthenticationConfigurationInfo{
	// 			Audience: to.Ptr("https://azurehealthcareapis.com"),
	// 			Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
	// 			SmartProxyEnabled: to.Ptr(true),
	// 		},
	// 		CorsConfiguration: &armm365securityandcompliance.ServiceCorsConfigurationInfo{
	// 			AllowCredentials: to.Ptr(false),
	// 			Headers: []*string{
	// 				to.Ptr("*")},
	// 				MaxAge: to.Ptr[int64](1440),
	// 				Methods: []*string{
	// 					to.Ptr("DELETE"),
	// 					to.Ptr("GET"),
	// 					to.Ptr("OPTIONS"),
	// 					to.Ptr("PATCH"),
	// 					to.Ptr("POST"),
	// 					to.Ptr("PUT")},
	// 					Origins: []*string{
	// 						to.Ptr("*")},
	// 					},
	// 					CosmosDbConfiguration: &armm365securityandcompliance.ServiceCosmosDbConfigurationInfo{
	// 						KeyVaultKeyURI: to.Ptr("https://my-vault.vault.azure.net/keys/my-key"),
	// 						OfferThroughput: to.Ptr[int64](1000),
	// 					},
	// 					ExportConfiguration: &armm365securityandcompliance.ServiceExportConfigurationInfo{
	// 						StorageAccountName: to.Ptr("existingStorageAccount"),
	// 					},
	// 					PrivateEndpointConnections: []*armm365securityandcompliance.PrivateEndpointConnection{
	// 					},
	// 					ProvisioningState: to.Ptr(armm365securityandcompliance.ProvisioningStateSucceeded),
	// 					PublicNetworkAccess: to.Ptr(armm365securityandcompliance.PublicNetworkAccessDisabled),
	// 				},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/SecurityCenterServiceCreateMinimum.json
func ExamplePrivateLinkServicesForM365SecurityCenterClient_BeginCreateOrUpdate_createOrUpdateAServiceWithMinimumParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armm365securityandcompliance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateLinkServicesForM365SecurityCenterClient().BeginCreateOrUpdate(ctx, "rg1", "service2", armm365securityandcompliance.PrivateLinkServicesForM365SecurityCenterDescription{
		Kind:     to.Ptr(armm365securityandcompliance.KindFhirR4),
		Location: to.Ptr("westus2"),
		Tags:     map[string]*string{},
		Properties: &armm365securityandcompliance.ServicesProperties{
			AccessPolicies: []*armm365securityandcompliance.ServiceAccessPolicyEntry{
				{
					ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
				}},
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
	// res.PrivateLinkServicesForM365SecurityCenterDescription = armm365securityandcompliance.PrivateLinkServicesForM365SecurityCenterDescription{
	// 	Name: to.Ptr("service2"),
	// 	Type: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter"),
	// 	Etag: to.Ptr("etagvalue"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/service2"),
	// 	Kind: to.Ptr(armm365securityandcompliance.KindFhirR4),
	// 	Location: to.Ptr("westus2"),
	// 	SystemData: &armm365securityandcompliance.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-24T13:30:28.958Z"); return t}()),
	// 		CreatedBy: to.Ptr("sove"),
	// 		CreatedByType: to.Ptr(armm365securityandcompliance.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-24T13:30:28.958Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sove"),
	// 		LastModifiedByType: to.Ptr(armm365securityandcompliance.CreatedByTypeUser),
	// 	},
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armm365securityandcompliance.ServicesProperties{
	// 		AccessPolicies: []*armm365securityandcompliance.ServiceAccessPolicyEntry{
	// 			{
	// 				ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
	// 		}},
	// 		AuthenticationConfiguration: &armm365securityandcompliance.ServiceAuthenticationConfigurationInfo{
	// 			Audience: to.Ptr("https://azurehealthcareapis.com"),
	// 			Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
	// 			SmartProxyEnabled: to.Ptr(false),
	// 		},
	// 		CorsConfiguration: &armm365securityandcompliance.ServiceCorsConfigurationInfo{
	// 			AllowCredentials: to.Ptr(false),
	// 			Headers: []*string{
	// 			},
	// 			Methods: []*string{
	// 			},
	// 			Origins: []*string{
	// 			},
	// 		},
	// 		CosmosDbConfiguration: &armm365securityandcompliance.ServiceCosmosDbConfigurationInfo{
	// 			OfferThroughput: to.Ptr[int64](1000),
	// 		},
	// 		PrivateEndpointConnections: []*armm365securityandcompliance.PrivateEndpointConnection{
	// 		},
	// 		ProvisioningState: to.Ptr(armm365securityandcompliance.ProvisioningStateSucceeded),
	// 		PublicNetworkAccess: to.Ptr(armm365securityandcompliance.PublicNetworkAccessDisabled),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/SecurityCenterServicePatch.json
func ExamplePrivateLinkServicesForM365SecurityCenterClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armm365securityandcompliance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateLinkServicesForM365SecurityCenterClient().BeginUpdate(ctx, "rg1", "service1", armm365securityandcompliance.ServicesPatchDescription{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
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
	// res.PrivateLinkServicesForM365SecurityCenterDescription = armm365securityandcompliance.PrivateLinkServicesForM365SecurityCenterDescription{
	// 	Name: to.Ptr("service1"),
	// 	Type: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter"),
	// 	Etag: to.Ptr("etagvalue"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/service1"),
	// 	Kind: to.Ptr(armm365securityandcompliance.KindFhirR4),
	// 	Location: to.Ptr("West US"),
	// 	SystemData: &armm365securityandcompliance.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-24T13:30:28.958Z"); return t}()),
	// 		CreatedBy: to.Ptr("sove"),
	// 		CreatedByType: to.Ptr(armm365securityandcompliance.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-24T13:30:28.958Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sove"),
	// 		LastModifiedByType: to.Ptr(armm365securityandcompliance.CreatedByTypeUser),
	// 	},
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armm365securityandcompliance.ServicesProperties{
	// 		AccessPolicies: []*armm365securityandcompliance.ServiceAccessPolicyEntry{
	// 			{
	// 				ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
	// 			},
	// 			{
	// 				ObjectID: to.Ptr("5b307da8-43d4-492b-8b66-b0294ade872f"),
	// 		}},
	// 		AuthenticationConfiguration: &armm365securityandcompliance.ServiceAuthenticationConfigurationInfo{
	// 			Audience: to.Ptr("https://azurehealthcareapis.com"),
	// 			Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
	// 			SmartProxyEnabled: to.Ptr(true),
	// 		},
	// 		CorsConfiguration: &armm365securityandcompliance.ServiceCorsConfigurationInfo{
	// 			AllowCredentials: to.Ptr(false),
	// 			Headers: []*string{
	// 				to.Ptr("*")},
	// 				MaxAge: to.Ptr[int64](1440),
	// 				Methods: []*string{
	// 					to.Ptr("DELETE"),
	// 					to.Ptr("GET"),
	// 					to.Ptr("OPTIONS"),
	// 					to.Ptr("PATCH"),
	// 					to.Ptr("POST"),
	// 					to.Ptr("PUT")},
	// 					Origins: []*string{
	// 						to.Ptr("*")},
	// 					},
	// 					CosmosDbConfiguration: &armm365securityandcompliance.ServiceCosmosDbConfigurationInfo{
	// 						KeyVaultKeyURI: to.Ptr("https://my-vault.vault.azure.net/keys/my-key"),
	// 						OfferThroughput: to.Ptr[int64](1000),
	// 					},
	// 					PrivateEndpointConnections: []*armm365securityandcompliance.PrivateEndpointConnection{
	// 					},
	// 					ProvisioningState: to.Ptr(armm365securityandcompliance.ProvisioningStateSucceeded),
	// 					PublicNetworkAccess: to.Ptr(armm365securityandcompliance.PublicNetworkAccessDisabled),
	// 				},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/SecurityCenterServiceDelete.json
func ExamplePrivateLinkServicesForM365SecurityCenterClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armm365securityandcompliance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateLinkServicesForM365SecurityCenterClient().BeginDelete(ctx, "rg1", "service1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/SecurityCenterServiceList.json
func ExamplePrivateLinkServicesForM365SecurityCenterClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armm365securityandcompliance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkServicesForM365SecurityCenterClient().NewListPager(nil)
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
		// page.PrivateLinkServicesForM365SecurityCenterDescriptionListResult = armm365securityandcompliance.PrivateLinkServicesForM365SecurityCenterDescriptionListResult{
		// 	Value: []*armm365securityandcompliance.PrivateLinkServicesForM365SecurityCenterDescription{
		// 		{
		// 			Name: to.Ptr("service1"),
		// 			Type: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter"),
		// 			Etag: to.Ptr("etag"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/service1"),
		// 			Kind: to.Ptr(armm365securityandcompliance.KindFhirR4),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armm365securityandcompliance.ServicesProperties{
		// 				AccessPolicies: []*armm365securityandcompliance.ServiceAccessPolicyEntry{
		// 					{
		// 						ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
		// 					},
		// 					{
		// 						ObjectID: to.Ptr("5b307da8-43d4-492b-8b66-b0294ade872f"),
		// 				}},
		// 				AuthenticationConfiguration: &armm365securityandcompliance.ServiceAuthenticationConfigurationInfo{
		// 					Audience: to.Ptr("https://azurehealthcareapis.com"),
		// 					Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
		// 					SmartProxyEnabled: to.Ptr(true),
		// 				},
		// 				CorsConfiguration: &armm365securityandcompliance.ServiceCorsConfigurationInfo{
		// 					AllowCredentials: to.Ptr(false),
		// 					Headers: []*string{
		// 						to.Ptr("*")},
		// 						MaxAge: to.Ptr[int64](1440),
		// 						Methods: []*string{
		// 							to.Ptr("DELETE"),
		// 							to.Ptr("GET"),
		// 							to.Ptr("OPTIONS"),
		// 							to.Ptr("PATCH"),
		// 							to.Ptr("POST"),
		// 							to.Ptr("PUT")},
		// 							Origins: []*string{
		// 								to.Ptr("*")},
		// 							},
		// 							CosmosDbConfiguration: &armm365securityandcompliance.ServiceCosmosDbConfigurationInfo{
		// 								KeyVaultKeyURI: to.Ptr("https://my-vault.vault.azure.net/keys/my-key"),
		// 								OfferThroughput: to.Ptr[int64](1000),
		// 							},
		// 							PrivateEndpointConnections: []*armm365securityandcompliance.PrivateEndpointConnection{
		// 							},
		// 							ProvisioningState: to.Ptr(armm365securityandcompliance.ProvisioningStateSucceeded),
		// 							PublicNetworkAccess: to.Ptr(armm365securityandcompliance.PublicNetworkAccessDisabled),
		// 						},
		// 				}},
		// 			}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/m365securityandcompliance/resource-manager/Microsoft.M365SecurityAndCompliance/preview/2021-03-25-preview/examples/SecurityCenterServiceListByResourceGroup.json
func ExamplePrivateLinkServicesForM365SecurityCenterClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armm365securityandcompliance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkServicesForM365SecurityCenterClient().NewListByResourceGroupPager("rgname", nil)
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
		// page.PrivateLinkServicesForM365SecurityCenterDescriptionListResult = armm365securityandcompliance.PrivateLinkServicesForM365SecurityCenterDescriptionListResult{
		// 	Value: []*armm365securityandcompliance.PrivateLinkServicesForM365SecurityCenterDescription{
		// 		{
		// 			Name: to.Ptr("service1"),
		// 			Type: to.Ptr("Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter"),
		// 			Etag: to.Ptr("etagvalue"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/dddb8dcb-effb-4290-bb47-ce1e8440c729"),
		// 			Kind: to.Ptr(armm365securityandcompliance.KindFhirR4),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armm365securityandcompliance.ServicesProperties{
		// 				AccessPolicies: []*armm365securityandcompliance.ServiceAccessPolicyEntry{
		// 					{
		// 						ObjectID: to.Ptr("c487e7d1-3210-41a3-8ccc-e9372b78da47"),
		// 					},
		// 					{
		// 						ObjectID: to.Ptr("5b307da8-43d4-492b-8b66-b0294ade872f"),
		// 				}},
		// 				AuthenticationConfiguration: &armm365securityandcompliance.ServiceAuthenticationConfigurationInfo{
		// 					Audience: to.Ptr("https://azurehealthcareapis.com"),
		// 					Authority: to.Ptr("https://login.microsoftonline.com/abfde7b2-df0f-47e6-aabf-2462b07508dc"),
		// 					SmartProxyEnabled: to.Ptr(true),
		// 				},
		// 				CorsConfiguration: &armm365securityandcompliance.ServiceCorsConfigurationInfo{
		// 					AllowCredentials: to.Ptr(false),
		// 					Headers: []*string{
		// 						to.Ptr("*")},
		// 						MaxAge: to.Ptr[int64](1440),
		// 						Methods: []*string{
		// 							to.Ptr("DELETE"),
		// 							to.Ptr("GET"),
		// 							to.Ptr("OPTIONS"),
		// 							to.Ptr("PATCH"),
		// 							to.Ptr("POST"),
		// 							to.Ptr("PUT")},
		// 							Origins: []*string{
		// 								to.Ptr("*")},
		// 							},
		// 							CosmosDbConfiguration: &armm365securityandcompliance.ServiceCosmosDbConfigurationInfo{
		// 								KeyVaultKeyURI: to.Ptr("https://my-vault.vault.azure.net/keys/my-key"),
		// 								OfferThroughput: to.Ptr[int64](1000),
		// 							},
		// 							PrivateEndpointConnections: []*armm365securityandcompliance.PrivateEndpointConnection{
		// 							},
		// 							ProvisioningState: to.Ptr(armm365securityandcompliance.ProvisioningStateSucceeded),
		// 							PublicNetworkAccess: to.Ptr(armm365securityandcompliance.PublicNetworkAccessDisabled),
		// 						},
		// 				}},
		// 			}
	}
}
