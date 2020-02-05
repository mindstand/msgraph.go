// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// InformationProtectionRequestBuilder is request builder for InformationProtection
type InformationProtectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns InformationProtectionRequest
func (b *InformationProtectionRequestBuilder) Request() *InformationProtectionRequest {
	return &InformationProtectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *InformationProtectionRequestBuilder) Delta() *InformationProtectionRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// InformationProtectionRequest is request for InformationProtection
type InformationProtectionRequest struct{ BaseRequest }

// Get performs GET request for InformationProtection
func (r *InformationProtectionRequest) Get(ctx context.Context) (resObj *InformationProtection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for InformationProtection
func (r *InformationProtectionRequest) Update(ctx context.Context, reqObj *InformationProtection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for InformationProtection
func (r *InformationProtectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DataLossPreventionPolicies returns request builder for DataLossPreventionPolicy collection
func (b *InformationProtectionRequestBuilder) DataLossPreventionPolicies() *InformationProtectionDataLossPreventionPoliciesCollectionRequestBuilder {
	bb := &InformationProtectionDataLossPreventionPoliciesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/dataLossPreventionPolicies"
	return bb
}

// InformationProtectionDataLossPreventionPoliciesCollectionRequestBuilder is request builder for DataLossPreventionPolicy collection
type InformationProtectionDataLossPreventionPoliciesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DataLossPreventionPolicy collection
func (b *InformationProtectionDataLossPreventionPoliciesCollectionRequestBuilder) Request() *InformationProtectionDataLossPreventionPoliciesCollectionRequest {
	return &InformationProtectionDataLossPreventionPoliciesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DataLossPreventionPolicy item
func (b *InformationProtectionDataLossPreventionPoliciesCollectionRequestBuilder) ID(id string) *DataLossPreventionPolicyRequestBuilder {
	bb := &DataLossPreventionPolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// InformationProtectionDataLossPreventionPoliciesCollectionRequest is request for DataLossPreventionPolicy collection
type InformationProtectionDataLossPreventionPoliciesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DataLossPreventionPolicy collection
func (r *InformationProtectionDataLossPreventionPoliciesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DataLossPreventionPolicy, error) {
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
	var values []DataLossPreventionPolicy
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
			value  []DataLossPreventionPolicy
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

// Get performs GET request for DataLossPreventionPolicy collection
func (r *InformationProtectionDataLossPreventionPoliciesCollectionRequest) Get(ctx context.Context) ([]DataLossPreventionPolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DataLossPreventionPolicy collection
func (r *InformationProtectionDataLossPreventionPoliciesCollectionRequest) Add(ctx context.Context, reqObj *DataLossPreventionPolicy) (resObj *DataLossPreventionPolicy, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Policy is navigation property
func (b *InformationProtectionRequestBuilder) Policy() *InformationProtectionPolicyRequestBuilder {
	bb := &InformationProtectionPolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/policy"
	return bb
}

// SensitivityLabels returns request builder for SensitivityLabel collection
func (b *InformationProtectionRequestBuilder) SensitivityLabels() *InformationProtectionSensitivityLabelsCollectionRequestBuilder {
	bb := &InformationProtectionSensitivityLabelsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sensitivityLabels"
	return bb
}

// InformationProtectionSensitivityLabelsCollectionRequestBuilder is request builder for SensitivityLabel collection
type InformationProtectionSensitivityLabelsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SensitivityLabel collection
func (b *InformationProtectionSensitivityLabelsCollectionRequestBuilder) Request() *InformationProtectionSensitivityLabelsCollectionRequest {
	return &InformationProtectionSensitivityLabelsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SensitivityLabel item
func (b *InformationProtectionSensitivityLabelsCollectionRequestBuilder) ID(id string) *SensitivityLabelRequestBuilder {
	bb := &SensitivityLabelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// InformationProtectionSensitivityLabelsCollectionRequest is request for SensitivityLabel collection
type InformationProtectionSensitivityLabelsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SensitivityLabel collection
func (r *InformationProtectionSensitivityLabelsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]SensitivityLabel, error) {
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
	var values []SensitivityLabel
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
			value  []SensitivityLabel
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

// Get performs GET request for SensitivityLabel collection
func (r *InformationProtectionSensitivityLabelsCollectionRequest) Get(ctx context.Context) ([]SensitivityLabel, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for SensitivityLabel collection
func (r *InformationProtectionSensitivityLabelsCollectionRequest) Add(ctx context.Context, reqObj *SensitivityLabel) (resObj *SensitivityLabel, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// SensitivityPolicySettings is navigation property
func (b *InformationProtectionRequestBuilder) SensitivityPolicySettings() *SensitivityPolicySettingsRequestBuilder {
	bb := &SensitivityPolicySettingsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sensitivityPolicySettings"
	return bb
}

// ThreatAssessmentRequests returns request builder for ThreatAssessmentRequestObject collection
func (b *InformationProtectionRequestBuilder) ThreatAssessmentRequests() *InformationProtectionThreatAssessmentRequestsCollectionRequestBuilder {
	bb := &InformationProtectionThreatAssessmentRequestsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/threatAssessmentRequests"
	return bb
}

// InformationProtectionThreatAssessmentRequestsCollectionRequestBuilder is request builder for ThreatAssessmentRequestObject collection
type InformationProtectionThreatAssessmentRequestsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ThreatAssessmentRequestObject collection
func (b *InformationProtectionThreatAssessmentRequestsCollectionRequestBuilder) Request() *InformationProtectionThreatAssessmentRequestsCollectionRequest {
	return &InformationProtectionThreatAssessmentRequestsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ThreatAssessmentRequestObject item
func (b *InformationProtectionThreatAssessmentRequestsCollectionRequestBuilder) ID(id string) *ThreatAssessmentRequestObjectRequestBuilder {
	bb := &ThreatAssessmentRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// InformationProtectionThreatAssessmentRequestsCollectionRequest is request for ThreatAssessmentRequestObject collection
type InformationProtectionThreatAssessmentRequestsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ThreatAssessmentRequestObject collection
func (r *InformationProtectionThreatAssessmentRequestsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ThreatAssessmentRequestObject, error) {
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
	var values []ThreatAssessmentRequestObject
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
			value  []ThreatAssessmentRequestObject
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

// Get performs GET request for ThreatAssessmentRequestObject collection
func (r *InformationProtectionThreatAssessmentRequestsCollectionRequest) Get(ctx context.Context) ([]ThreatAssessmentRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ThreatAssessmentRequestObject collection
func (r *InformationProtectionThreatAssessmentRequestsCollectionRequest) Add(ctx context.Context, reqObj *ThreatAssessmentRequestObject) (resObj *ThreatAssessmentRequestObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
