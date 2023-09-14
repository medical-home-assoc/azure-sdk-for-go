//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ManagedInstancePrivateEndpointConnectionsClient contains the methods for the ManagedInstancePrivateEndpointConnections group.
// Don't use this type directly, use NewManagedInstancePrivateEndpointConnectionsClient() instead.
type ManagedInstancePrivateEndpointConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedInstancePrivateEndpointConnectionsClient creates a new instance of ManagedInstancePrivateEndpointConnectionsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedInstancePrivateEndpointConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedInstancePrivateEndpointConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName+".ManagedInstancePrivateEndpointConnectionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedInstancePrivateEndpointConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Approve or reject a private endpoint connection with a given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - ManagedInstancePrivateEndpointConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for
//     the ManagedInstancePrivateEndpointConnectionsClient.BeginCreateOrUpdate method.
func (client *ManagedInstancePrivateEndpointConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, privateEndpointConnectionName string, parameters ManagedInstancePrivateEndpointConnection, options *ManagedInstancePrivateEndpointConnectionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ManagedInstancePrivateEndpointConnectionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, privateEndpointConnectionName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ManagedInstancePrivateEndpointConnectionsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ManagedInstancePrivateEndpointConnectionsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Approve or reject a private endpoint connection with a given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
func (client *ManagedInstancePrivateEndpointConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, privateEndpointConnectionName string, parameters ManagedInstancePrivateEndpointConnection, options *ManagedInstancePrivateEndpointConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, privateEndpointConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedInstancePrivateEndpointConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, privateEndpointConnectionName string, parameters ManagedInstancePrivateEndpointConnection, options *ManagedInstancePrivateEndpointConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a private endpoint connection with a given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - ManagedInstancePrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the ManagedInstancePrivateEndpointConnectionsClient.BeginDelete
//     method.
func (client *ManagedInstancePrivateEndpointConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, managedInstanceName string, privateEndpointConnectionName string, options *ManagedInstancePrivateEndpointConnectionsClientBeginDeleteOptions) (*runtime.Poller[ManagedInstancePrivateEndpointConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, managedInstanceName, privateEndpointConnectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ManagedInstancePrivateEndpointConnectionsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ManagedInstancePrivateEndpointConnectionsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a private endpoint connection with a given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
func (client *ManagedInstancePrivateEndpointConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, managedInstanceName string, privateEndpointConnectionName string, options *ManagedInstancePrivateEndpointConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managedInstanceName, privateEndpointConnectionName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ManagedInstancePrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, privateEndpointConnectionName string, options *ManagedInstancePrivateEndpointConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a private endpoint connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - privateEndpointConnectionName - The name of the private endpoint connection.
//   - options - ManagedInstancePrivateEndpointConnectionsClientGetOptions contains the optional parameters for the ManagedInstancePrivateEndpointConnectionsClient.Get
//     method.
func (client *ManagedInstancePrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, privateEndpointConnectionName string, options *ManagedInstancePrivateEndpointConnectionsClientGetOptions) (ManagedInstancePrivateEndpointConnectionsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, privateEndpointConnectionName, options)
	if err != nil {
		return ManagedInstancePrivateEndpointConnectionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedInstancePrivateEndpointConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedInstancePrivateEndpointConnectionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedInstancePrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, privateEndpointConnectionName string, options *ManagedInstancePrivateEndpointConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedInstancePrivateEndpointConnectionsClient) getHandleResponse(resp *http.Response) (ManagedInstancePrivateEndpointConnectionsClientGetResponse, error) {
	result := ManagedInstancePrivateEndpointConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstancePrivateEndpointConnection); err != nil {
		return ManagedInstancePrivateEndpointConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByManagedInstancePager - Gets all private endpoint connections on a server.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - options - ManagedInstancePrivateEndpointConnectionsClientListByManagedInstanceOptions contains the optional parameters
//     for the ManagedInstancePrivateEndpointConnectionsClient.NewListByManagedInstancePager method.
func (client *ManagedInstancePrivateEndpointConnectionsClient) NewListByManagedInstancePager(resourceGroupName string, managedInstanceName string, options *ManagedInstancePrivateEndpointConnectionsClientListByManagedInstanceOptions) *runtime.Pager[ManagedInstancePrivateEndpointConnectionsClientListByManagedInstanceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedInstancePrivateEndpointConnectionsClientListByManagedInstanceResponse]{
		More: func(page ManagedInstancePrivateEndpointConnectionsClientListByManagedInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedInstancePrivateEndpointConnectionsClientListByManagedInstanceResponse) (ManagedInstancePrivateEndpointConnectionsClientListByManagedInstanceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByManagedInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedInstancePrivateEndpointConnectionsClientListByManagedInstanceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ManagedInstancePrivateEndpointConnectionsClientListByManagedInstanceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedInstancePrivateEndpointConnectionsClientListByManagedInstanceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByManagedInstanceHandleResponse(resp)
		},
	})
}

// listByManagedInstanceCreateRequest creates the ListByManagedInstance request.
func (client *ManagedInstancePrivateEndpointConnectionsClient) listByManagedInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstancePrivateEndpointConnectionsClientListByManagedInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/privateEndpointConnections"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByManagedInstanceHandleResponse handles the ListByManagedInstance response.
func (client *ManagedInstancePrivateEndpointConnectionsClient) listByManagedInstanceHandleResponse(resp *http.Response) (ManagedInstancePrivateEndpointConnectionsClientListByManagedInstanceResponse, error) {
	result := ManagedInstancePrivateEndpointConnectionsClientListByManagedInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstancePrivateEndpointConnectionListResult); err != nil {
		return ManagedInstancePrivateEndpointConnectionsClientListByManagedInstanceResponse{}, err
	}
	return result, nil
}
