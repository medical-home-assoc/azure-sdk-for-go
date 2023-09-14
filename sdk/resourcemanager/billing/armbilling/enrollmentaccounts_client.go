//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbilling

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

// EnrollmentAccountsClient contains the methods for the EnrollmentAccounts group.
// Don't use this type directly, use NewEnrollmentAccountsClient() instead.
type EnrollmentAccountsClient struct {
	internal *arm.Client
}

// NewEnrollmentAccountsClient creates a new instance of EnrollmentAccountsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEnrollmentAccountsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*EnrollmentAccountsClient, error) {
	cl, err := arm.NewClient(moduleName+".EnrollmentAccountsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EnrollmentAccountsClient{
		internal: cl,
	}
	return client, nil
}

// Get - Gets a enrollment account by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-03-01-preview
//   - name - Enrollment Account name.
//   - options - EnrollmentAccountsClientGetOptions contains the optional parameters for the EnrollmentAccountsClient.Get method.
func (client *EnrollmentAccountsClient) Get(ctx context.Context, name string, options *EnrollmentAccountsClientGetOptions) (EnrollmentAccountsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, name, options)
	if err != nil {
		return EnrollmentAccountsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnrollmentAccountsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EnrollmentAccountsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EnrollmentAccountsClient) getCreateRequest(ctx context.Context, name string, options *EnrollmentAccountsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/enrollmentAccounts/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EnrollmentAccountsClient) getHandleResponse(resp *http.Response) (EnrollmentAccountsClientGetResponse, error) {
	result := EnrollmentAccountsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnrollmentAccountSummary); err != nil {
		return EnrollmentAccountsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the enrollment accounts the caller has access to.
//
// Generated from API version 2018-03-01-preview
//   - options - EnrollmentAccountsClientListOptions contains the optional parameters for the EnrollmentAccountsClient.NewListPager
//     method.
func (client *EnrollmentAccountsClient) NewListPager(options *EnrollmentAccountsClientListOptions) *runtime.Pager[EnrollmentAccountsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EnrollmentAccountsClientListResponse]{
		More: func(page EnrollmentAccountsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EnrollmentAccountsClientListResponse) (EnrollmentAccountsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EnrollmentAccountsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return EnrollmentAccountsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EnrollmentAccountsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *EnrollmentAccountsClient) listCreateRequest(ctx context.Context, options *EnrollmentAccountsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/enrollmentAccounts"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EnrollmentAccountsClient) listHandleResponse(resp *http.Response) (EnrollmentAccountsClientListResponse, error) {
	result := EnrollmentAccountsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnrollmentAccountListResult); err != nil {
		return EnrollmentAccountsClientListResponse{}, err
	}
	return result, nil
}
