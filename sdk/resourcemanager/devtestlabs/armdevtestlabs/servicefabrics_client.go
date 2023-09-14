//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevtestlabs

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

// ServiceFabricsClient contains the methods for the ServiceFabrics group.
// Don't use this type directly, use NewServiceFabricsClient() instead.
type ServiceFabricsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServiceFabricsClient creates a new instance of ServiceFabricsClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServiceFabricsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServiceFabricsClient, error) {
	cl, err := arm.NewClient(moduleName+".ServiceFabricsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServiceFabricsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or replace an existing service fabric. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - name - The name of the service fabric.
//   - serviceFabric - A Service Fabric.
//   - options - ServiceFabricsClientBeginCreateOrUpdateOptions contains the optional parameters for the ServiceFabricsClient.BeginCreateOrUpdate
//     method.
func (client *ServiceFabricsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, name string, serviceFabric ServiceFabric, options *ServiceFabricsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServiceFabricsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, labName, userName, name, serviceFabric, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServiceFabricsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServiceFabricsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or replace an existing service fabric. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
func (client *ServiceFabricsClient) createOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, name string, serviceFabric ServiceFabric, options *ServiceFabricsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, userName, name, serviceFabric, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServiceFabricsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, serviceFabric ServiceFabric, options *ServiceFabricsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, serviceFabric)
}

// BeginDelete - Delete service fabric. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - name - The name of the service fabric.
//   - options - ServiceFabricsClientBeginDeleteOptions contains the optional parameters for the ServiceFabricsClient.BeginDelete
//     method.
func (client *ServiceFabricsClient) BeginDelete(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *ServiceFabricsClientBeginDeleteOptions) (*runtime.Poller[ServiceFabricsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, labName, userName, name, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServiceFabricsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServiceFabricsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete service fabric. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
func (client *ServiceFabricsClient) deleteOperation(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *ServiceFabricsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, userName, name, options)
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
func (client *ServiceFabricsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *ServiceFabricsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get service fabric.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - name - The name of the service fabric.
//   - options - ServiceFabricsClientGetOptions contains the optional parameters for the ServiceFabricsClient.Get method.
func (client *ServiceFabricsClient) Get(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *ServiceFabricsClientGetOptions) (ServiceFabricsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, userName, name, options)
	if err != nil {
		return ServiceFabricsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceFabricsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceFabricsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServiceFabricsClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *ServiceFabricsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceFabricsClient) getHandleResponse(resp *http.Response) (ServiceFabricsClientGetResponse, error) {
	result := ServiceFabricsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceFabric); err != nil {
		return ServiceFabricsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List service fabrics in a given user profile.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - options - ServiceFabricsClientListOptions contains the optional parameters for the ServiceFabricsClient.NewListPager method.
func (client *ServiceFabricsClient) NewListPager(resourceGroupName string, labName string, userName string, options *ServiceFabricsClientListOptions) *runtime.Pager[ServiceFabricsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServiceFabricsClientListResponse]{
		More: func(page ServiceFabricsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServiceFabricsClientListResponse) (ServiceFabricsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, labName, userName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServiceFabricsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServiceFabricsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServiceFabricsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ServiceFabricsClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, options *ServiceFabricsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServiceFabricsClient) listHandleResponse(resp *http.Response) (ServiceFabricsClientListResponse, error) {
	result := ServiceFabricsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceFabricList); err != nil {
		return ServiceFabricsClientListResponse{}, err
	}
	return result, nil
}

// ListApplicableSchedules - Lists the applicable start/stop schedules, if any.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - name - The name of the service fabric.
//   - options - ServiceFabricsClientListApplicableSchedulesOptions contains the optional parameters for the ServiceFabricsClient.ListApplicableSchedules
//     method.
func (client *ServiceFabricsClient) ListApplicableSchedules(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *ServiceFabricsClientListApplicableSchedulesOptions) (ServiceFabricsClientListApplicableSchedulesResponse, error) {
	req, err := client.listApplicableSchedulesCreateRequest(ctx, resourceGroupName, labName, userName, name, options)
	if err != nil {
		return ServiceFabricsClientListApplicableSchedulesResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceFabricsClientListApplicableSchedulesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceFabricsClientListApplicableSchedulesResponse{}, runtime.NewResponseError(resp)
	}
	return client.listApplicableSchedulesHandleResponse(resp)
}

// listApplicableSchedulesCreateRequest creates the ListApplicableSchedules request.
func (client *ServiceFabricsClient) listApplicableSchedulesCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *ServiceFabricsClientListApplicableSchedulesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}/listApplicableSchedules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listApplicableSchedulesHandleResponse handles the ListApplicableSchedules response.
func (client *ServiceFabricsClient) listApplicableSchedulesHandleResponse(resp *http.Response) (ServiceFabricsClientListApplicableSchedulesResponse, error) {
	result := ServiceFabricsClientListApplicableSchedulesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicableSchedule); err != nil {
		return ServiceFabricsClientListApplicableSchedulesResponse{}, err
	}
	return result, nil
}

// BeginStart - Start a service fabric. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - name - The name of the service fabric.
//   - options - ServiceFabricsClientBeginStartOptions contains the optional parameters for the ServiceFabricsClient.BeginStart
//     method.
func (client *ServiceFabricsClient) BeginStart(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *ServiceFabricsClientBeginStartOptions) (*runtime.Poller[ServiceFabricsClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, labName, userName, name, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServiceFabricsClientStartResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServiceFabricsClientStartResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Start - Start a service fabric. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
func (client *ServiceFabricsClient) start(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *ServiceFabricsClientBeginStartOptions) (*http.Response, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, labName, userName, name, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// startCreateRequest creates the Start request.
func (client *ServiceFabricsClient) startCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *ServiceFabricsClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStop - Stop a service fabric This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - name - The name of the service fabric.
//   - options - ServiceFabricsClientBeginStopOptions contains the optional parameters for the ServiceFabricsClient.BeginStop
//     method.
func (client *ServiceFabricsClient) BeginStop(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *ServiceFabricsClientBeginStopOptions) (*runtime.Poller[ServiceFabricsClientStopResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stop(ctx, resourceGroupName, labName, userName, name, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServiceFabricsClientStopResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServiceFabricsClientStopResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Stop - Stop a service fabric This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
func (client *ServiceFabricsClient) stop(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *ServiceFabricsClientBeginStopOptions) (*http.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, labName, userName, name, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// stopCreateRequest creates the Stop request.
func (client *ServiceFabricsClient) stopCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *ServiceFabricsClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Update - Allows modifying tags of service fabrics. All other properties will be ignored.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - name - The name of the service fabric.
//   - serviceFabric - A Service Fabric.
//   - options - ServiceFabricsClientUpdateOptions contains the optional parameters for the ServiceFabricsClient.Update method.
func (client *ServiceFabricsClient) Update(ctx context.Context, resourceGroupName string, labName string, userName string, name string, serviceFabric ServiceFabricFragment, options *ServiceFabricsClientUpdateOptions) (ServiceFabricsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labName, userName, name, serviceFabric, options)
	if err != nil {
		return ServiceFabricsClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceFabricsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceFabricsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ServiceFabricsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, serviceFabric ServiceFabricFragment, options *ServiceFabricsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, serviceFabric)
}

// updateHandleResponse handles the Update response.
func (client *ServiceFabricsClient) updateHandleResponse(resp *http.Response) (ServiceFabricsClientUpdateResponse, error) {
	result := ServiceFabricsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceFabric); err != nil {
		return ServiceFabricsClientUpdateResponse{}, err
	}
	return result, nil
}
