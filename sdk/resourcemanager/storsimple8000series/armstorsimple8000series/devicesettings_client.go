//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstorsimple8000series

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// DeviceSettingsClient contains the methods for the DeviceSettings group.
// Don't use this type directly, use NewDeviceSettingsClient() instead.
type DeviceSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDeviceSettingsClient creates a new instance of DeviceSettingsClient with the specified values.
//   - subscriptionID - The subscription id
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDeviceSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DeviceSettingsClient, error) {
	cl, err := arm.NewClient(moduleName+".DeviceSettingsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DeviceSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdateAlertSettings - Creates or updates the alert settings of the specified device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
//   - deviceName - The device name
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - parameters - The alert settings to be added or updated.
//   - options - DeviceSettingsClientBeginCreateOrUpdateAlertSettingsOptions contains the optional parameters for the DeviceSettingsClient.BeginCreateOrUpdateAlertSettings
//     method.
func (client *DeviceSettingsClient) BeginCreateOrUpdateAlertSettings(ctx context.Context, deviceName string, resourceGroupName string, managerName string, parameters AlertSettings, options *DeviceSettingsClientBeginCreateOrUpdateAlertSettingsOptions) (*runtime.Poller[DeviceSettingsClientCreateOrUpdateAlertSettingsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateAlertSettings(ctx, deviceName, resourceGroupName, managerName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DeviceSettingsClientCreateOrUpdateAlertSettingsResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[DeviceSettingsClientCreateOrUpdateAlertSettingsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdateAlertSettings - Creates or updates the alert settings of the specified device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
func (client *DeviceSettingsClient) createOrUpdateAlertSettings(ctx context.Context, deviceName string, resourceGroupName string, managerName string, parameters AlertSettings, options *DeviceSettingsClientBeginCreateOrUpdateAlertSettingsOptions) (*http.Response, error) {
	req, err := client.createOrUpdateAlertSettingsCreateRequest(ctx, deviceName, resourceGroupName, managerName, parameters, options)
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

// createOrUpdateAlertSettingsCreateRequest creates the CreateOrUpdateAlertSettings request.
func (client *DeviceSettingsClient) createOrUpdateAlertSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, parameters AlertSettings, options *DeviceSettingsClientBeginCreateOrUpdateAlertSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/alertSettings/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginCreateOrUpdateTimeSettings - Creates or updates the time settings of the specified device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
//   - deviceName - The device name
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - parameters - The time settings to be added or updated.
//   - options - DeviceSettingsClientBeginCreateOrUpdateTimeSettingsOptions contains the optional parameters for the DeviceSettingsClient.BeginCreateOrUpdateTimeSettings
//     method.
func (client *DeviceSettingsClient) BeginCreateOrUpdateTimeSettings(ctx context.Context, deviceName string, resourceGroupName string, managerName string, parameters TimeSettings, options *DeviceSettingsClientBeginCreateOrUpdateTimeSettingsOptions) (*runtime.Poller[DeviceSettingsClientCreateOrUpdateTimeSettingsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateTimeSettings(ctx, deviceName, resourceGroupName, managerName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DeviceSettingsClientCreateOrUpdateTimeSettingsResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[DeviceSettingsClientCreateOrUpdateTimeSettingsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdateTimeSettings - Creates or updates the time settings of the specified device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
func (client *DeviceSettingsClient) createOrUpdateTimeSettings(ctx context.Context, deviceName string, resourceGroupName string, managerName string, parameters TimeSettings, options *DeviceSettingsClientBeginCreateOrUpdateTimeSettingsOptions) (*http.Response, error) {
	req, err := client.createOrUpdateTimeSettingsCreateRequest(ctx, deviceName, resourceGroupName, managerName, parameters, options)
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

// createOrUpdateTimeSettingsCreateRequest creates the CreateOrUpdateTimeSettings request.
func (client *DeviceSettingsClient) createOrUpdateTimeSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, parameters TimeSettings, options *DeviceSettingsClientBeginCreateOrUpdateTimeSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/timeSettings/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// GetAlertSettings - Gets the alert settings of the specified device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
//   - deviceName - The device name
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - DeviceSettingsClientGetAlertSettingsOptions contains the optional parameters for the DeviceSettingsClient.GetAlertSettings
//     method.
func (client *DeviceSettingsClient) GetAlertSettings(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *DeviceSettingsClientGetAlertSettingsOptions) (DeviceSettingsClientGetAlertSettingsResponse, error) {
	req, err := client.getAlertSettingsCreateRequest(ctx, deviceName, resourceGroupName, managerName, options)
	if err != nil {
		return DeviceSettingsClientGetAlertSettingsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeviceSettingsClientGetAlertSettingsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeviceSettingsClientGetAlertSettingsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAlertSettingsHandleResponse(resp)
}

// getAlertSettingsCreateRequest creates the GetAlertSettings request.
func (client *DeviceSettingsClient) getAlertSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *DeviceSettingsClientGetAlertSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/alertSettings/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAlertSettingsHandleResponse handles the GetAlertSettings response.
func (client *DeviceSettingsClient) getAlertSettingsHandleResponse(resp *http.Response) (DeviceSettingsClientGetAlertSettingsResponse, error) {
	result := DeviceSettingsClientGetAlertSettingsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertSettings); err != nil {
		return DeviceSettingsClientGetAlertSettingsResponse{}, err
	}
	return result, nil
}

// GetNetworkSettings - Gets the network settings of the specified device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
//   - deviceName - The device name
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - DeviceSettingsClientGetNetworkSettingsOptions contains the optional parameters for the DeviceSettingsClient.GetNetworkSettings
//     method.
func (client *DeviceSettingsClient) GetNetworkSettings(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *DeviceSettingsClientGetNetworkSettingsOptions) (DeviceSettingsClientGetNetworkSettingsResponse, error) {
	req, err := client.getNetworkSettingsCreateRequest(ctx, deviceName, resourceGroupName, managerName, options)
	if err != nil {
		return DeviceSettingsClientGetNetworkSettingsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeviceSettingsClientGetNetworkSettingsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeviceSettingsClientGetNetworkSettingsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getNetworkSettingsHandleResponse(resp)
}

// getNetworkSettingsCreateRequest creates the GetNetworkSettings request.
func (client *DeviceSettingsClient) getNetworkSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *DeviceSettingsClientGetNetworkSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/networkSettings/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNetworkSettingsHandleResponse handles the GetNetworkSettings response.
func (client *DeviceSettingsClient) getNetworkSettingsHandleResponse(resp *http.Response) (DeviceSettingsClientGetNetworkSettingsResponse, error) {
	result := DeviceSettingsClientGetNetworkSettingsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkSettings); err != nil {
		return DeviceSettingsClientGetNetworkSettingsResponse{}, err
	}
	return result, nil
}

// GetSecuritySettings - Returns the Security properties of the specified device name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
//   - deviceName - The device name
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - DeviceSettingsClientGetSecuritySettingsOptions contains the optional parameters for the DeviceSettingsClient.GetSecuritySettings
//     method.
func (client *DeviceSettingsClient) GetSecuritySettings(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *DeviceSettingsClientGetSecuritySettingsOptions) (DeviceSettingsClientGetSecuritySettingsResponse, error) {
	req, err := client.getSecuritySettingsCreateRequest(ctx, deviceName, resourceGroupName, managerName, options)
	if err != nil {
		return DeviceSettingsClientGetSecuritySettingsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeviceSettingsClientGetSecuritySettingsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeviceSettingsClientGetSecuritySettingsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSecuritySettingsHandleResponse(resp)
}

// getSecuritySettingsCreateRequest creates the GetSecuritySettings request.
func (client *DeviceSettingsClient) getSecuritySettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *DeviceSettingsClientGetSecuritySettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/securitySettings/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSecuritySettingsHandleResponse handles the GetSecuritySettings response.
func (client *DeviceSettingsClient) getSecuritySettingsHandleResponse(resp *http.Response) (DeviceSettingsClientGetSecuritySettingsResponse, error) {
	result := DeviceSettingsClientGetSecuritySettingsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecuritySettings); err != nil {
		return DeviceSettingsClientGetSecuritySettingsResponse{}, err
	}
	return result, nil
}

// GetTimeSettings - Gets the time settings of the specified device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
//   - deviceName - The device name
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - DeviceSettingsClientGetTimeSettingsOptions contains the optional parameters for the DeviceSettingsClient.GetTimeSettings
//     method.
func (client *DeviceSettingsClient) GetTimeSettings(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *DeviceSettingsClientGetTimeSettingsOptions) (DeviceSettingsClientGetTimeSettingsResponse, error) {
	req, err := client.getTimeSettingsCreateRequest(ctx, deviceName, resourceGroupName, managerName, options)
	if err != nil {
		return DeviceSettingsClientGetTimeSettingsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeviceSettingsClientGetTimeSettingsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeviceSettingsClientGetTimeSettingsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getTimeSettingsHandleResponse(resp)
}

// getTimeSettingsCreateRequest creates the GetTimeSettings request.
func (client *DeviceSettingsClient) getTimeSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *DeviceSettingsClientGetTimeSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/timeSettings/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getTimeSettingsHandleResponse handles the GetTimeSettings response.
func (client *DeviceSettingsClient) getTimeSettingsHandleResponse(resp *http.Response) (DeviceSettingsClientGetTimeSettingsResponse, error) {
	result := DeviceSettingsClientGetTimeSettingsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TimeSettings); err != nil {
		return DeviceSettingsClientGetTimeSettingsResponse{}, err
	}
	return result, nil
}

// BeginSyncRemotemanagementCertificate - sync Remote management Certificate between appliance and Service
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
//   - deviceName - The device name
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - DeviceSettingsClientBeginSyncRemotemanagementCertificateOptions contains the optional parameters for the DeviceSettingsClient.BeginSyncRemotemanagementCertificate
//     method.
func (client *DeviceSettingsClient) BeginSyncRemotemanagementCertificate(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *DeviceSettingsClientBeginSyncRemotemanagementCertificateOptions) (*runtime.Poller[DeviceSettingsClientSyncRemotemanagementCertificateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.syncRemotemanagementCertificate(ctx, deviceName, resourceGroupName, managerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DeviceSettingsClientSyncRemotemanagementCertificateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[DeviceSettingsClientSyncRemotemanagementCertificateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// SyncRemotemanagementCertificate - sync Remote management Certificate between appliance and Service
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
func (client *DeviceSettingsClient) syncRemotemanagementCertificate(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *DeviceSettingsClientBeginSyncRemotemanagementCertificateOptions) (*http.Response, error) {
	req, err := client.syncRemotemanagementCertificateCreateRequest(ctx, deviceName, resourceGroupName, managerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// syncRemotemanagementCertificateCreateRequest creates the SyncRemotemanagementCertificate request.
func (client *DeviceSettingsClient) syncRemotemanagementCertificateCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *DeviceSettingsClientBeginSyncRemotemanagementCertificateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/securitySettings/default/syncRemoteManagementCertificate"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginUpdateNetworkSettings - Updates the network settings on the specified device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
//   - deviceName - The device name
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - parameters - The network settings to be updated.
//   - options - DeviceSettingsClientBeginUpdateNetworkSettingsOptions contains the optional parameters for the DeviceSettingsClient.BeginUpdateNetworkSettings
//     method.
func (client *DeviceSettingsClient) BeginUpdateNetworkSettings(ctx context.Context, deviceName string, resourceGroupName string, managerName string, parameters NetworkSettingsPatch, options *DeviceSettingsClientBeginUpdateNetworkSettingsOptions) (*runtime.Poller[DeviceSettingsClientUpdateNetworkSettingsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateNetworkSettings(ctx, deviceName, resourceGroupName, managerName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DeviceSettingsClientUpdateNetworkSettingsResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[DeviceSettingsClientUpdateNetworkSettingsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdateNetworkSettings - Updates the network settings on the specified device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
func (client *DeviceSettingsClient) updateNetworkSettings(ctx context.Context, deviceName string, resourceGroupName string, managerName string, parameters NetworkSettingsPatch, options *DeviceSettingsClientBeginUpdateNetworkSettingsOptions) (*http.Response, error) {
	req, err := client.updateNetworkSettingsCreateRequest(ctx, deviceName, resourceGroupName, managerName, parameters, options)
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

// updateNetworkSettingsCreateRequest creates the UpdateNetworkSettings request.
func (client *DeviceSettingsClient) updateNetworkSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, parameters NetworkSettingsPatch, options *DeviceSettingsClientBeginUpdateNetworkSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/networkSettings/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginUpdateSecuritySettings - Patch Security properties of the specified device name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
//   - deviceName - The device name
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - parameters - The security settings properties to be patched.
//   - options - DeviceSettingsClientBeginUpdateSecuritySettingsOptions contains the optional parameters for the DeviceSettingsClient.BeginUpdateSecuritySettings
//     method.
func (client *DeviceSettingsClient) BeginUpdateSecuritySettings(ctx context.Context, deviceName string, resourceGroupName string, managerName string, parameters SecuritySettingsPatch, options *DeviceSettingsClientBeginUpdateSecuritySettingsOptions) (*runtime.Poller[DeviceSettingsClientUpdateSecuritySettingsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateSecuritySettings(ctx, deviceName, resourceGroupName, managerName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DeviceSettingsClientUpdateSecuritySettingsResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[DeviceSettingsClientUpdateSecuritySettingsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdateSecuritySettings - Patch Security properties of the specified device name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
func (client *DeviceSettingsClient) updateSecuritySettings(ctx context.Context, deviceName string, resourceGroupName string, managerName string, parameters SecuritySettingsPatch, options *DeviceSettingsClientBeginUpdateSecuritySettingsOptions) (*http.Response, error) {
	req, err := client.updateSecuritySettingsCreateRequest(ctx, deviceName, resourceGroupName, managerName, parameters, options)
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

// updateSecuritySettingsCreateRequest creates the UpdateSecuritySettings request.
func (client *DeviceSettingsClient) updateSecuritySettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, parameters SecuritySettingsPatch, options *DeviceSettingsClientBeginUpdateSecuritySettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/securitySettings/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", deviceName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
