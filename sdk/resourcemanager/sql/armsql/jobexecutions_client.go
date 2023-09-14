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
	"strconv"
	"strings"
	"time"
)

// JobExecutionsClient contains the methods for the JobExecutions group.
// Don't use this type directly, use NewJobExecutionsClient() instead.
type JobExecutionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewJobExecutionsClient creates a new instance of JobExecutionsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewJobExecutionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*JobExecutionsClient, error) {
	cl, err := arm.NewClient(moduleName+".JobExecutionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &JobExecutionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Cancel - Requests cancellation of a job execution.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - jobAgentName - The name of the job agent.
//   - jobName - The name of the job.
//   - jobExecutionID - The id of the job execution to cancel.
//   - options - JobExecutionsClientCancelOptions contains the optional parameters for the JobExecutionsClient.Cancel method.
func (client *JobExecutionsClient) Cancel(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobExecutionsClientCancelOptions) (JobExecutionsClientCancelResponse, error) {
	var err error
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID, options)
	if err != nil {
		return JobExecutionsClientCancelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobExecutionsClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return JobExecutionsClientCancelResponse{}, err
	}
	return JobExecutionsClientCancelResponse{}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *JobExecutionsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobExecutionsClientCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions/{jobExecutionId}/cancel"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	urlPath = strings.ReplaceAll(urlPath, "{jobExecutionId}", url.PathEscape(jobExecutionID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginCreate - Starts an elastic job execution.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - jobAgentName - The name of the job agent.
//   - jobName - The name of the job to get.
//   - options - JobExecutionsClientBeginCreateOptions contains the optional parameters for the JobExecutionsClient.BeginCreate
//     method.
func (client *JobExecutionsClient) BeginCreate(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobExecutionsClientBeginCreateOptions) (*runtime.Poller[JobExecutionsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, serverName, jobAgentName, jobName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[JobExecutionsClientCreateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[JobExecutionsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Starts an elastic job execution.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
func (client *JobExecutionsClient) create(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobExecutionsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, options)
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

// createCreateRequest creates the Create request.
func (client *JobExecutionsClient) createCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobExecutionsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/start"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginCreateOrUpdate - Creates or updates a job execution.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - jobAgentName - The name of the job agent.
//   - jobName - The name of the job to get.
//   - jobExecutionID - The job execution id to create the job execution under.
//   - options - JobExecutionsClientBeginCreateOrUpdateOptions contains the optional parameters for the JobExecutionsClient.BeginCreateOrUpdate
//     method.
func (client *JobExecutionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobExecutionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[JobExecutionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[JobExecutionsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[JobExecutionsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a job execution.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
func (client *JobExecutionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobExecutionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *JobExecutionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobExecutionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions/{jobExecutionId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	urlPath = strings.ReplaceAll(urlPath, "{jobExecutionId}", url.PathEscape(jobExecutionID))
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
	return req, nil
}

// Get - Gets a job execution.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - jobAgentName - The name of the job agent.
//   - jobName - The name of the job.
//   - jobExecutionID - The id of the job execution
//   - options - JobExecutionsClientGetOptions contains the optional parameters for the JobExecutionsClient.Get method.
func (client *JobExecutionsClient) Get(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobExecutionsClientGetOptions) (JobExecutionsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID, options)
	if err != nil {
		return JobExecutionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobExecutionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return JobExecutionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *JobExecutionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *JobExecutionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions/{jobExecutionId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	urlPath = strings.ReplaceAll(urlPath, "{jobExecutionId}", url.PathEscape(jobExecutionID))
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
func (client *JobExecutionsClient) getHandleResponse(resp *http.Response) (JobExecutionsClientGetResponse, error) {
	result := JobExecutionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobExecution); err != nil {
		return JobExecutionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAgentPager - Lists all executions in a job agent.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - jobAgentName - The name of the job agent.
//   - options - JobExecutionsClientListByAgentOptions contains the optional parameters for the JobExecutionsClient.NewListByAgentPager
//     method.
func (client *JobExecutionsClient) NewListByAgentPager(resourceGroupName string, serverName string, jobAgentName string, options *JobExecutionsClientListByAgentOptions) *runtime.Pager[JobExecutionsClientListByAgentResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobExecutionsClientListByAgentResponse]{
		More: func(page JobExecutionsClientListByAgentResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobExecutionsClientListByAgentResponse) (JobExecutionsClientListByAgentResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAgentCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JobExecutionsClientListByAgentResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return JobExecutionsClientListByAgentResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobExecutionsClientListByAgentResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAgentHandleResponse(resp)
		},
	})
}

// listByAgentCreateRequest creates the ListByAgent request.
func (client *JobExecutionsClient) listByAgentCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, options *JobExecutionsClientListByAgentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/executions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.CreateTimeMin != nil {
		reqQP.Set("createTimeMin", options.CreateTimeMin.Format(time.RFC3339Nano))
	}
	if options != nil && options.CreateTimeMax != nil {
		reqQP.Set("createTimeMax", options.CreateTimeMax.Format(time.RFC3339Nano))
	}
	if options != nil && options.EndTimeMin != nil {
		reqQP.Set("endTimeMin", options.EndTimeMin.Format(time.RFC3339Nano))
	}
	if options != nil && options.EndTimeMax != nil {
		reqQP.Set("endTimeMax", options.EndTimeMax.Format(time.RFC3339Nano))
	}
	if options != nil && options.IsActive != nil {
		reqQP.Set("isActive", strconv.FormatBool(*options.IsActive))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAgentHandleResponse handles the ListByAgent response.
func (client *JobExecutionsClient) listByAgentHandleResponse(resp *http.Response) (JobExecutionsClientListByAgentResponse, error) {
	result := JobExecutionsClientListByAgentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobExecutionListResult); err != nil {
		return JobExecutionsClientListByAgentResponse{}, err
	}
	return result, nil
}

// NewListByJobPager - Lists a job's executions.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - jobAgentName - The name of the job agent.
//   - jobName - The name of the job to get.
//   - options - JobExecutionsClientListByJobOptions contains the optional parameters for the JobExecutionsClient.NewListByJobPager
//     method.
func (client *JobExecutionsClient) NewListByJobPager(resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobExecutionsClientListByJobOptions) *runtime.Pager[JobExecutionsClientListByJobResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobExecutionsClientListByJobResponse]{
		More: func(page JobExecutionsClientListByJobResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobExecutionsClientListByJobResponse) (JobExecutionsClientListByJobResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByJobCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return JobExecutionsClientListByJobResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return JobExecutionsClientListByJobResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return JobExecutionsClientListByJobResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByJobHandleResponse(resp)
		},
	})
}

// listByJobCreateRequest creates the ListByJob request.
func (client *JobExecutionsClient) listByJobCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobExecutionsClientListByJobOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.CreateTimeMin != nil {
		reqQP.Set("createTimeMin", options.CreateTimeMin.Format(time.RFC3339Nano))
	}
	if options != nil && options.CreateTimeMax != nil {
		reqQP.Set("createTimeMax", options.CreateTimeMax.Format(time.RFC3339Nano))
	}
	if options != nil && options.EndTimeMin != nil {
		reqQP.Set("endTimeMin", options.EndTimeMin.Format(time.RFC3339Nano))
	}
	if options != nil && options.EndTimeMax != nil {
		reqQP.Set("endTimeMax", options.EndTimeMax.Format(time.RFC3339Nano))
	}
	if options != nil && options.IsActive != nil {
		reqQP.Set("isActive", strconv.FormatBool(*options.IsActive))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByJobHandleResponse handles the ListByJob response.
func (client *JobExecutionsClient) listByJobHandleResponse(resp *http.Response) (JobExecutionsClientListByJobResponse, error) {
	result := JobExecutionsClientListByJobResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobExecutionListResult); err != nil {
		return JobExecutionsClientListByJobResponse{}, err
	}
	return result, nil
}
