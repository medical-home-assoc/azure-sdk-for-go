//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armiotcentral

// AppsClientCheckNameAvailabilityResponse contains the response from method AppsClient.CheckNameAvailability.
type AppsClientCheckNameAvailabilityResponse struct {
	AppAvailabilityInfo
}

// AppsClientCheckSubdomainAvailabilityResponse contains the response from method AppsClient.CheckSubdomainAvailability.
type AppsClientCheckSubdomainAvailabilityResponse struct {
	AppAvailabilityInfo
}

// AppsClientCreateOrUpdateResponse contains the response from method AppsClient.BeginCreateOrUpdate.
type AppsClientCreateOrUpdateResponse struct {
	App
}

// AppsClientDeleteResponse contains the response from method AppsClient.BeginDelete.
type AppsClientDeleteResponse struct {
	// placeholder for future response values
}

// AppsClientGetResponse contains the response from method AppsClient.Get.
type AppsClientGetResponse struct {
	App
}

// AppsClientListByResourceGroupResponse contains the response from method AppsClient.NewListByResourceGroupPager.
type AppsClientListByResourceGroupResponse struct {
	AppListResult
}

// AppsClientListBySubscriptionResponse contains the response from method AppsClient.NewListBySubscriptionPager.
type AppsClientListBySubscriptionResponse struct {
	AppListResult
}

// AppsClientListTemplatesResponse contains the response from method AppsClient.NewListTemplatesPager.
type AppsClientListTemplatesResponse struct {
	AppTemplatesResult
}

// AppsClientUpdateResponse contains the response from method AppsClient.BeginUpdate.
type AppsClientUpdateResponse struct {
	App
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}
