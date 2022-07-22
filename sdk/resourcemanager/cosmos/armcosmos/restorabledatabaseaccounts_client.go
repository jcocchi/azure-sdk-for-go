//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcosmos

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RestorableDatabaseAccountsClient contains the methods for the RestorableDatabaseAccounts group.
// Don't use this type directly, use NewRestorableDatabaseAccountsClient() instead.
type RestorableDatabaseAccountsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRestorableDatabaseAccountsClient creates a new instance of RestorableDatabaseAccountsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRestorableDatabaseAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RestorableDatabaseAccountsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &RestorableDatabaseAccountsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// GetByLocation - Retrieves the properties of an existing Azure Cosmos DB restorable database account. This call requires
// 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/read/*' permission.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-15
// location - Cosmos DB region, with spaces between words and each word capitalized.
// instanceID - The instanceId GUID of a restorable database account.
// options - RestorableDatabaseAccountsClientGetByLocationOptions contains the optional parameters for the RestorableDatabaseAccountsClient.GetByLocation
// method.
func (client *RestorableDatabaseAccountsClient) GetByLocation(ctx context.Context, location string, instanceID string, options *RestorableDatabaseAccountsClientGetByLocationOptions) (RestorableDatabaseAccountsClientGetByLocationResponse, error) {
	req, err := client.getByLocationCreateRequest(ctx, location, instanceID, options)
	if err != nil {
		return RestorableDatabaseAccountsClientGetByLocationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RestorableDatabaseAccountsClientGetByLocationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RestorableDatabaseAccountsClientGetByLocationResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByLocationHandleResponse(resp)
}

// getByLocationCreateRequest creates the GetByLocation request.
func (client *RestorableDatabaseAccountsClient) getByLocationCreateRequest(ctx context.Context, location string, instanceID string, options *RestorableDatabaseAccountsClientGetByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByLocationHandleResponse handles the GetByLocation response.
func (client *RestorableDatabaseAccountsClient) getByLocationHandleResponse(resp *http.Response) (RestorableDatabaseAccountsClientGetByLocationResponse, error) {
	result := RestorableDatabaseAccountsClientGetByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableDatabaseAccountGetResult); err != nil {
		return RestorableDatabaseAccountsClientGetByLocationResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the restorable Azure Cosmos DB database accounts available under the subscription. This call requires
// 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/read' permission.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-15
// options - RestorableDatabaseAccountsClientListOptions contains the optional parameters for the RestorableDatabaseAccountsClient.List
// method.
func (client *RestorableDatabaseAccountsClient) NewListPager(options *RestorableDatabaseAccountsClientListOptions) *runtime.Pager[RestorableDatabaseAccountsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RestorableDatabaseAccountsClientListResponse]{
		More: func(page RestorableDatabaseAccountsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *RestorableDatabaseAccountsClientListResponse) (RestorableDatabaseAccountsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, options)
			if err != nil {
				return RestorableDatabaseAccountsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RestorableDatabaseAccountsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RestorableDatabaseAccountsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RestorableDatabaseAccountsClient) listCreateRequest(ctx context.Context, options *RestorableDatabaseAccountsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/restorableDatabaseAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorableDatabaseAccountsClient) listHandleResponse(resp *http.Response) (RestorableDatabaseAccountsClientListResponse, error) {
	result := RestorableDatabaseAccountsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableDatabaseAccountsListResult); err != nil {
		return RestorableDatabaseAccountsClientListResponse{}, err
	}
	return result, nil
}

// NewListByLocationPager - Lists all the restorable Azure Cosmos DB database accounts available under the subscription and
// in a region. This call requires 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/read'
// permission.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-15
// location - Cosmos DB region, with spaces between words and each word capitalized.
// options - RestorableDatabaseAccountsClientListByLocationOptions contains the optional parameters for the RestorableDatabaseAccountsClient.ListByLocation
// method.
func (client *RestorableDatabaseAccountsClient) NewListByLocationPager(location string, options *RestorableDatabaseAccountsClientListByLocationOptions) *runtime.Pager[RestorableDatabaseAccountsClientListByLocationResponse] {
	return runtime.NewPager(runtime.PagingHandler[RestorableDatabaseAccountsClientListByLocationResponse]{
		More: func(page RestorableDatabaseAccountsClientListByLocationResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *RestorableDatabaseAccountsClientListByLocationResponse) (RestorableDatabaseAccountsClientListByLocationResponse, error) {
			req, err := client.listByLocationCreateRequest(ctx, location, options)
			if err != nil {
				return RestorableDatabaseAccountsClientListByLocationResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return RestorableDatabaseAccountsClientListByLocationResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RestorableDatabaseAccountsClientListByLocationResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByLocationHandleResponse(resp)
		},
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *RestorableDatabaseAccountsClient) listByLocationCreateRequest(ctx context.Context, location string, options *RestorableDatabaseAccountsClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *RestorableDatabaseAccountsClient) listByLocationHandleResponse(resp *http.Response) (RestorableDatabaseAccountsClientListByLocationResponse, error) {
	result := RestorableDatabaseAccountsClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableDatabaseAccountsListResult); err != nil {
		return RestorableDatabaseAccountsClientListByLocationResponse{}, err
	}
	return result, nil
}