//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package armconnectedvmware

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// VirtualNetworksClient contains the methods for the VirtualNetworks group.
// Don't use this type directly, use NewVirtualNetworksClient() instead.
type VirtualNetworksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualNetworksClient creates a new instance of VirtualNetworksClient with the specified values.
//   - subscriptionID - The Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualNetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualNetworksClient, error) {
	cl, err := arm.NewClient(moduleName+".VirtualNetworksClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualNetworksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create Or Update virtual network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-10-preview
//   - resourceGroupName - The Resource Group Name.
//   - virtualNetworkName - Name of the virtual network resource.
//   - body - Request payload.
//   - options - VirtualNetworksClientBeginCreateOptions contains the optional parameters for the VirtualNetworksClient.BeginCreate
//     method.
func (client *VirtualNetworksClient) BeginCreate(ctx context.Context, resourceGroupName string, virtualNetworkName string, body VirtualNetwork, options *VirtualNetworksClientBeginCreateOptions) (*runtime.Poller[VirtualNetworksClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, virtualNetworkName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualNetworksClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualNetworksClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Create Or Update virtual network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-10-preview
func (client *VirtualNetworksClient) create(ctx context.Context, resourceGroupName string, virtualNetworkName string, body VirtualNetwork, options *VirtualNetworksClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, virtualNetworkName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *VirtualNetworksClient) createCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, body VirtualNetwork, options *VirtualNetworksClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualNetworks/{virtualNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Implements virtual network DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-10-preview
//   - resourceGroupName - The Resource Group Name.
//   - virtualNetworkName - Name of the virtual network resource.
//   - options - VirtualNetworksClientBeginDeleteOptions contains the optional parameters for the VirtualNetworksClient.BeginDelete
//     method.
func (client *VirtualNetworksClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksClientBeginDeleteOptions) (*runtime.Poller[VirtualNetworksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualNetworkName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualNetworksClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualNetworksClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Implements virtual network DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-10-preview
func (client *VirtualNetworksClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualNetworksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualNetworks/{virtualNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	if options != nil && options.Force != nil {
		reqQP.Set("force", strconv.FormatBool(*options.Force))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Implements virtual network GET method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-10-preview
//   - resourceGroupName - The Resource Group Name.
//   - virtualNetworkName - Name of the virtual network resource.
//   - options - VirtualNetworksClientGetOptions contains the optional parameters for the VirtualNetworksClient.Get method.
func (client *VirtualNetworksClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksClientGetOptions) (VirtualNetworksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
	if err != nil {
		return VirtualNetworksClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworksClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualNetworks/{virtualNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworksClient) getHandleResponse(resp *http.Response) (VirtualNetworksClientGetResponse, error) {
	result := VirtualNetworksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetwork); err != nil {
		return VirtualNetworksClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List of virtualNetworks in a subscription.
//
// Generated from API version 2022-01-10-preview
//   - options - VirtualNetworksClientListOptions contains the optional parameters for the VirtualNetworksClient.NewListPager
//     method.
func (client *VirtualNetworksClient) NewListPager(options *VirtualNetworksClientListOptions) *runtime.Pager[VirtualNetworksClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualNetworksClientListResponse]{
		More: func(page VirtualNetworksClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualNetworksClientListResponse) (VirtualNetworksClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualNetworksClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VirtualNetworksClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualNetworksClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualNetworksClient) listCreateRequest(ctx context.Context, options *VirtualNetworksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/virtualNetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualNetworksClient) listHandleResponse(resp *http.Response) (VirtualNetworksClientListResponse, error) {
	result := VirtualNetworksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworksList); err != nil {
		return VirtualNetworksClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List of virtualNetworks in a resource group.
//
// Generated from API version 2022-01-10-preview
//   - resourceGroupName - The Resource Group Name.
//   - options - VirtualNetworksClientListByResourceGroupOptions contains the optional parameters for the VirtualNetworksClient.NewListByResourceGroupPager
//     method.
func (client *VirtualNetworksClient) NewListByResourceGroupPager(resourceGroupName string, options *VirtualNetworksClientListByResourceGroupOptions) *runtime.Pager[VirtualNetworksClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualNetworksClientListByResourceGroupResponse]{
		More: func(page VirtualNetworksClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualNetworksClientListByResourceGroupResponse) (VirtualNetworksClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualNetworksClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VirtualNetworksClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualNetworksClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VirtualNetworksClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualNetworksClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualNetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VirtualNetworksClient) listByResourceGroupHandleResponse(resp *http.Response) (VirtualNetworksClientListByResourceGroupResponse, error) {
	result := VirtualNetworksClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworksList); err != nil {
		return VirtualNetworksClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - API to update certain properties of the virtual network resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-10-preview
//   - resourceGroupName - The Resource Group Name.
//   - virtualNetworkName - Name of the virtual network resource.
//   - body - Resource properties to update.
//   - options - VirtualNetworksClientUpdateOptions contains the optional parameters for the VirtualNetworksClient.Update method.
func (client *VirtualNetworksClient) Update(ctx context.Context, resourceGroupName string, virtualNetworkName string, body ResourcePatch, options *VirtualNetworksClientUpdateOptions) (VirtualNetworksClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, virtualNetworkName, body, options)
	if err != nil {
		return VirtualNetworksClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworksClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworksClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *VirtualNetworksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, body ResourcePatch, options *VirtualNetworksClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/virtualNetworks/{virtualNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleResponse handles the Update response.
func (client *VirtualNetworksClient) updateHandleResponse(resp *http.Response) (VirtualNetworksClientUpdateResponse, error) {
	result := VirtualNetworksClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetwork); err != nil {
		return VirtualNetworksClientUpdateResponse{}, err
	}
	return result, nil
}
