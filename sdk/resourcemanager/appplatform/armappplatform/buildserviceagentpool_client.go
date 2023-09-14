//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappplatform

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

// BuildServiceAgentPoolClient contains the methods for the BuildServiceAgentPool group.
// Don't use this type directly, use NewBuildServiceAgentPoolClient() instead.
type BuildServiceAgentPoolClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBuildServiceAgentPoolClient creates a new instance of BuildServiceAgentPoolClient with the specified values.
//   - subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBuildServiceAgentPoolClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BuildServiceAgentPoolClient, error) {
	cl, err := arm.NewClient(moduleName+".BuildServiceAgentPoolClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BuildServiceAgentPoolClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get build service agent pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - buildServiceName - The name of the build service resource.
//   - agentPoolName - The name of the build service agent pool resource.
//   - options - BuildServiceAgentPoolClientGetOptions contains the optional parameters for the BuildServiceAgentPoolClient.Get
//     method.
func (client *BuildServiceAgentPoolClient) Get(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, agentPoolName string, options *BuildServiceAgentPoolClientGetOptions) (BuildServiceAgentPoolClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, buildServiceName, agentPoolName, options)
	if err != nil {
		return BuildServiceAgentPoolClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BuildServiceAgentPoolClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BuildServiceAgentPoolClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BuildServiceAgentPoolClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, agentPoolName string, options *BuildServiceAgentPoolClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/agentPools/{agentPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if buildServiceName == "" {
		return nil, errors.New("parameter buildServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildServiceName}", url.PathEscape(buildServiceName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BuildServiceAgentPoolClient) getHandleResponse(resp *http.Response) (BuildServiceAgentPoolClientGetResponse, error) {
	result := BuildServiceAgentPoolClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BuildServiceAgentPoolResource); err != nil {
		return BuildServiceAgentPoolClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List build service agent pool.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - buildServiceName - The name of the build service resource.
//   - options - BuildServiceAgentPoolClientListOptions contains the optional parameters for the BuildServiceAgentPoolClient.NewListPager
//     method.
func (client *BuildServiceAgentPoolClient) NewListPager(resourceGroupName string, serviceName string, buildServiceName string, options *BuildServiceAgentPoolClientListOptions) *runtime.Pager[BuildServiceAgentPoolClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BuildServiceAgentPoolClientListResponse]{
		More: func(page BuildServiceAgentPoolClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BuildServiceAgentPoolClientListResponse) (BuildServiceAgentPoolClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, buildServiceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BuildServiceAgentPoolClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return BuildServiceAgentPoolClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BuildServiceAgentPoolClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BuildServiceAgentPoolClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, options *BuildServiceAgentPoolClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/agentPools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if buildServiceName == "" {
		return nil, errors.New("parameter buildServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildServiceName}", url.PathEscape(buildServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BuildServiceAgentPoolClient) listHandleResponse(resp *http.Response) (BuildServiceAgentPoolClientListResponse, error) {
	result := BuildServiceAgentPoolClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BuildServiceAgentPoolResourceCollection); err != nil {
		return BuildServiceAgentPoolClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdatePut - Create or update build service agent pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - buildServiceName - The name of the build service resource.
//   - agentPoolName - The name of the build service agent pool resource.
//   - agentPoolResource - Parameters for the update operation
//   - options - BuildServiceAgentPoolClientBeginUpdatePutOptions contains the optional parameters for the BuildServiceAgentPoolClient.BeginUpdatePut
//     method.
func (client *BuildServiceAgentPoolClient) BeginUpdatePut(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, agentPoolName string, agentPoolResource BuildServiceAgentPoolResource, options *BuildServiceAgentPoolClientBeginUpdatePutOptions) (*runtime.Poller[BuildServiceAgentPoolClientUpdatePutResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updatePut(ctx, resourceGroupName, serviceName, buildServiceName, agentPoolName, agentPoolResource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[BuildServiceAgentPoolClientUpdatePutResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[BuildServiceAgentPoolClientUpdatePutResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdatePut - Create or update build service agent pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
func (client *BuildServiceAgentPoolClient) updatePut(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, agentPoolName string, agentPoolResource BuildServiceAgentPoolResource, options *BuildServiceAgentPoolClientBeginUpdatePutOptions) (*http.Response, error) {
	req, err := client.updatePutCreateRequest(ctx, resourceGroupName, serviceName, buildServiceName, agentPoolName, agentPoolResource, options)
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

// updatePutCreateRequest creates the UpdatePut request.
func (client *BuildServiceAgentPoolClient) updatePutCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, agentPoolName string, agentPoolResource BuildServiceAgentPoolResource, options *BuildServiceAgentPoolClientBeginUpdatePutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/agentPools/{agentPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if buildServiceName == "" {
		return nil, errors.New("parameter buildServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{buildServiceName}", url.PathEscape(buildServiceName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, agentPoolResource)
}
