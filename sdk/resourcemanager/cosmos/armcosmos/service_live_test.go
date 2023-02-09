//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/testutil"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite

	ctx               context.Context
	cred              azcore.TokenCredential
	options           *arm.ClientOptions
	accountName       string
	location          string
	resourceGroupName string
	subscriptionId    string
}

func (testsuite *ServiceTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.accountName = "databaseaccountservice2"
	testsuite.location = testutil.GetEnv("LOCATION", "westus")
	testsuite.resourceGroupName = testutil.GetEnv("RESOURCE_GROUP_NAME", "scenarioTestTempGroup")
	testsuite.subscriptionId = testutil.GetEnv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")

	testutil.StartRecording(testsuite.T(), "sdk/resourcemanager/cosmos/armcosmos/testdata")
	resourceGroup, _, err := testutil.CreateResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.location)
	testsuite.Require().NoError(err)
	testsuite.resourceGroupName = *resourceGroup.Name
	testsuite.Prepare()
}

func (testsuite *ServiceTestSuite) TearDownSuite() {
	_, err := testutil.DeleteResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName)
	testsuite.Require().NoError(err)
	testutil.StopRecording(testsuite.T())
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (testsuite *ServiceTestSuite) Prepare() {
	var err error
	// From step DatabaseAccount_Create
	fmt.Println("Call operation: DatabaseAccounts_CreateOrUpdate")
	databaseAccountsClient, err := armcosmos.NewDatabaseAccountsClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	databaseAccountsClientCreateOrUpdateResponsePoller, err := databaseAccountsClient.BeginCreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, armcosmos.DatabaseAccountCreateUpdateParameters{
		Location: to.Ptr(testsuite.location),
		Properties: &armcosmos.DatabaseAccountCreateUpdateProperties{
			CreateMode:               to.Ptr(armcosmos.CreateModeDefault),
			DatabaseAccountOfferType: to.Ptr("Standard"),
			Locations: []*armcosmos.Location{
				{
					FailoverPriority: to.Ptr[int32](0),
					IsZoneRedundant:  to.Ptr(false),
					LocationName:     to.Ptr("eastus"),
				}},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, databaseAccountsClientCreateOrUpdateResponsePoller)
	testsuite.Require().NoError(err)
}

// Microsoft.DocumentDB/databaseAccounts/services
func (testsuite *ServiceTestSuite) TestService() {
	var err error
	// From step Service_CreateSqlDedicatedGatewayService
	fmt.Println("Call operation: Service_Create")
	serviceClient, err := armcosmos.NewServiceClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	serviceClientCreateResponsePoller, err := serviceClient.BeginCreate(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, "SqlDedicatedGateway", armcosmos.ServiceResourceCreateUpdateParameters{
		Properties: &armcosmos.ServiceResourceCreateUpdateProperties{
			InstanceCount: to.Ptr[int32](1),
			InstanceSize:  to.Ptr(armcosmos.ServiceSizeCosmosD4S),
			ServiceType:   to.Ptr(armcosmos.ServiceTypeSQLDedicatedGateway),
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, serviceClientCreateResponsePoller)
	testsuite.Require().NoError(err)

	// From step Service_GetSqlDedicatedGatewayService
	fmt.Println("Call operation: Service_Get")
	_, err = serviceClient.Get(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, "SqlDedicatedGateway", nil)
	testsuite.Require().NoError(err)

	// From step Service_List
	fmt.Println("Call operation: Service_List")
	serviceClientNewListPager := serviceClient.NewListPager(testsuite.resourceGroupName, testsuite.accountName, nil)
	for serviceClientNewListPager.More() {
		_, err := serviceClientNewListPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step Service_DeleteSqlDedicatedGatewayService
	fmt.Println("Call operation: Service_Delete")
	serviceClientDeleteResponsePoller, err := serviceClient.BeginDelete(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, "SqlDedicatedGateway", nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, serviceClientDeleteResponsePoller)
	testsuite.Require().NoError(err)
}
