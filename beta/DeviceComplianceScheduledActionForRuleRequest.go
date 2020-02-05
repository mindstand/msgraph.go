// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DeviceComplianceScheduledActionForRuleRequestBuilder is request builder for DeviceComplianceScheduledActionForRule
type DeviceComplianceScheduledActionForRuleRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceComplianceScheduledActionForRuleRequest
func (b *DeviceComplianceScheduledActionForRuleRequestBuilder) Request() *DeviceComplianceScheduledActionForRuleRequest {
	return &DeviceComplianceScheduledActionForRuleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *DeviceComplianceScheduledActionForRuleRequestBuilder) Delta() *DeviceComplianceScheduledActionForRuleRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// DeviceComplianceScheduledActionForRuleRequest is request for DeviceComplianceScheduledActionForRule
type DeviceComplianceScheduledActionForRuleRequest struct{ BaseRequest }

// Get performs GET request for DeviceComplianceScheduledActionForRule
func (r *DeviceComplianceScheduledActionForRuleRequest) Get(ctx context.Context) (resObj *DeviceComplianceScheduledActionForRule, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceComplianceScheduledActionForRule
func (r *DeviceComplianceScheduledActionForRuleRequest) Update(ctx context.Context, reqObj *DeviceComplianceScheduledActionForRule) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceComplianceScheduledActionForRule
func (r *DeviceComplianceScheduledActionForRuleRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ScheduledActionConfigurations returns request builder for DeviceComplianceActionItem collection
func (b *DeviceComplianceScheduledActionForRuleRequestBuilder) ScheduledActionConfigurations() *DeviceComplianceScheduledActionForRuleScheduledActionConfigurationsCollectionRequestBuilder {
	bb := &DeviceComplianceScheduledActionForRuleScheduledActionConfigurationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/scheduledActionConfigurations"
	return bb
}

// DeviceComplianceScheduledActionForRuleScheduledActionConfigurationsCollectionRequestBuilder is request builder for DeviceComplianceActionItem collection
type DeviceComplianceScheduledActionForRuleScheduledActionConfigurationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceComplianceActionItem collection
func (b *DeviceComplianceScheduledActionForRuleScheduledActionConfigurationsCollectionRequestBuilder) Request() *DeviceComplianceScheduledActionForRuleScheduledActionConfigurationsCollectionRequest {
	return &DeviceComplianceScheduledActionForRuleScheduledActionConfigurationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceComplianceActionItem item
func (b *DeviceComplianceScheduledActionForRuleScheduledActionConfigurationsCollectionRequestBuilder) ID(id string) *DeviceComplianceActionItemRequestBuilder {
	bb := &DeviceComplianceActionItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceComplianceScheduledActionForRuleScheduledActionConfigurationsCollectionRequest is request for DeviceComplianceActionItem collection
type DeviceComplianceScheduledActionForRuleScheduledActionConfigurationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DeviceComplianceActionItem collection
func (r *DeviceComplianceScheduledActionForRuleScheduledActionConfigurationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DeviceComplianceActionItem, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DeviceComplianceActionItem
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []DeviceComplianceActionItem
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for DeviceComplianceActionItem collection
func (r *DeviceComplianceScheduledActionForRuleScheduledActionConfigurationsCollectionRequest) Get(ctx context.Context) ([]DeviceComplianceActionItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DeviceComplianceActionItem collection
func (r *DeviceComplianceScheduledActionForRuleScheduledActionConfigurationsCollectionRequest) Add(ctx context.Context, reqObj *DeviceComplianceActionItem) (resObj *DeviceComplianceActionItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
