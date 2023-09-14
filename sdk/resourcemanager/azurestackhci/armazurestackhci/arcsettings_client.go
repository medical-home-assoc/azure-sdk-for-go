//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armazurestackhci

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

// ArcSettingsClient contains the methods for the ArcSettings group.
// Don't use this type directly, use NewArcSettingsClient() instead.
type ArcSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewArcSettingsClient creates a new instance of ArcSettingsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewArcSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ArcSettingsClient, error) {
	cl, err := arm.NewClient(moduleName+".ArcSettingsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ArcSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create ArcSetting for HCI cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - arcSettingName - The name of the proxy resource holding details of HCI ArcSetting information.
//   - arcSetting - Parameters supplied to the Create ArcSetting resource for this HCI cluster.
//   - options - ArcSettingsClientCreateOptions contains the optional parameters for the ArcSettingsClient.Create method.
func (client *ArcSettingsClient) Create(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, arcSetting ArcSetting, options *ArcSettingsClientCreateOptions) (ArcSettingsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, arcSetting, options)
	if err != nil {
		return ArcSettingsClientCreateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArcSettingsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArcSettingsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *ArcSettingsClient) createCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, arcSetting ArcSetting, options *ArcSettingsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, arcSetting)
}

// createHandleResponse handles the Create response.
func (client *ArcSettingsClient) createHandleResponse(resp *http.Response) (ArcSettingsClientCreateResponse, error) {
	result := ArcSettingsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArcSetting); err != nil {
		return ArcSettingsClientCreateResponse{}, err
	}
	return result, nil
}

