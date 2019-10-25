// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// OfficeClientConfigurationAssignRequestParameter undocumented
type OfficeClientConfigurationAssignRequestParameter struct {
	// OfficeConfigurationAssignments undocumented
	OfficeConfigurationAssignments []OfficeClientConfigurationAssignment `json:"officeConfigurationAssignments,omitempty"`
}

// OfficeClientConfigurationCollectionUpdatePrioritiesRequestParameter undocumented
type OfficeClientConfigurationCollectionUpdatePrioritiesRequestParameter struct {
	// OfficeConfigurationPolicyIDs undocumented
	OfficeConfigurationPolicyIDs []string `json:"officeConfigurationPolicyIds,omitempty"`
	// OfficeConfigurationPriorities undocumented
	OfficeConfigurationPriorities []int `json:"officeConfigurationPriorities,omitempty"`
}

//
type OfficeClientConfigurationCollectionUpdatePrioritiesRequestBuilder struct{ BaseRequestBuilder }

// UpdatePriorities action undocumented
func (b *OfficeConfigurationClientConfigurationsCollectionRequestBuilder) UpdatePriorities(reqObj *OfficeClientConfigurationCollectionUpdatePrioritiesRequestParameter) *OfficeClientConfigurationCollectionUpdatePrioritiesRequestBuilder {
	bb := &OfficeClientConfigurationCollectionUpdatePrioritiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/updatePriorities"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type OfficeClientConfigurationCollectionUpdatePrioritiesRequest struct{ BaseRequest }

//
func (b *OfficeClientConfigurationCollectionUpdatePrioritiesRequestBuilder) Request() *OfficeClientConfigurationCollectionUpdatePrioritiesRequest {
	return &OfficeClientConfigurationCollectionUpdatePrioritiesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *OfficeClientConfigurationCollectionUpdatePrioritiesRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type OfficeClientConfigurationAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *OfficeClientConfigurationRequestBuilder) Assign(reqObj *OfficeClientConfigurationAssignRequestParameter) *OfficeClientConfigurationAssignRequestBuilder {
	bb := &OfficeClientConfigurationAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type OfficeClientConfigurationAssignRequest struct{ BaseRequest }

//
func (b *OfficeClientConfigurationAssignRequestBuilder) Request() *OfficeClientConfigurationAssignRequest {
	return &OfficeClientConfigurationAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *OfficeClientConfigurationAssignRequest) Paging(method, path string, obj interface{}) ([][]OfficeClientConfigurationAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]OfficeClientConfigurationAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]OfficeClientConfigurationAssignment
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
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

//
func (r *OfficeClientConfigurationAssignRequest) Get() ([][]OfficeClientConfigurationAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}
