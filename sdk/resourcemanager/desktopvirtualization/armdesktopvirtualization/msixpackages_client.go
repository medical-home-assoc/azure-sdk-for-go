//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdesktopvirtualization

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

// MSIXPackagesClient contains the methods for the MSIXPackages group.
// Don't use this type directly, use NewMSIXPackagesClient() instead.
type MSIXPackagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMSIXPackagesClient creates a new instance of MSIXPackagesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMSIXPackagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MSIXPackagesClient, error) {
	cl, err := arm.NewClient(moduleName+".MSIXPackagesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MSIXPackagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a MSIX package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-09
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - msixPackageFullName - The version specific package full name of the MSIX package within specified hostpool
//   - msixPackage - Object containing MSIX Package definitions.
//   - options - MSIXPackagesClientCreateOrUpdateOptions contains the optional parameters for the MSIXPackagesClient.CreateOrUpdate
//     method.
func (client *MSIXPackagesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hostPoolName string, msixPackageFullName string, msixPackage MSIXPackage, options *MSIXPackagesClientCreateOrUpdateOptions) (MSIXPackagesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hostPoolName, msixPackageFullName, msixPackage, options)
	if err != nil {
		return MSIXPackagesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MSIXPackagesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return MSIXPackagesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MSIXPackagesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, msixPackageFullName string, msixPackage MSIXPackage, options *MSIXPackagesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/msixPackages/{msixPackageFullName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	if msixPackageFullName == "" {
		return nil, errors.New("parameter msixPackageFullName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{msixPackageFullName}", url.PathEscape(msixPackageFullName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-09")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, msixPackage)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *MSIXPackagesClient) createOrUpdateHandleResponse(resp *http.Response) (MSIXPackagesClientCreateOrUpdateResponse, error) {
	result := MSIXPackagesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MSIXPackage); err != nil {
		return MSIXPackagesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Remove an MSIX Package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-09
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - msixPackageFullName - The version specific package full name of the MSIX package within specified hostpool
//   - options - MSIXPackagesClientDeleteOptions contains the optional parameters for the MSIXPackagesClient.Delete method.
func (client *MSIXPackagesClient) Delete(ctx context.Context, resourceGroupName string, hostPoolName string, msixPackageFullName string, options *MSIXPackagesClientDeleteOptions) (MSIXPackagesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hostPoolName, msixPackageFullName, options)
	if err != nil {
		return MSIXPackagesClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MSIXPackagesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return MSIXPackagesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return MSIXPackagesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MSIXPackagesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, msixPackageFullName string, options *MSIXPackagesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/msixPackages/{msixPackageFullName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	if msixPackageFullName == "" {
		return nil, errors.New("parameter msixPackageFullName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{msixPackageFullName}", url.PathEscape(msixPackageFullName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-09")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a msixpackage.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-09
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - msixPackageFullName - The version specific package full name of the MSIX package within specified hostpool
//   - options - MSIXPackagesClientGetOptions contains the optional parameters for the MSIXPackagesClient.Get method.
func (client *MSIXPackagesClient) Get(ctx context.Context, resourceGroupName string, hostPoolName string, msixPackageFullName string, options *MSIXPackagesClientGetOptions) (MSIXPackagesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, hostPoolName, msixPackageFullName, options)
	if err != nil {
		return MSIXPackagesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MSIXPackagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MSIXPackagesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MSIXPackagesClient) getCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, msixPackageFullName string, options *MSIXPackagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/msixPackages/{msixPackageFullName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	if msixPackageFullName == "" {
		return nil, errors.New("parameter msixPackageFullName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{msixPackageFullName}", url.PathEscape(msixPackageFullName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-09")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MSIXPackagesClient) getHandleResponse(resp *http.Response) (MSIXPackagesClientGetResponse, error) {
	result := MSIXPackagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MSIXPackage); err != nil {
		return MSIXPackagesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List MSIX packages in hostpool.
//
// Generated from API version 2022-09-09
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - options - MSIXPackagesClientListOptions contains the optional parameters for the MSIXPackagesClient.NewListPager method.
func (client *MSIXPackagesClient) NewListPager(resourceGroupName string, hostPoolName string, options *MSIXPackagesClientListOptions) *runtime.Pager[MSIXPackagesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MSIXPackagesClientListResponse]{
		More: func(page MSIXPackagesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MSIXPackagesClientListResponse) (MSIXPackagesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, hostPoolName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MSIXPackagesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return MSIXPackagesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MSIXPackagesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *MSIXPackagesClient) listCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *MSIXPackagesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/msixPackages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-09")
	if options != nil && options.PageSize != nil {
		reqQP.Set("pageSize", strconv.FormatInt(int64(*options.PageSize), 10))
	}
	if options != nil && options.IsDescending != nil {
		reqQP.Set("isDescending", strconv.FormatBool(*options.IsDescending))
	}
	if options != nil && options.InitialSkip != nil {
		reqQP.Set("initialSkip", strconv.FormatInt(int64(*options.InitialSkip), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MSIXPackagesClient) listHandleResponse(resp *http.Response) (MSIXPackagesClientListResponse, error) {
	result := MSIXPackagesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MSIXPackageList); err != nil {
		return MSIXPackagesClientListResponse{}, err
	}
	return result, nil
}

// Update - Update an MSIX Package.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-09
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - msixPackageFullName - The version specific package full name of the MSIX package within specified hostpool
//   - options - MSIXPackagesClientUpdateOptions contains the optional parameters for the MSIXPackagesClient.Update method.
func (client *MSIXPackagesClient) Update(ctx context.Context, resourceGroupName string, hostPoolName string, msixPackageFullName string, options *MSIXPackagesClientUpdateOptions) (MSIXPackagesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, hostPoolName, msixPackageFullName, options)
	if err != nil {
		return MSIXPackagesClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MSIXPackagesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MSIXPackagesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *MSIXPackagesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, msixPackageFullName string, options *MSIXPackagesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/msixPackages/{msixPackageFullName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	if msixPackageFullName == "" {
		return nil, errors.New("parameter msixPackageFullName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{msixPackageFullName}", url.PathEscape(msixPackageFullName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-09")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.MsixPackage != nil {
		return req, runtime.MarshalAsJSON(req, *options.MsixPackage)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *MSIXPackagesClient) updateHandleResponse(resp *http.Response) (MSIXPackagesClientUpdateResponse, error) {
	result := MSIXPackagesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MSIXPackage); err != nil {
		return MSIXPackagesClientUpdateResponse{}, err
	}
	return result, nil
}
