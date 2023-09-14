//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurityinsights

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

// IncidentsClient contains the methods for the Incidents group.
// Don't use this type directly, use NewIncidentsClient() instead.
type IncidentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIncidentsClient creates a new instance of IncidentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIncidentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IncidentsClient, error) {
	cl, err := arm.NewClient(moduleName+".IncidentsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IncidentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an incident.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - incidentID - Incident ID
//   - incident - The incident
//   - options - IncidentsClientCreateOrUpdateOptions contains the optional parameters for the IncidentsClient.CreateOrUpdate
//     method.
func (client *IncidentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incident Incident, options *IncidentsClientCreateOrUpdateOptions) (IncidentsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, incident, options)
	if err != nil {
		return IncidentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IncidentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return IncidentsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IncidentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incident Incident, options *IncidentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, incident)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IncidentsClient) createOrUpdateHandleResponse(resp *http.Response) (IncidentsClientCreateOrUpdateResponse, error) {
	result := IncidentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Incident); err != nil {
		return IncidentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a given incident.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - incidentID - Incident ID
//   - options - IncidentsClientDeleteOptions contains the optional parameters for the IncidentsClient.Delete method.
func (client *IncidentsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientDeleteOptions) (IncidentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
	if err != nil {
		return IncidentsClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IncidentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IncidentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return IncidentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IncidentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a given incident.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - incidentID - Incident ID
//   - options - IncidentsClientGetOptions contains the optional parameters for the IncidentsClient.Get method.
func (client *IncidentsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientGetOptions) (IncidentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
	if err != nil {
		return IncidentsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IncidentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IncidentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IncidentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IncidentsClient) getHandleResponse(resp *http.Response) (IncidentsClientGetResponse, error) {
	result := IncidentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Incident); err != nil {
		return IncidentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all incidents.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - IncidentsClientListOptions contains the optional parameters for the IncidentsClient.NewListPager method.
func (client *IncidentsClient) NewListPager(resourceGroupName string, workspaceName string, options *IncidentsClientListOptions) *runtime.Pager[IncidentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IncidentsClientListResponse]{
		More: func(page IncidentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IncidentsClientListResponse) (IncidentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IncidentsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return IncidentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IncidentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *IncidentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *IncidentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IncidentsClient) listHandleResponse(resp *http.Response) (IncidentsClientListResponse, error) {
	result := IncidentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentList); err != nil {
		return IncidentsClientListResponse{}, err
	}
	return result, nil
}

// ListAlerts - Gets all alerts for an incident.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - incidentID - Incident ID
//   - options - IncidentsClientListAlertsOptions contains the optional parameters for the IncidentsClient.ListAlerts method.
func (client *IncidentsClient) ListAlerts(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListAlertsOptions) (IncidentsClientListAlertsResponse, error) {
	req, err := client.listAlertsCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
	if err != nil {
		return IncidentsClientListAlertsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IncidentsClientListAlertsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IncidentsClientListAlertsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listAlertsHandleResponse(resp)
}

// listAlertsCreateRequest creates the ListAlerts request.
func (client *IncidentsClient) listAlertsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListAlertsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/alerts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAlertsHandleResponse handles the ListAlerts response.
func (client *IncidentsClient) listAlertsHandleResponse(resp *http.Response) (IncidentsClientListAlertsResponse, error) {
	result := IncidentsClientListAlertsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentAlertList); err != nil {
		return IncidentsClientListAlertsResponse{}, err
	}
	return result, nil
}

// ListBookmarks - Gets all bookmarks for an incident.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - incidentID - Incident ID
//   - options - IncidentsClientListBookmarksOptions contains the optional parameters for the IncidentsClient.ListBookmarks method.
func (client *IncidentsClient) ListBookmarks(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListBookmarksOptions) (IncidentsClientListBookmarksResponse, error) {
	req, err := client.listBookmarksCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
	if err != nil {
		return IncidentsClientListBookmarksResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IncidentsClientListBookmarksResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IncidentsClientListBookmarksResponse{}, runtime.NewResponseError(resp)
	}
	return client.listBookmarksHandleResponse(resp)
}

// listBookmarksCreateRequest creates the ListBookmarks request.
func (client *IncidentsClient) listBookmarksCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListBookmarksOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/bookmarks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBookmarksHandleResponse handles the ListBookmarks response.
func (client *IncidentsClient) listBookmarksHandleResponse(resp *http.Response) (IncidentsClientListBookmarksResponse, error) {
	result := IncidentsClientListBookmarksResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentBookmarkList); err != nil {
		return IncidentsClientListBookmarksResponse{}, err
	}
	return result, nil
}

// ListEntities - Gets all entities for an incident.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - incidentID - Incident ID
//   - options - IncidentsClientListEntitiesOptions contains the optional parameters for the IncidentsClient.ListEntities method.
func (client *IncidentsClient) ListEntities(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListEntitiesOptions) (IncidentsClientListEntitiesResponse, error) {
	req, err := client.listEntitiesCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
	if err != nil {
		return IncidentsClientListEntitiesResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IncidentsClientListEntitiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IncidentsClientListEntitiesResponse{}, runtime.NewResponseError(resp)
	}
	return client.listEntitiesHandleResponse(resp)
}

// listEntitiesCreateRequest creates the ListEntities request.
func (client *IncidentsClient) listEntitiesCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListEntitiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/entities"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listEntitiesHandleResponse handles the ListEntities response.
func (client *IncidentsClient) listEntitiesHandleResponse(resp *http.Response) (IncidentsClientListEntitiesResponse, error) {
	result := IncidentsClientListEntitiesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentEntitiesResponse); err != nil {
		return IncidentsClientListEntitiesResponse{}, err
	}
	return result, nil
}
