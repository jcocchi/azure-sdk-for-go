//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/Policy.json
func ExamplePoliciesClient_GetByBillingProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewPoliciesClient(cred, nil)
	res, err := client.GetByBillingProfile(ctx,
		"<billing-account-name>",
		"<billing-profile-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Policy.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdatePolicy.json
func ExamplePoliciesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewPoliciesClient(cred, nil)
	res, err := client.Update(ctx,
		"<billing-account-name>",
		"<billing-profile-name>",
		armbilling.Policy{
			Properties: &armbilling.PolicyProperties{
				MarketplacePurchases: armbilling.MarketplacePurchasesPolicyOnlyFreeAllowed.ToPtr(),
				ReservationPurchases: armbilling.ReservationPurchasesPolicyNotAllowed.ToPtr(),
				ViewCharges:          armbilling.ViewChargesPolicyAllowed.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Policy.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/CustomerPolicy.json
func ExamplePoliciesClient_GetByCustomer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewPoliciesClient(cred, nil)
	res, err := client.GetByCustomer(ctx,
		"<billing-account-name>",
		"<customer-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CustomerPolicy.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdateCustomerPolicy.json
func ExamplePoliciesClient_UpdateCustomer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewPoliciesClient(cred, nil)
	res, err := client.UpdateCustomer(ctx,
		"<billing-account-name>",
		"<customer-name>",
		armbilling.CustomerPolicy{
			Properties: &armbilling.CustomerPolicyProperties{
				ViewCharges: armbilling.ViewChargesNotAllowed.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CustomerPolicy.ID: %s\n", *res.ID)
}
