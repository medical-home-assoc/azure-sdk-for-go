//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armkusto

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

// OperationsResultsLocationClient contains the methods for the OperationsResultsLocation group.
// Don't use this type directly, use NewOperationsResultsLocationClient() instead.
type OperationsResultsLocationClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOperationsResultsLocationClient creates a new instance of OperationsResultsLocationClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOperationsResultsLocationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationsResultsLocationClient, error) {
	cl, err := arm.NewClient(moduleName+".OperationsResultsLocationClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OperationsResultsLocationClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Returns operation results.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-02
//   - location - The name of Azure region.
//   - operationID - The ID of an ongoing async operation.
//   - options - OperationsResultsLocationClientGetOptions contains the optional parameters for the OperationsResultsLocationClient.Get
//     method.
func (client *OperationsResultsLocationClient) Get(ctx context.Context, location string, operationID string, options *OperationsResultsLocationClientGetOptions) (OperationsResultsLocationClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, operationID, options)
	if err != nil {
		return OperationsResultsLocationClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsResultsLocationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return OperationsResultsLocationClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OperationsResultsLocationClient) getCreateRequest(ctx context.Context, location string, operationID string, options *OperationsResultsLocationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Kusto/locations/{location}/operationResults/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OperationsResultsLocationClient) getHandleResponse(resp *http.Response) (OperationsResultsLocationClientGetResponse, error) {
	result := OperationsResultsLocationClientGetResponse{}
	if val := resp.Header.Get("Azure-AsyncOperation"); val != "" {
		result.AzureAsyncOperation = &val
	}
	return result, nil
}
