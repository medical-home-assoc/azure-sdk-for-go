//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListBackends.json
func ExampleBackendClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBackendClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.BackendClientListByServiceOptions{Filter: nil,
		Top:  nil,
		Skip: nil,
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
		// page.BackendCollection = armapimanagement.BackendCollection{
		// 	Count: to.Ptr[int64](2),
		// 	Value: []*armapimanagement.BackendContract{
		// 		{
		// 			Name: to.Ptr("proxybackend"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/backends"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/proxybackend"),
		// 			Properties: &armapimanagement.BackendContractProperties{
		// 				Description: to.Ptr("description5308"),
		// 				Credentials: &armapimanagement.BackendCredentialsContract{
		// 					Authorization: &armapimanagement.BackendAuthorizationHeaderCredentials{
		// 						Parameter: to.Ptr("opensesma"),
		// 						Scheme: to.Ptr("Basic"),
		// 					},
		// 					Header: map[string][]*string{
		// 						"x-my-1": []*string{
		// 							to.Ptr("val1"),
		// 							to.Ptr("val2")},
		// 						},
		// 						Query: map[string][]*string{
		// 							"sv": []*string{
		// 								to.Ptr("xx"),
		// 								to.Ptr("bb"),
		// 								to.Ptr("cc")},
		// 							},
		// 						},
		// 						Proxy: &armapimanagement.BackendProxyContract{
		// 							Password: to.Ptr("<password>"),
		// 							URL: to.Ptr("http://192.168.1.1:8080"),
		// 							Username: to.Ptr("Contoso\\admin"),
		// 						},
		// 						TLS: &armapimanagement.BackendTLSProperties{
		// 							ValidateCertificateChain: to.Ptr(false),
		// 							ValidateCertificateName: to.Ptr(false),
		// 						},
		// 						URL: to.Ptr("https://backendname2644/"),
		// 						Protocol: to.Ptr(armapimanagement.BackendProtocolHTTP),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("sfbackend"),
		// 					Type: to.Ptr("Microsoft.ApiManagement/service/backends"),
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/sfbackend"),
		// 					Properties: &armapimanagement.BackendContractProperties{
		// 						Description: to.Ptr("Service Fabric Test App 1"),
		// 						Properties: &armapimanagement.BackendProperties{
		// 							ServiceFabricCluster: &armapimanagement.BackendServiceFabricClusterProperties{
		// 								ClientCertificateID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/cert1"),
		// 								ManagementEndpoints: []*string{
		// 									to.Ptr("https://somecluster.com")},
		// 									MaxPartitionResolutionRetries: to.Ptr[int32](5),
		// 									ServerX509Names: []*armapimanagement.X509CertificateName{
		// 										{
		// 											Name: to.Ptr("ServerCommonName1"),
		// 											IssuerCertificateThumbprint: to.Ptr("IssuerCertificateThumbprint1"),
		// 									}},
		// 								},
		// 							},
		// 							URL: to.Ptr("fabric:/mytestapp/mytestservice"),
		// 							Protocol: to.Ptr(armapimanagement.BackendProtocolHTTP),
		// 						},
		// 				}},
		// 			}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadBackend.json
func ExampleBackendClient_GetEntityTag() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewBackendClient().GetEntityTag(ctx, "rg1", "apimService1", "sfbackend", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetBackend.json
func ExampleBackendClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackendClient().Get(ctx, "rg1", "apimService1", "sfbackend", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackendContract = armapimanagement.BackendContract{
	// 	Name: to.Ptr("sfbackend"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/backends"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/sfbackend"),
	// 	Properties: &armapimanagement.BackendContractProperties{
	// 		Description: to.Ptr("Service Fabric Test App 1"),
	// 		Properties: &armapimanagement.BackendProperties{
	// 			ServiceFabricCluster: &armapimanagement.BackendServiceFabricClusterProperties{
	// 				ClientCertificateID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/cert1"),
	// 				ManagementEndpoints: []*string{
	// 					to.Ptr("https://somecluster.com")},
	// 					MaxPartitionResolutionRetries: to.Ptr[int32](5),
	// 					ServerX509Names: []*armapimanagement.X509CertificateName{
	// 						{
	// 							Name: to.Ptr("ServerCommonName1"),
	// 							IssuerCertificateThumbprint: to.Ptr("IssuerCertificateThumbprint1"),
	// 					}},
	// 				},
	// 			},
	// 			URL: to.Ptr("fabric:/mytestapp/mytestservice"),
	// 			Protocol: to.Ptr(armapimanagement.BackendProtocolHTTP),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateBackendProxyBackend.json
func ExampleBackendClient_CreateOrUpdate_apiManagementCreateBackendProxyBackend() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackendClient().CreateOrUpdate(ctx, "rg1", "apimService1", "proxybackend", armapimanagement.BackendContract{
		Properties: &armapimanagement.BackendContractProperties{
			Description: to.Ptr("description5308"),
			Credentials: &armapimanagement.BackendCredentialsContract{
				Authorization: &armapimanagement.BackendAuthorizationHeaderCredentials{
					Parameter: to.Ptr("opensesma"),
					Scheme:    to.Ptr("Basic"),
				},
				Header: map[string][]*string{
					"x-my-1": {
						to.Ptr("val1"),
						to.Ptr("val2")},
				},
				Query: map[string][]*string{
					"sv": {
						to.Ptr("xx"),
						to.Ptr("bb"),
						to.Ptr("cc")},
				},
			},
			Proxy: &armapimanagement.BackendProxyContract{
				Password: to.Ptr("<password>"),
				URL:      to.Ptr("http://192.168.1.1:8080"),
				Username: to.Ptr("Contoso\\admin"),
			},
			TLS: &armapimanagement.BackendTLSProperties{
				ValidateCertificateChain: to.Ptr(true),
				ValidateCertificateName:  to.Ptr(true),
			},
			URL:      to.Ptr("https://backendname2644/"),
			Protocol: to.Ptr(armapimanagement.BackendProtocolHTTP),
		},
	}, &armapimanagement.BackendClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackendContract = armapimanagement.BackendContract{
	// 	Name: to.Ptr("proxybackend"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/backends"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/proxybackend"),
	// 	Properties: &armapimanagement.BackendContractProperties{
	// 		Description: to.Ptr("description5308"),
	// 		Credentials: &armapimanagement.BackendCredentialsContract{
	// 			Authorization: &armapimanagement.BackendAuthorizationHeaderCredentials{
	// 				Parameter: to.Ptr("opensesma"),
	// 				Scheme: to.Ptr("Basic"),
	// 			},
	// 			Header: map[string][]*string{
	// 				"x-my-1": []*string{
	// 					to.Ptr("val1"),
	// 					to.Ptr("val2")},
	// 				},
	// 				Query: map[string][]*string{
	// 					"sv": []*string{
	// 						to.Ptr("xx"),
	// 						to.Ptr("bb"),
	// 						to.Ptr("cc")},
	// 					},
	// 				},
	// 				Proxy: &armapimanagement.BackendProxyContract{
	// 					Password: to.Ptr("<password>"),
	// 					URL: to.Ptr("http://192.168.1.1:8080"),
	// 					Username: to.Ptr("Contoso\\admin"),
	// 				},
	// 				TLS: &armapimanagement.BackendTLSProperties{
	// 					ValidateCertificateChain: to.Ptr(false),
	// 					ValidateCertificateName: to.Ptr(false),
	// 				},
	// 				URL: to.Ptr("https://backendname2644/"),
	// 				Protocol: to.Ptr(armapimanagement.BackendProtocolHTTP),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateBackendServiceFabric.json
func ExampleBackendClient_CreateOrUpdate_apiManagementCreateBackendServiceFabric() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackendClient().CreateOrUpdate(ctx, "rg1", "apimService1", "sfbackend", armapimanagement.BackendContract{
		Properties: &armapimanagement.BackendContractProperties{
			Description: to.Ptr("Service Fabric Test App 1"),
			Properties: &armapimanagement.BackendProperties{
				ServiceFabricCluster: &armapimanagement.BackendServiceFabricClusterProperties{
					ClientCertificateID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/cert1"),
					ManagementEndpoints: []*string{
						to.Ptr("https://somecluster.com")},
					MaxPartitionResolutionRetries: to.Ptr[int32](5),
					ServerX509Names: []*armapimanagement.X509CertificateName{
						{
							Name:                        to.Ptr("ServerCommonName1"),
							IssuerCertificateThumbprint: to.Ptr("IssuerCertificateThumbprint1"),
						}},
				},
			},
			URL:      to.Ptr("fabric:/mytestapp/mytestservice"),
			Protocol: to.Ptr(armapimanagement.BackendProtocolHTTP),
		},
	}, &armapimanagement.BackendClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackendContract = armapimanagement.BackendContract{
	// 	Name: to.Ptr("sfbackend"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/backends"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/sfbackend"),
	// 	Properties: &armapimanagement.BackendContractProperties{
	// 		Description: to.Ptr("Service Fabric Test App 1"),
	// 		Properties: &armapimanagement.BackendProperties{
	// 			ServiceFabricCluster: &armapimanagement.BackendServiceFabricClusterProperties{
	// 				ClientCertificateID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/cert1"),
	// 				ManagementEndpoints: []*string{
	// 					to.Ptr("https://somecluster.com")},
	// 					MaxPartitionResolutionRetries: to.Ptr[int32](5),
	// 					ServerX509Names: []*armapimanagement.X509CertificateName{
	// 						{
	// 							Name: to.Ptr("ServerCommonName1"),
	// 							IssuerCertificateThumbprint: to.Ptr("IssuerCertificateThumbprint1"),
	// 					}},
	// 				},
	// 			},
	// 			URL: to.Ptr("fabric:/mytestapp/mytestservice"),
	// 			Protocol: to.Ptr(armapimanagement.BackendProtocolHTTP),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateBackend.json
func ExampleBackendClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBackendClient().Update(ctx, "rg1", "apimService1", "proxybackend", "*", armapimanagement.BackendUpdateParameters{
		Properties: &armapimanagement.BackendUpdateParameterProperties{
			Description: to.Ptr("description5308"),
			TLS: &armapimanagement.BackendTLSProperties{
				ValidateCertificateChain: to.Ptr(false),
				ValidateCertificateName:  to.Ptr(true),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackendContract = armapimanagement.BackendContract{
	// 	Name: to.Ptr("proxybackend"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/backends"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/proxybackend"),
	// 	Properties: &armapimanagement.BackendContractProperties{
	// 		Description: to.Ptr("description5308"),
	// 		Credentials: &armapimanagement.BackendCredentialsContract{
	// 			Authorization: &armapimanagement.BackendAuthorizationHeaderCredentials{
	// 				Parameter: to.Ptr("opensesma"),
	// 				Scheme: to.Ptr("Basic"),
	// 			},
	// 			Header: map[string][]*string{
	// 				"x-my-1": []*string{
	// 					to.Ptr("val1"),
	// 					to.Ptr("val2")},
	// 				},
	// 				Query: map[string][]*string{
	// 					"sv": []*string{
	// 						to.Ptr("xx"),
	// 						to.Ptr("bb"),
	// 						to.Ptr("cc")},
	// 					},
	// 				},
	// 				Proxy: &armapimanagement.BackendProxyContract{
	// 					Password: to.Ptr("<password>"),
	// 					URL: to.Ptr("http://192.168.1.1:8080"),
	// 					Username: to.Ptr("Contoso\\admin"),
	// 				},
	// 				TLS: &armapimanagement.BackendTLSProperties{
	// 					ValidateCertificateChain: to.Ptr(false),
	// 					ValidateCertificateName: to.Ptr(true),
	// 				},
	// 				URL: to.Ptr("https://backendname2644/"),
	// 				Protocol: to.Ptr(armapimanagement.BackendProtocolHTTP),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteBackend.json
func ExampleBackendClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewBackendClient().Delete(ctx, "rg1", "apimService1", "sfbackend", "*", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementBackendReconnect.json
func ExampleBackendClient_Reconnect() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewBackendClient().Reconnect(ctx, "rg1", "apimService1", "proxybackend", &armapimanagement.BackendClientReconnectOptions{Parameters: &armapimanagement.BackendReconnectContract{
		Properties: &armapimanagement.BackendReconnectProperties{
			After: to.Ptr("PT3S"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