// BeginCreateIdentity - Create Aad identity for arc settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - arcSettingName - The name of the proxy resource holding details of HCI ArcSetting information.
//   - options - ArcSettingsClientBeginCreateIdentityOptions contains the optional parameters for the ArcSettingsClient.BeginCreateIdentity
//     method.
func (client *ArcSettingsClient) BeginCreateIdentity(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, options *ArcSettingsClientBeginCreateIdentityOptions) (*runtime.Poller[ArcSettingsClientCreateIdentityResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createIdentity(ctx, resourceGroupName, clusterName, arcSettingName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ArcSettingsClientCreateIdentityResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ArcSettingsClientCreateIdentityResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateIdentity - Create Aad identity for arc settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01
func (client *ArcSettingsClient) createIdentity(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, options *ArcSettingsClientBeginCreateIdentityOptions) (*http.Response, error) {
	req, err := client.createIdentityCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, options)
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

// createIdentityCreateRequest creates the CreateIdentity request.
func (client *ArcSettingsClient) createIdentityCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, options *ArcSettingsClientBeginCreateIdentityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}/createArcIdentity"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDelete - Delete ArcSetting resource details of HCI Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - arcSettingName - The name of the proxy resource holding details of HCI ArcSetting information.
//   - options - ArcSettingsClientBeginDeleteOptions contains the optional parameters for the ArcSettingsClient.BeginDelete method.
func (client *ArcSettingsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, options *ArcSettingsClientBeginDeleteOptions) (*runtime.Poller[ArcSettingsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, arcSettingName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ArcSettingsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ArcSettingsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete ArcSetting resource details of HCI Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01
func (client *ArcSettingsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, options *ArcSettingsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, options)
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
func (client *ArcSettingsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, options *ArcSettingsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GeneratePassword - Generate password for arc settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - arcSettingName - The name of the proxy resource holding details of HCI ArcSetting information.
//   - options - ArcSettingsClientGeneratePasswordOptions contains the optional parameters for the ArcSettingsClient.GeneratePassword
//     method.
func (client *ArcSettingsClient) GeneratePassword(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, options *ArcSettingsClientGeneratePasswordOptions) (ArcSettingsClientGeneratePasswordResponse, error) {
	req, err := client.generatePasswordCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, options)
	if err != nil {
		return ArcSettingsClientGeneratePasswordResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArcSettingsClientGeneratePasswordResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArcSettingsClientGeneratePasswordResponse{}, runtime.NewResponseError(resp)
	}
	return client.generatePasswordHandleResponse(resp)
}

// generatePasswordCreateRequest creates the GeneratePassword request.
func (client *ArcSettingsClient) generatePasswordCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, options *ArcSettingsClientGeneratePasswordOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}/generatePassword"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// generatePasswordHandleResponse handles the GeneratePassword response.
func (client *ArcSettingsClient) generatePasswordHandleResponse(resp *http.Response) (ArcSettingsClientGeneratePasswordResponse, error) {
	result := ArcSettingsClientGeneratePasswordResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PasswordCredential); err != nil {
		return ArcSettingsClientGeneratePasswordResponse{}, err
	}
	return result, nil
}

// Get - Get ArcSetting resource details of HCI Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - arcSettingName - The name of the proxy resource holding details of HCI ArcSetting information.
//   - options - ArcSettingsClientGetOptions contains the optional parameters for the ArcSettingsClient.Get method.
func (client *ArcSettingsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, options *ArcSettingsClientGetOptions) (ArcSettingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, options)
	if err != nil {
		return ArcSettingsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArcSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArcSettingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ArcSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, options *ArcSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ArcSettingsClient) getHandleResponse(resp *http.Response) (ArcSettingsClientGetResponse, error) {
	result := ArcSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArcSetting); err != nil {
		return ArcSettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByClusterPager - Get ArcSetting resources of HCI Cluster.
//
// Generated from API version 2022-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - options - ArcSettingsClientListByClusterOptions contains the optional parameters for the ArcSettingsClient.NewListByClusterPager
//     method.
func (client *ArcSettingsClient) NewListByClusterPager(resourceGroupName string, clusterName string, options *ArcSettingsClientListByClusterOptions) *runtime.Pager[ArcSettingsClientListByClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[ArcSettingsClientListByClusterResponse]{
		More: func(page ArcSettingsClientListByClusterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ArcSettingsClientListByClusterResponse) (ArcSettingsClientListByClusterResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByClusterCreateRequest(ctx, resourceGroupName, clusterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ArcSettingsClientListByClusterResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ArcSettingsClientListByClusterResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ArcSettingsClientListByClusterResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByClusterHandleResponse(resp)
		},
	})
}

// listByClusterCreateRequest creates the ListByCluster request.
func (client *ArcSettingsClient) listByClusterCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ArcSettingsClientListByClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByClusterHandleResponse handles the ListByCluster response.
func (client *ArcSettingsClient) listByClusterHandleResponse(resp *http.Response) (ArcSettingsClientListByClusterResponse, error) {
	result := ArcSettingsClientListByClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArcSettingList); err != nil {
		return ArcSettingsClientListByClusterResponse{}, err
	}
	return result, nil
}

// Update - Update ArcSettings for HCI cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - arcSettingName - The name of the proxy resource holding details of HCI ArcSetting information.
//   - arcSetting - ArcSettings parameters that needs to be updated
//   - options - ArcSettingsClientUpdateOptions contains the optional parameters for the ArcSettingsClient.Update method.
func (client *ArcSettingsClient) Update(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, arcSetting ArcSettingsPatch, options *ArcSettingsClientUpdateOptions) (ArcSettingsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, arcSettingName, arcSetting, options)
	if err != nil {
		return ArcSettingsClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArcSettingsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArcSettingsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ArcSettingsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, arcSettingName string, arcSetting ArcSettingsPatch, options *ArcSettingsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/clusters/{clusterName}/arcSettings/{arcSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if arcSettingName == "" {
		return nil, errors.New("parameter arcSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{arcSettingName}", url.PathEscape(arcSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, arcSetting)
}

// updateHandleResponse handles the Update response.
func (client *ArcSettingsClient) updateHandleResponse(resp *http.Response) (ArcSettingsClientUpdateResponse, error) {
	result := ArcSettingsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArcSetting); err != nil {
		return ArcSettingsClientUpdateResponse{}, err
	}
	return result, nil
}
