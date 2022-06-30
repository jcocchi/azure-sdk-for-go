//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armworkloads

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

// SAPVirtualInstancesClient contains the methods for the SAPVirtualInstances group.
// Don't use this type directly, use NewSAPVirtualInstancesClient() instead.
type SAPVirtualInstancesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSAPVirtualInstancesClient creates a new instance of SAPVirtualInstancesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSAPVirtualInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SAPVirtualInstancesClient, error) {
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
	client := &SAPVirtualInstancesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Creates an Virtual Instance for SAP.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// sapVirtualInstanceName - The name of the Virtual Instances for SAP.
// body - The Virtual Instance for SAP request body.
// options - SAPVirtualInstancesClientBeginCreateOptions contains the optional parameters for the SAPVirtualInstancesClient.BeginCreate
// method.
func (client *SAPVirtualInstancesClient) BeginCreate(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, body SAPVirtualInstance, options *SAPVirtualInstancesClientBeginCreateOptions) (*runtime.Poller[SAPVirtualInstancesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, sapVirtualInstanceName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SAPVirtualInstancesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SAPVirtualInstancesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates an Virtual Instance for SAP.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
func (client *SAPVirtualInstancesClient) create(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, body SAPVirtualInstance, options *SAPVirtualInstancesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *SAPVirtualInstancesClient) createCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, body SAPVirtualInstance, options *SAPVirtualInstancesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Deletes an Virtual Instance for SAP.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// sapVirtualInstanceName - The name of the Virtual Instances for SAP.
// options - SAPVirtualInstancesClientBeginDeleteOptions contains the optional parameters for the SAPVirtualInstancesClient.BeginDelete
// method.
func (client *SAPVirtualInstancesClient) BeginDelete(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, options *SAPVirtualInstancesClientBeginDeleteOptions) (*runtime.Poller[SAPVirtualInstancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, sapVirtualInstanceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SAPVirtualInstancesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SAPVirtualInstancesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes an Virtual Instance for SAP.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
func (client *SAPVirtualInstancesClient) deleteOperation(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, options *SAPVirtualInstancesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SAPVirtualInstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, options *SAPVirtualInstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an Virtual Instance for SAP.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// sapVirtualInstanceName - The name of the Virtual Instances for SAP.
// options - SAPVirtualInstancesClientGetOptions contains the optional parameters for the SAPVirtualInstancesClient.Get method.
func (client *SAPVirtualInstancesClient) Get(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, options *SAPVirtualInstancesClientGetOptions) (SAPVirtualInstancesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, options)
	if err != nil {
		return SAPVirtualInstancesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SAPVirtualInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SAPVirtualInstancesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SAPVirtualInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, options *SAPVirtualInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SAPVirtualInstancesClient) getHandleResponse(resp *http.Response) (SAPVirtualInstancesClientGetResponse, error) {
	result := SAPVirtualInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPVirtualInstance); err != nil {
		return SAPVirtualInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all Virtual Instances for SAP in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - SAPVirtualInstancesClientListByResourceGroupOptions contains the optional parameters for the SAPVirtualInstancesClient.ListByResourceGroup
// method.
func (client *SAPVirtualInstancesClient) NewListByResourceGroupPager(resourceGroupName string, options *SAPVirtualInstancesClientListByResourceGroupOptions) *runtime.Pager[SAPVirtualInstancesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SAPVirtualInstancesClientListByResourceGroupResponse]{
		More: func(page SAPVirtualInstancesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SAPVirtualInstancesClientListByResourceGroupResponse) (SAPVirtualInstancesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SAPVirtualInstancesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SAPVirtualInstancesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SAPVirtualInstancesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SAPVirtualInstancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SAPVirtualInstancesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SAPVirtualInstancesClient) listByResourceGroupHandleResponse(resp *http.Response) (SAPVirtualInstancesClientListByResourceGroupResponse, error) {
	result := SAPVirtualInstancesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPVirtualInstanceList); err != nil {
		return SAPVirtualInstancesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets all Virtual Instances for SAP in the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// options - SAPVirtualInstancesClientListBySubscriptionOptions contains the optional parameters for the SAPVirtualInstancesClient.ListBySubscription
// method.
func (client *SAPVirtualInstancesClient) NewListBySubscriptionPager(options *SAPVirtualInstancesClientListBySubscriptionOptions) *runtime.Pager[SAPVirtualInstancesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[SAPVirtualInstancesClientListBySubscriptionResponse]{
		More: func(page SAPVirtualInstancesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SAPVirtualInstancesClientListBySubscriptionResponse) (SAPVirtualInstancesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SAPVirtualInstancesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SAPVirtualInstancesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SAPVirtualInstancesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SAPVirtualInstancesClient) listBySubscriptionCreateRequest(ctx context.Context, options *SAPVirtualInstancesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Workloads/sapVirtualInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SAPVirtualInstancesClient) listBySubscriptionHandleResponse(resp *http.Response) (SAPVirtualInstancesClientListBySubscriptionResponse, error) {
	result := SAPVirtualInstancesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPVirtualInstanceList); err != nil {
		return SAPVirtualInstancesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginStart - Starts the SAP System.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// sapVirtualInstanceName - The name of the Virtual Instances for SAP.
// options - SAPVirtualInstancesClientBeginStartOptions contains the optional parameters for the SAPVirtualInstancesClient.BeginStart
// method.
func (client *SAPVirtualInstancesClient) BeginStart(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, options *SAPVirtualInstancesClientBeginStartOptions) (*runtime.Poller[SAPVirtualInstancesClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, sapVirtualInstanceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SAPVirtualInstancesClientStartResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SAPVirtualInstancesClientStartResponse](options.ResumeToken, client.pl, nil)
	}
}

// Start - Starts the SAP System.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
func (client *SAPVirtualInstancesClient) start(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, options *SAPVirtualInstancesClientBeginStartOptions) (*http.Response, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// startCreateRequest creates the Start request.
func (client *SAPVirtualInstancesClient) startCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, options *SAPVirtualInstancesClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStop - Stops the SAP System.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// sapVirtualInstanceName - The name of the Virtual Instances for SAP.
// options - SAPVirtualInstancesClientBeginStopOptions contains the optional parameters for the SAPVirtualInstancesClient.BeginStop
// method.
func (client *SAPVirtualInstancesClient) BeginStop(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, options *SAPVirtualInstancesClientBeginStopOptions) (*runtime.Poller[SAPVirtualInstancesClientStopResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stop(ctx, resourceGroupName, sapVirtualInstanceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SAPVirtualInstancesClientStopResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SAPVirtualInstancesClientStopResponse](options.ResumeToken, client.pl, nil)
	}
}

// Stop - Stops the SAP System.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
func (client *SAPVirtualInstancesClient) stop(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, options *SAPVirtualInstancesClientBeginStopOptions) (*http.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// stopCreateRequest creates the Stop request.
func (client *SAPVirtualInstancesClient) stopCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, options *SAPVirtualInstancesClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// Update - Updates an Virtual Instance for SAP.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// sapVirtualInstanceName - The name of the Virtual Instances for SAP.
// body - The Update Virtual Instance for SAP request body.
// options - SAPVirtualInstancesClientUpdateOptions contains the optional parameters for the SAPVirtualInstancesClient.Update
// method.
func (client *SAPVirtualInstancesClient) Update(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, body UpdateSAPVirtualInstanceRequest, options *SAPVirtualInstancesClientUpdateOptions) (SAPVirtualInstancesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, body, options)
	if err != nil {
		return SAPVirtualInstancesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SAPVirtualInstancesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SAPVirtualInstancesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SAPVirtualInstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, body UpdateSAPVirtualInstanceRequest, options *SAPVirtualInstancesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleResponse handles the Update response.
func (client *SAPVirtualInstancesClient) updateHandleResponse(resp *http.Response) (SAPVirtualInstancesClientUpdateResponse, error) {
	result := SAPVirtualInstancesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPVirtualInstance); err != nil {
		return SAPVirtualInstancesClientUpdateResponse{}, err
	}
	return result, nil
}