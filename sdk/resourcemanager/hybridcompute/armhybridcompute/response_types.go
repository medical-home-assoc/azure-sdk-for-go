//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridcompute

// MachineExtensionsClientCreateOrUpdateResponse contains the response from method MachineExtensionsClient.BeginCreateOrUpdate.
type MachineExtensionsClientCreateOrUpdateResponse struct {
	MachineExtension
}

// MachineExtensionsClientDeleteResponse contains the response from method MachineExtensionsClient.BeginDelete.
type MachineExtensionsClientDeleteResponse struct {
	// placeholder for future response values
}

// MachineExtensionsClientGetResponse contains the response from method MachineExtensionsClient.Get.
type MachineExtensionsClientGetResponse struct {
	MachineExtension
}

// MachineExtensionsClientListResponse contains the response from method MachineExtensionsClient.NewListPager.
type MachineExtensionsClientListResponse struct {
	MachineExtensionsListResult
}

// MachineExtensionsClientUpdateResponse contains the response from method MachineExtensionsClient.BeginUpdate.
type MachineExtensionsClientUpdateResponse struct {
	MachineExtension
}

// MachinesClientCreateOrUpdateResponse contains the response from method MachinesClient.CreateOrUpdate.
type MachinesClientCreateOrUpdateResponse struct {
	Machine
}

// MachinesClientDeleteResponse contains the response from method MachinesClient.Delete.
type MachinesClientDeleteResponse struct {
	// placeholder for future response values
}

// MachinesClientGetResponse contains the response from method MachinesClient.Get.
type MachinesClientGetResponse struct {
	Machine
}

// MachinesClientListByResourceGroupResponse contains the response from method MachinesClient.NewListByResourceGroupPager.
type MachinesClientListByResourceGroupResponse struct {
	MachineListResult
}

// MachinesClientListBySubscriptionResponse contains the response from method MachinesClient.NewListBySubscriptionPager.
type MachinesClientListBySubscriptionResponse struct {
	MachineListResult
}

// MachinesClientUpdateResponse contains the response from method MachinesClient.Update.
type MachinesClientUpdateResponse struct {
	Machine
}

// ManagementClientUpgradeExtensionsResponse contains the response from method ManagementClient.BeginUpgradeExtensions.
type ManagementClientUpgradeExtensionsResponse struct {
	// placeholder for future response values
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse contains the response from method PrivateEndpointConnectionsClient.NewListByPrivateLinkScopePager.
type PrivateEndpointConnectionsClientListByPrivateLinkScopeResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByPrivateLinkScopeResponse contains the response from method PrivateLinkResourcesClient.NewListByPrivateLinkScopePager.
type PrivateLinkResourcesClientListByPrivateLinkScopeResponse struct {
	PrivateLinkResourceListResult
}

// PrivateLinkScopesClientCreateOrUpdateResponse contains the response from method PrivateLinkScopesClient.CreateOrUpdate.
type PrivateLinkScopesClientCreateOrUpdateResponse struct {
	PrivateLinkScope
}

// PrivateLinkScopesClientDeleteResponse contains the response from method PrivateLinkScopesClient.BeginDelete.
type PrivateLinkScopesClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateLinkScopesClientGetResponse contains the response from method PrivateLinkScopesClient.Get.
type PrivateLinkScopesClientGetResponse struct {
	PrivateLinkScope
}

// PrivateLinkScopesClientGetValidationDetailsForMachineResponse contains the response from method PrivateLinkScopesClient.GetValidationDetailsForMachine.
type PrivateLinkScopesClientGetValidationDetailsForMachineResponse struct {
	PrivateLinkScopeValidationDetails
}

// PrivateLinkScopesClientGetValidationDetailsResponse contains the response from method PrivateLinkScopesClient.GetValidationDetails.
type PrivateLinkScopesClientGetValidationDetailsResponse struct {
	PrivateLinkScopeValidationDetails
}

// PrivateLinkScopesClientListByResourceGroupResponse contains the response from method PrivateLinkScopesClient.NewListByResourceGroupPager.
type PrivateLinkScopesClientListByResourceGroupResponse struct {
	PrivateLinkScopeListResult
}

// PrivateLinkScopesClientListResponse contains the response from method PrivateLinkScopesClient.NewListPager.
type PrivateLinkScopesClientListResponse struct {
	PrivateLinkScopeListResult
}

// PrivateLinkScopesClientUpdateTagsResponse contains the response from method PrivateLinkScopesClient.UpdateTags.
type PrivateLinkScopesClientUpdateTagsResponse struct {
	PrivateLinkScope
}
