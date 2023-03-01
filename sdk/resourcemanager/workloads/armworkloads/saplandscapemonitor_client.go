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

// SapLandscapeMonitorClient contains the methods for the SapLandscapeMonitor group.
// Don't use this type directly, use NewSapLandscapeMonitorClient() instead.
type SapLandscapeMonitorClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSapLandscapeMonitorClient creates a new instance of SapLandscapeMonitorClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSapLandscapeMonitorClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SapLandscapeMonitorClient, error) {
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
	client := &SapLandscapeMonitorClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Creates a SAP Landscape Monitor Dashboard for the specified subscription, resource group, and resource name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Name of the SAP monitor resource.
//   - sapLandscapeMonitorParameter - Request body representing a configuration for Sap Landscape Monitor Dashboard
//   - options - SapLandscapeMonitorClientCreateOptions contains the optional parameters for the SapLandscapeMonitorClient.Create
//     method.
func (client *SapLandscapeMonitorClient) Create(ctx context.Context, resourceGroupName string, monitorName string, sapLandscapeMonitorParameter SapLandscapeMonitor, options *SapLandscapeMonitorClientCreateOptions) (SapLandscapeMonitorClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, monitorName, sapLandscapeMonitorParameter, options)
	if err != nil {
		return SapLandscapeMonitorClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SapLandscapeMonitorClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SapLandscapeMonitorClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *SapLandscapeMonitorClient) createCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, sapLandscapeMonitorParameter SapLandscapeMonitor, options *SapLandscapeMonitorClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/monitors/{monitorName}/sapLandscapeMonitor/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, sapLandscapeMonitorParameter)
}

// createHandleResponse handles the Create response.
func (client *SapLandscapeMonitorClient) createHandleResponse(resp *http.Response) (SapLandscapeMonitorClientCreateResponse, error) {
	result := SapLandscapeMonitorClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SapLandscapeMonitor); err != nil {
		return SapLandscapeMonitorClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a SAP Landscape Monitor Dashboard with the specified subscription, resource group, and SAP monitor name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Name of the SAP monitor resource.
//   - options - SapLandscapeMonitorClientDeleteOptions contains the optional parameters for the SapLandscapeMonitorClient.Delete
//     method.
func (client *SapLandscapeMonitorClient) Delete(ctx context.Context, resourceGroupName string, monitorName string, options *SapLandscapeMonitorClientDeleteOptions) (SapLandscapeMonitorClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return SapLandscapeMonitorClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SapLandscapeMonitorClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SapLandscapeMonitorClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return SapLandscapeMonitorClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SapLandscapeMonitorClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *SapLandscapeMonitorClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/monitors/{monitorName}/sapLandscapeMonitor/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets configuration values for Single Pane Of Glass for SAP monitor for the specified subscription, resource group,
// and resource name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Name of the SAP monitor resource.
//   - options - SapLandscapeMonitorClientGetOptions contains the optional parameters for the SapLandscapeMonitorClient.Get method.
func (client *SapLandscapeMonitorClient) Get(ctx context.Context, resourceGroupName string, monitorName string, options *SapLandscapeMonitorClientGetOptions) (SapLandscapeMonitorClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return SapLandscapeMonitorClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SapLandscapeMonitorClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SapLandscapeMonitorClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SapLandscapeMonitorClient) getCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *SapLandscapeMonitorClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/monitors/{monitorName}/sapLandscapeMonitor/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SapLandscapeMonitorClient) getHandleResponse(resp *http.Response) (SapLandscapeMonitorClientGetResponse, error) {
	result := SapLandscapeMonitorClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SapLandscapeMonitor); err != nil {
		return SapLandscapeMonitorClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets configuration values for Single Pane Of Glass for SAP monitor for the specified subscription, resource group,
// and resource name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Name of the SAP monitor resource.
//   - options - SapLandscapeMonitorClientListOptions contains the optional parameters for the SapLandscapeMonitorClient.List
//     method.
func (client *SapLandscapeMonitorClient) List(ctx context.Context, resourceGroupName string, monitorName string, options *SapLandscapeMonitorClientListOptions) (SapLandscapeMonitorClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return SapLandscapeMonitorClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SapLandscapeMonitorClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SapLandscapeMonitorClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *SapLandscapeMonitorClient) listCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *SapLandscapeMonitorClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/monitors/{monitorName}/sapLandscapeMonitor"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SapLandscapeMonitorClient) listHandleResponse(resp *http.Response) (SapLandscapeMonitorClientListResponse, error) {
	result := SapLandscapeMonitorClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SapLandscapeMonitorListResult); err != nil {
		return SapLandscapeMonitorClientListResponse{}, err
	}
	return result, nil
}

// Update - Patches the SAP Landscape Monitor Dashboard for the specified subscription, resource group, and SAP monitor name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Name of the SAP monitor resource.
//   - sapLandscapeMonitorParameter - Request body representing a configuration for Sap Landscape Monitor Dashboard
//   - options - SapLandscapeMonitorClientUpdateOptions contains the optional parameters for the SapLandscapeMonitorClient.Update
//     method.
func (client *SapLandscapeMonitorClient) Update(ctx context.Context, resourceGroupName string, monitorName string, sapLandscapeMonitorParameter SapLandscapeMonitor, options *SapLandscapeMonitorClientUpdateOptions) (SapLandscapeMonitorClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, monitorName, sapLandscapeMonitorParameter, options)
	if err != nil {
		return SapLandscapeMonitorClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SapLandscapeMonitorClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SapLandscapeMonitorClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SapLandscapeMonitorClient) updateCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, sapLandscapeMonitorParameter SapLandscapeMonitor, options *SapLandscapeMonitorClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/monitors/{monitorName}/sapLandscapeMonitor/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, sapLandscapeMonitorParameter)
}

// updateHandleResponse handles the Update response.
func (client *SapLandscapeMonitorClient) updateHandleResponse(resp *http.Response) (SapLandscapeMonitorClientUpdateResponse, error) {
	result := SapLandscapeMonitorClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SapLandscapeMonitor); err != nil {
		return SapLandscapeMonitorClientUpdateResponse{}, err
	}
	return result, nil
}