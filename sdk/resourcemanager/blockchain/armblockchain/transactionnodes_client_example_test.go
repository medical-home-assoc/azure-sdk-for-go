//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armblockchain_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/TransactionNodes_Get.json
func ExampleTransactionNodesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblockchain.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTransactionNodesClient().Get(ctx, "contosemember1", "txnode2", "mygroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TransactionNode = armblockchain.TransactionNode{
	// 	Name: to.Ptr("txnode2"),
	// 	Type: to.Ptr("Microsoft.Blockchain/blockchainMembers/transactionNodes"),
	// 	ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourceGroups/mygroup/providers/Microsoft.Blockchain/blockchainMembers/contosemember1/transactionNodes/txnode2"),
	// 	Location: to.Ptr("southeastasia"),
	// 	Properties: &armblockchain.TransactionNodeProperties{
	// 		DNS: to.Ptr("txnode2-contosemember1.blockchain.ppe.azure-int.net"),
	// 		ProvisioningState: to.Ptr(armblockchain.NodeProvisioningStateSucceeded),
	// 		PublicKey: to.Ptr("h7Q10I/1dLK/hzX8FkVrfl03D/aX8jW3YNoxJ/n4vkY="),
	// 		UserName: to.Ptr("txnode2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/TransactionNodes_Create.json
func ExampleTransactionNodesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblockchain.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTransactionNodesClient().BeginCreate(ctx, "contosemember1", "txnode2", "mygroup", &armblockchain.TransactionNodesClientBeginCreateOptions{TransactionNode: &armblockchain.TransactionNode{
		Location: to.Ptr("southeastasia"),
		Properties: &armblockchain.TransactionNodeProperties{
			Password: to.Ptr("<password>"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TransactionNode = armblockchain.TransactionNode{
	// 	Name: to.Ptr("txnode2"),
	// 	Type: to.Ptr("Microsoft.Blockchain/blockchainMembers/transactionNodes"),
	// 	ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourceGroups/mygroup/providers/Microsoft.Blockchain/blockchainMembers/contosemember1/transactionNodes/txnode2"),
	// 	Location: to.Ptr("southeastasia"),
	// 	Properties: &armblockchain.TransactionNodeProperties{
	// 		ProvisioningState: to.Ptr(armblockchain.NodeProvisioningStateSucceeded),
	// 		UserName: to.Ptr("txnode2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/TransactionNodes_Delete.json
func ExampleTransactionNodesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblockchain.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTransactionNodesClient().BeginDelete(ctx, "contosemember1", "txNode2", "mygroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/TransactionNodes_Update.json
func ExampleTransactionNodesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblockchain.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTransactionNodesClient().Update(ctx, "contosemember1", "txnode2", "mygroup", &armblockchain.TransactionNodesClientUpdateOptions{TransactionNode: &armblockchain.TransactionNodeUpdate{
		Properties: &armblockchain.TransactionNodePropertiesUpdate{
			Password: to.Ptr("<password>"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TransactionNode = armblockchain.TransactionNode{
	// 	Name: to.Ptr("txnode2"),
	// 	Type: to.Ptr("Microsoft.Blockchain/blockchainMembers/transactionNodes"),
	// 	ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourceGroups/mygroup/providers/Microsoft.Blockchain/blockchainMembers/contosemember1/transactionNodes/txnode2"),
	// 	Location: to.Ptr("southeastasia"),
	// 	Properties: &armblockchain.TransactionNodeProperties{
	// 		DNS: to.Ptr("txnode1-contosemember1.blockchain.azure.com"),
	// 		ProvisioningState: to.Ptr(armblockchain.NodeProvisioningStateSucceeded),
	// 		PublicKey: to.Ptr("DbRYTorBtY7rZfNfByUQpdC+hD3k/0lfA7+UnH4ovWM="),
	// 		UserName: to.Ptr("txnode2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/TransactionNodes_List.json
func ExampleTransactionNodesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblockchain.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTransactionNodesClient().NewListPager("contosemember1", "mygroup", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.TransactionNodeCollection = armblockchain.TransactionNodeCollection{
		// 	Value: []*armblockchain.TransactionNode{
		// 		{
		// 			Name: to.Ptr("txnode2"),
		// 			Type: to.Ptr("Microsoft.Blockchain/blockchainMembers/transactionNodes"),
		// 			ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourceGroups/mygroup/providers/Microsoft.Blockchain/blockchainMembers/contosemember1/transactionNodes/txnode2"),
		// 			Location: to.Ptr("southeastasia"),
		// 			Properties: &armblockchain.TransactionNodeProperties{
		// 				DNS: to.Ptr("txnode2-contosemember1.blockchain.azure.com"),
		// 				ProvisioningState: to.Ptr(armblockchain.NodeProvisioningStateSucceeded),
		// 				PublicKey: to.Ptr("DbRYTorBtY7rZfNfByUQpdC+hD3k/0lfA7+UnH4ovWM="),
		// 				UserName: to.Ptr("txnode2"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/TransactionNodes_ListApiKeys.json
func ExampleTransactionNodesClient_ListAPIKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblockchain.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTransactionNodesClient().ListAPIKeys(ctx, "contosemember1", "txnode2", "mygroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.APIKeyCollection = armblockchain.APIKeyCollection{
	// 	Keys: []*armblockchain.APIKey{
	// 		{
	// 			KeyName: to.Ptr("key1"),
	// 			Value: to.Ptr("-EnIUzu29xj60xPJmstyCURo"),
	// 		},
	// 		{
	// 			KeyName: to.Ptr("key2"),
	// 			Value: to.Ptr("I8P-q4u_WDuCZUBYXnJ3yYX7"),
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/TransactionNodes_ListRegenerateApiKeys.json
func ExampleTransactionNodesClient_ListRegenerateAPIKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armblockchain.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTransactionNodesClient().ListRegenerateAPIKeys(ctx, "contosemember1", "txnode2", "mygroup", &armblockchain.TransactionNodesClientListRegenerateAPIKeysOptions{APIKey: &armblockchain.APIKey{
		KeyName: to.Ptr("key1"),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.APIKeyCollection = armblockchain.APIKeyCollection{
	// 	Keys: []*armblockchain.APIKey{
	// 		{
	// 			KeyName: to.Ptr("key1"),
	// 			Value: to.Ptr("-EnIUzu29xj60xPJmstyCURo"),
	// 		},
	// 		{
	// 			KeyName: to.Ptr("key2"),
	// 			Value: to.Ptr("-EHio4yRJLxajDaxRNaoD7cZ"),
	// 	}},
	// }
}
