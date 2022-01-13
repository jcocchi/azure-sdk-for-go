//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakestore_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
)

// x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_List.json
func ExampleAccountsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatalakestore.NewAccountsClient("<subscription-id>", cred, nil)
	pager := client.List(&armdatalakestore.AccountsClientListOptions{Filter: to.StringPtr("<filter>"),
		Top:     to.Int32Ptr(1),
		Skip:    to.Int32Ptr(1),
		Select:  to.StringPtr("<select>"),
		Orderby: to.StringPtr("<orderby>"),
		Count:   to.BoolPtr(false),
	})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_ListByResourceGroup.json
func ExampleAccountsClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatalakestore.NewAccountsClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		&armdatalakestore.AccountsClientListByResourceGroupOptions{Filter: to.StringPtr("<filter>"),
			Top:     to.Int32Ptr(1),
			Skip:    to.Int32Ptr(1),
			Select:  to.StringPtr("<select>"),
			Orderby: to.StringPtr("<orderby>"),
			Count:   to.BoolPtr(false),
		})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_Create.json
func ExampleAccountsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatalakestore.NewAccountsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armdatalakestore.CreateDataLakeStoreAccountParameters{
			Identity: &armdatalakestore.EncryptionIdentity{
				Type: to.StringPtr("<type>"),
			},
			Location: to.StringPtr("<location>"),
			Properties: &armdatalakestore.CreateDataLakeStoreAccountProperties{
				DefaultGroup: to.StringPtr("<default-group>"),
				EncryptionConfig: &armdatalakestore.EncryptionConfig{
					Type: armdatalakestore.EncryptionConfigTypeUserManaged.ToPtr(),
					KeyVaultMetaInfo: &armdatalakestore.KeyVaultMetaInfo{
						EncryptionKeyName:    to.StringPtr("<encryption-key-name>"),
						EncryptionKeyVersion: to.StringPtr("<encryption-key-version>"),
						KeyVaultResourceID:   to.StringPtr("<key-vault-resource-id>"),
					},
				},
				EncryptionState:       armdatalakestore.EncryptionStateEnabled.ToPtr(),
				FirewallAllowAzureIPs: armdatalakestore.FirewallAllowAzureIPsStateEnabled.ToPtr(),
				FirewallRules: []*armdatalakestore.CreateFirewallRuleWithAccountParameters{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armdatalakestore.CreateOrUpdateFirewallRuleProperties{
							EndIPAddress:   to.StringPtr("<end-ipaddress>"),
							StartIPAddress: to.StringPtr("<start-ipaddress>"),
						},
					}},
				FirewallState:          armdatalakestore.FirewallStateEnabled.ToPtr(),
				NewTier:                armdatalakestore.TierTypeConsumption.ToPtr(),
				TrustedIDProviderState: armdatalakestore.TrustedIDProviderStateEnabled.ToPtr(),
				TrustedIDProviders: []*armdatalakestore.CreateTrustedIDProviderWithAccountParameters{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armdatalakestore.CreateOrUpdateTrustedIDProviderProperties{
							IDProvider: to.StringPtr("<idprovider>"),
						},
					}},
			},
			Tags: map[string]*string{
				"test_key": to.StringPtr("test_value"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountsClientCreateResult)
}

// x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_Get.json
func ExampleAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatalakestore.NewAccountsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountsClientGetResult)
}

// x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_Update.json
func ExampleAccountsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatalakestore.NewAccountsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armdatalakestore.UpdateDataLakeStoreAccountParameters{
			Properties: &armdatalakestore.UpdateDataLakeStoreAccountProperties{
				DefaultGroup: to.StringPtr("<default-group>"),
				EncryptionConfig: &armdatalakestore.UpdateEncryptionConfig{
					KeyVaultMetaInfo: &armdatalakestore.UpdateKeyVaultMetaInfo{
						EncryptionKeyVersion: to.StringPtr("<encryption-key-version>"),
					},
				},
				FirewallAllowAzureIPs:  armdatalakestore.FirewallAllowAzureIPsStateEnabled.ToPtr(),
				FirewallState:          armdatalakestore.FirewallStateEnabled.ToPtr(),
				NewTier:                armdatalakestore.TierTypeConsumption.ToPtr(),
				TrustedIDProviderState: armdatalakestore.TrustedIDProviderStateEnabled.ToPtr(),
			},
			Tags: map[string]*string{
				"test_key": to.StringPtr("test_value"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountsClientUpdateResult)
}

// x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_Delete.json
func ExampleAccountsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatalakestore.NewAccountsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<account-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_EnableKeyVault.json
func ExampleAccountsClient_EnableKeyVault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatalakestore.NewAccountsClient("<subscription-id>", cred, nil)
	_, err = client.EnableKeyVault(ctx,
		"<resource-group-name>",
		"<account-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/Accounts_CheckNameAvailability.json
func ExampleAccountsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatalakestore.NewAccountsClient("<subscription-id>", cred, nil)
	res, err := client.CheckNameAvailability(ctx,
		"<location>",
		armdatalakestore.CheckNameAvailabilityParameters{
			Name: to.StringPtr("<name>"),
			Type: armdatalakestore.CheckNameAvailabilityParametersType("Microsoft.DataLakeStore/accounts").ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountsClientCheckNameAvailabilityResult)
}