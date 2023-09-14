//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armautomanage

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewBestPracticesClient() *BestPracticesClient {
	subClient, _ := NewBestPracticesClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBestPracticesVersionsClient() *BestPracticesVersionsClient {
	subClient, _ := NewBestPracticesVersionsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConfigurationProfilesClient() *ConfigurationProfilesClient {
	subClient, _ := NewConfigurationProfilesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConfigurationProfilesVersionsClient() *ConfigurationProfilesVersionsClient {
	subClient, _ := NewConfigurationProfilesVersionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConfigurationProfileAssignmentsClient() *ConfigurationProfileAssignmentsClient {
	subClient, _ := NewConfigurationProfileAssignmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReportsClient() *ReportsClient {
	subClient, _ := NewReportsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewServicePrincipalsClient() *ServicePrincipalsClient {
	subClient, _ := NewServicePrincipalsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConfigurationProfileHCRPAssignmentsClient() *ConfigurationProfileHCRPAssignmentsClient {
	subClient, _ := NewConfigurationProfileHCRPAssignmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewHCRPReportsClient() *HCRPReportsClient {
	subClient, _ := NewHCRPReportsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConfigurationProfileHCIAssignmentsClient() *ConfigurationProfileHCIAssignmentsClient {
	subClient, _ := NewConfigurationProfileHCIAssignmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewHCIReportsClient() *HCIReportsClient {
	subClient, _ := NewHCIReportsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
