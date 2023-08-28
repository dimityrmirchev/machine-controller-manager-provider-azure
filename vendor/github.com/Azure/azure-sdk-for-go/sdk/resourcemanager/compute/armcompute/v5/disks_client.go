//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// DisksClient contains the methods for the Disks group.
// Don't use this type directly, use NewDisksClient() instead.
type DisksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDisksClient creates a new instance of DisksClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDisksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DisksClient, error) {
	cl, err := arm.NewClient(moduleName+".DisksClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DisksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a disk.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-02
//   - resourceGroupName - The name of the resource group.
//   - diskName - The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported
//     characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80
//     characters.
//   - disk - Disk object supplied in the body of the Put disk operation.
//   - options - DisksClientBeginCreateOrUpdateOptions contains the optional parameters for the DisksClient.BeginCreateOrUpdate
//     method.
func (client *DisksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, diskName string, disk Disk, options *DisksClientBeginCreateOrUpdateOptions) (*runtime.Poller[DisksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, diskName, disk, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DisksClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DisksClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a disk.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-02
func (client *DisksClient) createOrUpdate(ctx context.Context, resourceGroupName string, diskName string, disk Disk, options *DisksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DisksClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, diskName, disk, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DisksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, diskName string, disk Disk, options *DisksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskName == "" {
		return nil, errors.New("parameter diskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, disk); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a disk.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-02
//   - resourceGroupName - The name of the resource group.
//   - diskName - The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported
//     characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80
//     characters.
//   - options - DisksClientBeginDeleteOptions contains the optional parameters for the DisksClient.BeginDelete method.
func (client *DisksClient) BeginDelete(ctx context.Context, resourceGroupName string, diskName string, options *DisksClientBeginDeleteOptions) (*runtime.Poller[DisksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, diskName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DisksClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DisksClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a disk.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-02
func (client *DisksClient) deleteOperation(ctx context.Context, resourceGroupName string, diskName string, options *DisksClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DisksClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, diskName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DisksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, diskName string, options *DisksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskName == "" {
		return nil, errors.New("parameter diskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets information about a disk.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-02
//   - resourceGroupName - The name of the resource group.
//   - diskName - The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported
//     characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80
//     characters.
//   - options - DisksClientGetOptions contains the optional parameters for the DisksClient.Get method.
func (client *DisksClient) Get(ctx context.Context, resourceGroupName string, diskName string, options *DisksClientGetOptions) (DisksClientGetResponse, error) {
	var err error
	const operationName = "DisksClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, diskName, options)
	if err != nil {
		return DisksClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DisksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DisksClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DisksClient) getCreateRequest(ctx context.Context, resourceGroupName string, diskName string, options *DisksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskName == "" {
		return nil, errors.New("parameter diskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DisksClient) getHandleResponse(resp *http.Response) (DisksClientGetResponse, error) {
	result := DisksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Disk); err != nil {
		return DisksClientGetResponse{}, err
	}
	return result, nil
}

// BeginGrantAccess - Grants access to a disk.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-02
//   - resourceGroupName - The name of the resource group.
//   - diskName - The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported
//     characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80
//     characters.
//   - grantAccessData - Access data object supplied in the body of the get disk access operation.
//   - options - DisksClientBeginGrantAccessOptions contains the optional parameters for the DisksClient.BeginGrantAccess method.
func (client *DisksClient) BeginGrantAccess(ctx context.Context, resourceGroupName string, diskName string, grantAccessData GrantAccessData, options *DisksClientBeginGrantAccessOptions) (*runtime.Poller[DisksClientGrantAccessResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.grantAccess(ctx, resourceGroupName, diskName, grantAccessData, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DisksClientGrantAccessResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DisksClientGrantAccessResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// GrantAccess - Grants access to a disk.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-02
func (client *DisksClient) grantAccess(ctx context.Context, resourceGroupName string, diskName string, grantAccessData GrantAccessData, options *DisksClientBeginGrantAccessOptions) (*http.Response, error) {
	var err error
	const operationName = "DisksClient.BeginGrantAccess"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.grantAccessCreateRequest(ctx, resourceGroupName, diskName, grantAccessData, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// grantAccessCreateRequest creates the GrantAccess request.
func (client *DisksClient) grantAccessCreateRequest(ctx context.Context, resourceGroupName string, diskName string, grantAccessData GrantAccessData, options *DisksClientBeginGrantAccessOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}/beginGetAccess"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskName == "" {
		return nil, errors.New("parameter diskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, grantAccessData); err != nil {
		return nil, err
	}
	return req, nil
}

// NewListPager - Lists all the disks under a subscription.
//
// Generated from API version 2023-01-02
//   - options - DisksClientListOptions contains the optional parameters for the DisksClient.NewListPager method.
func (client *DisksClient) NewListPager(options *DisksClientListOptions) *runtime.Pager[DisksClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DisksClientListResponse]{
		More: func(page DisksClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DisksClientListResponse) (DisksClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DisksClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DisksClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DisksClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DisksClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DisksClient) listCreateRequest(ctx context.Context, options *DisksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/disks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DisksClient) listHandleResponse(resp *http.Response) (DisksClientListResponse, error) {
	result := DisksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskList); err != nil {
		return DisksClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all the disks under a resource group.
//
// Generated from API version 2023-01-02
//   - resourceGroupName - The name of the resource group.
//   - options - DisksClientListByResourceGroupOptions contains the optional parameters for the DisksClient.NewListByResourceGroupPager
//     method.
func (client *DisksClient) NewListByResourceGroupPager(resourceGroupName string, options *DisksClientListByResourceGroupOptions) *runtime.Pager[DisksClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DisksClientListByResourceGroupResponse]{
		More: func(page DisksClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DisksClientListByResourceGroupResponse) (DisksClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DisksClient.NewListByResourceGroupPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DisksClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DisksClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DisksClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DisksClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DisksClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DisksClient) listByResourceGroupHandleResponse(resp *http.Response) (DisksClientListByResourceGroupResponse, error) {
	result := DisksClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskList); err != nil {
		return DisksClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginRevokeAccess - Revokes access to a disk.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-02
//   - resourceGroupName - The name of the resource group.
//   - diskName - The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported
//     characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80
//     characters.
//   - options - DisksClientBeginRevokeAccessOptions contains the optional parameters for the DisksClient.BeginRevokeAccess method.
func (client *DisksClient) BeginRevokeAccess(ctx context.Context, resourceGroupName string, diskName string, options *DisksClientBeginRevokeAccessOptions) (*runtime.Poller[DisksClientRevokeAccessResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.revokeAccess(ctx, resourceGroupName, diskName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DisksClientRevokeAccessResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DisksClientRevokeAccessResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// RevokeAccess - Revokes access to a disk.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-02
func (client *DisksClient) revokeAccess(ctx context.Context, resourceGroupName string, diskName string, options *DisksClientBeginRevokeAccessOptions) (*http.Response, error) {
	var err error
	const operationName = "DisksClient.BeginRevokeAccess"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.revokeAccessCreateRequest(ctx, resourceGroupName, diskName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// revokeAccessCreateRequest creates the RevokeAccess request.
func (client *DisksClient) revokeAccessCreateRequest(ctx context.Context, resourceGroupName string, diskName string, options *DisksClientBeginRevokeAccessOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}/endGetAccess"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskName == "" {
		return nil, errors.New("parameter diskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginUpdate - Updates (patches) a disk.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-02
//   - resourceGroupName - The name of the resource group.
//   - diskName - The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported
//     characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80
//     characters.
//   - disk - Disk object supplied in the body of the Patch disk operation.
//   - options - DisksClientBeginUpdateOptions contains the optional parameters for the DisksClient.BeginUpdate method.
func (client *DisksClient) BeginUpdate(ctx context.Context, resourceGroupName string, diskName string, disk DiskUpdate, options *DisksClientBeginUpdateOptions) (*runtime.Poller[DisksClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, diskName, disk, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DisksClientUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DisksClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates (patches) a disk.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-02
func (client *DisksClient) update(ctx context.Context, resourceGroupName string, diskName string, disk DiskUpdate, options *DisksClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DisksClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, diskName, disk, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *DisksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, diskName string, disk DiskUpdate, options *DisksClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskName == "" {
		return nil, errors.New("parameter diskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, disk); err != nil {
		return nil, err
	}
	return req, nil
}
