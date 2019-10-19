// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// InformationProtectionLabelCollectionExtractLabelRequestParameter undocumented
type InformationProtectionLabelCollectionExtractLabelRequestParameter struct {
	// ContentInfo undocumented
	ContentInfo *ContentInfo `json:"contentInfo,omitempty"`
}

// InformationProtectionLabelCollectionEvaluateApplicationRequestParameter undocumented
type InformationProtectionLabelCollectionEvaluateApplicationRequestParameter struct {
	// ContentInfo undocumented
	ContentInfo *ContentInfo `json:"contentInfo,omitempty"`
	// LabelingOptions undocumented
	LabelingOptions *LabelingOptions `json:"labelingOptions,omitempty"`
}

// InformationProtectionLabelCollectionEvaluateRemovalRequestParameter undocumented
type InformationProtectionLabelCollectionEvaluateRemovalRequestParameter struct {
	// ContentInfo undocumented
	ContentInfo *ContentInfo `json:"contentInfo,omitempty"`
	// DowngradeJustification undocumented
	DowngradeJustification *DowngradeJustification `json:"downgradeJustification,omitempty"`
}

// InformationProtectionLabelCollectionEvaluateClassificationResultsRequestParameter undocumented
type InformationProtectionLabelCollectionEvaluateClassificationResultsRequestParameter struct {
	// ContentInfo undocumented
	ContentInfo *ContentInfo `json:"contentInfo,omitempty"`
	// ClassificationResults undocumented
	ClassificationResults []ClassificationResult `json:"classificationResults,omitempty"`
}

//
type InformationProtectionLabelCollectionExtractLabelRequestBuilder struct{ BaseRequestBuilder }

// ExtractLabel action undocumented
func (b *InformationProtectionPolicyLabelsCollectionRequestBuilder) ExtractLabel(reqObj *InformationProtectionLabelCollectionExtractLabelRequestParameter) *InformationProtectionLabelCollectionExtractLabelRequestBuilder {
	bb := &InformationProtectionLabelCollectionExtractLabelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/extractLabel"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type InformationProtectionLabelCollectionExtractLabelRequest struct{ BaseRequest }

//
func (b *InformationProtectionLabelCollectionExtractLabelRequestBuilder) Request() *InformationProtectionLabelCollectionExtractLabelRequest {
	return &InformationProtectionLabelCollectionExtractLabelRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *InformationProtectionLabelCollectionExtractLabelRequest) Do(method, path string, reqObj interface{}) (resObj *InformationProtectionContentLabel, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *InformationProtectionLabelCollectionExtractLabelRequest) Post() (*InformationProtectionContentLabel, error) {
	return r.Do("POST", "", r.requestObject)
}

//
type InformationProtectionLabelCollectionEvaluateApplicationRequestBuilder struct{ BaseRequestBuilder }

// EvaluateApplication action undocumented
func (b *InformationProtectionPolicyLabelsCollectionRequestBuilder) EvaluateApplication(reqObj *InformationProtectionLabelCollectionEvaluateApplicationRequestParameter) *InformationProtectionLabelCollectionEvaluateApplicationRequestBuilder {
	bb := &InformationProtectionLabelCollectionEvaluateApplicationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/evaluateApplication"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type InformationProtectionLabelCollectionEvaluateApplicationRequest struct{ BaseRequest }

//
func (b *InformationProtectionLabelCollectionEvaluateApplicationRequestBuilder) Request() *InformationProtectionLabelCollectionEvaluateApplicationRequest {
	return &InformationProtectionLabelCollectionEvaluateApplicationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *InformationProtectionLabelCollectionEvaluateApplicationRequest) Do(method, path string, reqObj interface{}) (resObj *[]InformationProtectionAction, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *InformationProtectionLabelCollectionEvaluateApplicationRequest) Paging(method, path string, obj interface{}) ([][]InformationProtectionAction, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]InformationProtectionAction
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]InformationProtectionAction
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
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
func (r *InformationProtectionLabelCollectionEvaluateApplicationRequest) Get() ([][]InformationProtectionAction, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

//
type InformationProtectionLabelCollectionEvaluateRemovalRequestBuilder struct{ BaseRequestBuilder }

// EvaluateRemoval action undocumented
func (b *InformationProtectionPolicyLabelsCollectionRequestBuilder) EvaluateRemoval(reqObj *InformationProtectionLabelCollectionEvaluateRemovalRequestParameter) *InformationProtectionLabelCollectionEvaluateRemovalRequestBuilder {
	bb := &InformationProtectionLabelCollectionEvaluateRemovalRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/evaluateRemoval"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type InformationProtectionLabelCollectionEvaluateRemovalRequest struct{ BaseRequest }

//
func (b *InformationProtectionLabelCollectionEvaluateRemovalRequestBuilder) Request() *InformationProtectionLabelCollectionEvaluateRemovalRequest {
	return &InformationProtectionLabelCollectionEvaluateRemovalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *InformationProtectionLabelCollectionEvaluateRemovalRequest) Do(method, path string, reqObj interface{}) (resObj *[]InformationProtectionAction, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *InformationProtectionLabelCollectionEvaluateRemovalRequest) Paging(method, path string, obj interface{}) ([][]InformationProtectionAction, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]InformationProtectionAction
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]InformationProtectionAction
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
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
func (r *InformationProtectionLabelCollectionEvaluateRemovalRequest) Get() ([][]InformationProtectionAction, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

//
type InformationProtectionLabelCollectionEvaluateClassificationResultsRequestBuilder struct{ BaseRequestBuilder }

// EvaluateClassificationResults action undocumented
func (b *InformationProtectionPolicyLabelsCollectionRequestBuilder) EvaluateClassificationResults(reqObj *InformationProtectionLabelCollectionEvaluateClassificationResultsRequestParameter) *InformationProtectionLabelCollectionEvaluateClassificationResultsRequestBuilder {
	bb := &InformationProtectionLabelCollectionEvaluateClassificationResultsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/evaluateClassificationResults"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type InformationProtectionLabelCollectionEvaluateClassificationResultsRequest struct{ BaseRequest }

//
func (b *InformationProtectionLabelCollectionEvaluateClassificationResultsRequestBuilder) Request() *InformationProtectionLabelCollectionEvaluateClassificationResultsRequest {
	return &InformationProtectionLabelCollectionEvaluateClassificationResultsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *InformationProtectionLabelCollectionEvaluateClassificationResultsRequest) Do(method, path string, reqObj interface{}) (resObj *[]InformationProtectionAction, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *InformationProtectionLabelCollectionEvaluateClassificationResultsRequest) Paging(method, path string, obj interface{}) ([][]InformationProtectionAction, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]InformationProtectionAction
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]InformationProtectionAction
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
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
func (r *InformationProtectionLabelCollectionEvaluateClassificationResultsRequest) Get() ([][]InformationProtectionAction, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}
