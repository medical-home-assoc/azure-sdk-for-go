//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armproviderhub

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

// CustomRolloutsClient contains the methods for the CustomRollouts group.
// Don't use this type directly, use NewCustomRolloutsClient() instead.
type CustomRolloutsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCustomRolloutsClient creates a new instance of CustomRolloutsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCustomRolloutsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CustomRolloutsClient, error) {
	cl, err := arm.NewClient(moduleName+".CustomRolloutsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CustomRolloutsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the rollout details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-20
//   - providerNamespace - The name of the resource provider hosted within ProviderHub.
//   - rolloutName - The rollout name.
//   - properties - The custom rollout properties supplied to the CreateOrUpdate operation.
//   - options - CustomRolloutsClientCreateOrUpdateOptions contains the optional parameters for the CustomRolloutsClient.CreateOrUpdate
//     method.
func (client *CustomRolloutsClient) CreateOrUpdate(ctx context.Context, providerNamespace string, rolloutName string, properties CustomRollout, options *CustomRolloutsClientCreateOrUpdateOptions) (CustomRolloutsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, providerNamespace, rolloutName, properties, options)
	if err != nil {
		return CustomRolloutsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomRolloutsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomRolloutsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CustomRolloutsClient) createOrUpdateCreateRequest(ctx context.Context, providerNamespace string, rolloutName string, properties CustomRollout, options *CustomRolloutsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/customRollouts/{rolloutName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	if rolloutName == "" {
		return nil, errors.New("parameter rolloutName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rolloutName}", url.PathEscape(rolloutName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, properties)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CustomRolloutsClient) createOrUpdateHandleResponse(resp *http.Response) (CustomRolloutsClientCreateOrUpdateResponse, error) {
	result := CustomRolloutsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomRollout); err != nil {
		return CustomRolloutsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Gets the custom rollout details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-20
//   - providerNamespace - The name of the resource provider hosted within ProviderHub.
//   - rolloutName - The rollout name.
//   - options - CustomRolloutsClientGetOptions contains the optional parameters for the CustomRolloutsClient.Get method.
func (client *CustomRolloutsClient) Get(ctx context.Context, providerNamespace string, rolloutName string, options *CustomRolloutsClientGetOptions) (CustomRolloutsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, providerNamespace, rolloutName, options)
	if err != nil {
		return CustomRolloutsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomRolloutsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomRolloutsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CustomRolloutsClient) getCreateRequest(ctx context.Context, providerNamespace string, rolloutName string, options *CustomRolloutsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/customRollouts/{rolloutName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	if rolloutName == "" {
		return nil, errors.New("parameter rolloutName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{rolloutName}", url.PathEscape(rolloutName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomRolloutsClient) getHandleResponse(resp *http.Response) (CustomRolloutsClientGetResponse, error) {
	result := CustomRolloutsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomRollout); err != nil {
		return CustomRolloutsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByProviderRegistrationPager - Gets the list of the custom rollouts for the given provider.
//
// Generated from API version 2020-11-20
//   - providerNamespace - The name of the resource provider hosted within ProviderHub.
//   - options - CustomRolloutsClientListByProviderRegistrationOptions contains the optional parameters for the CustomRolloutsClient.NewListByProviderRegistrationPager
//     method.
func (client *CustomRolloutsClient) NewListByProviderRegistrationPager(providerNamespace string, options *CustomRolloutsClientListByProviderRegistrationOptions) *runtime.Pager[CustomRolloutsClientListByProviderRegistrationResponse] {
	return runtime.NewPager(runtime.PagingHandler[CustomRolloutsClientListByProviderRegistrationResponse]{
		More: func(page CustomRolloutsClientListByProviderRegistrationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CustomRolloutsClientListByProviderRegistrationResponse) (CustomRolloutsClientListByProviderRegistrationResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByProviderRegistrationCreateRequest(ctx, providerNamespace, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CustomRolloutsClientListByProviderRegistrationResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CustomRolloutsClientListByProviderRegistrationResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CustomRolloutsClientListByProviderRegistrationResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByProviderRegistrationHandleResponse(resp)
		},
	})
}

// listByProviderRegistrationCreateRequest creates the ListByProviderRegistration request.
func (client *CustomRolloutsClient) listByProviderRegistrationCreateRequest(ctx context.Context, providerNamespace string, options *CustomRolloutsClientListByProviderRegistrationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/customRollouts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProviderRegistrationHandleResponse handles the ListByProviderRegistration response.
func (client *CustomRolloutsClient) listByProviderRegistrationHandleResponse(resp *http.Response) (CustomRolloutsClientListByProviderRegistrationResponse, error) {
	result := CustomRolloutsClientListByProviderRegistrationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomRolloutArrayResponseWithContinuation); err != nil {
		return CustomRolloutsClientListByProviderRegistrationResponse{}, err
	}
	return result, nil
}
