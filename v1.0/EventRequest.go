// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// EventRequestBuilder is request builder for Event
type EventRequestBuilder struct{ BaseRequestBuilder }

// Request returns EventRequest
func (b *EventRequestBuilder) Request() *EventRequest {
	return &EventRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EventRequest is request for Event
type EventRequest struct{ BaseRequest }

// Do performs HTTP request for Event
func (r *EventRequest) Do(method, path string, reqObj interface{}) (resObj *Event, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Event
func (r *EventRequest) Get() (*Event, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Event
func (r *EventRequest) Update(reqObj *Event) (*Event, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Event
func (r *EventRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Attachments returns request builder for Attachment collection
func (b *EventRequestBuilder) Attachments() *EventAttachmentsCollectionRequestBuilder {
	bb := &EventAttachmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/attachments"
	return bb
}

// EventAttachmentsCollectionRequestBuilder is request builder for Attachment collection
type EventAttachmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Attachment collection
func (b *EventAttachmentsCollectionRequestBuilder) Request() *EventAttachmentsCollectionRequest {
	return &EventAttachmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Attachment item
func (b *EventAttachmentsCollectionRequestBuilder) ID(id string) *AttachmentRequestBuilder {
	bb := &AttachmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EventAttachmentsCollectionRequest is request for Attachment collection
type EventAttachmentsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for Attachment collection
func (r *EventAttachmentsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Attachment, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Attachment collection
func (r *EventAttachmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]Attachment, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Attachment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Attachment
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

// Get performs GET request for Attachment collection
func (r *EventAttachmentsCollectionRequest) Get() ([]Attachment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Attachment collection
func (r *EventAttachmentsCollectionRequest) Add(reqObj *Attachment) (*Attachment, error) {
	return r.Do("POST", "", reqObj)
}

// Calendar is navigation property
func (b *EventRequestBuilder) Calendar() *CalendarRequestBuilder {
	bb := &CalendarRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/calendar"
	return bb
}

// Extensions returns request builder for Extension collection
func (b *EventRequestBuilder) Extensions() *EventExtensionsCollectionRequestBuilder {
	bb := &EventExtensionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/extensions"
	return bb
}

// EventExtensionsCollectionRequestBuilder is request builder for Extension collection
type EventExtensionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Extension collection
func (b *EventExtensionsCollectionRequestBuilder) Request() *EventExtensionsCollectionRequest {
	return &EventExtensionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Extension item
func (b *EventExtensionsCollectionRequestBuilder) ID(id string) *ExtensionRequestBuilder {
	bb := &ExtensionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EventExtensionsCollectionRequest is request for Extension collection
type EventExtensionsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for Extension collection
func (r *EventExtensionsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Extension, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Extension collection
func (r *EventExtensionsCollectionRequest) Paging(method, path string, obj interface{}) ([]Extension, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Extension
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Extension
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

// Get performs GET request for Extension collection
func (r *EventExtensionsCollectionRequest) Get() ([]Extension, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Extension collection
func (r *EventExtensionsCollectionRequest) Add(reqObj *Extension) (*Extension, error) {
	return r.Do("POST", "", reqObj)
}

// Instances returns request builder for Event collection
func (b *EventRequestBuilder) Instances() *EventInstancesCollectionRequestBuilder {
	bb := &EventInstancesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/instances"
	return bb
}

// EventInstancesCollectionRequestBuilder is request builder for Event collection
type EventInstancesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Event collection
func (b *EventInstancesCollectionRequestBuilder) Request() *EventInstancesCollectionRequest {
	return &EventInstancesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Event item
func (b *EventInstancesCollectionRequestBuilder) ID(id string) *EventRequestBuilder {
	bb := &EventRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EventInstancesCollectionRequest is request for Event collection
type EventInstancesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for Event collection
func (r *EventInstancesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Event, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Event collection
func (r *EventInstancesCollectionRequest) Paging(method, path string, obj interface{}) ([]Event, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Event
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Event
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

// Get performs GET request for Event collection
func (r *EventInstancesCollectionRequest) Get() ([]Event, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Event collection
func (r *EventInstancesCollectionRequest) Add(reqObj *Event) (*Event, error) {
	return r.Do("POST", "", reqObj)
}

// MultiValueExtendedProperties returns request builder for MultiValueLegacyExtendedProperty collection
func (b *EventRequestBuilder) MultiValueExtendedProperties() *EventMultiValueExtendedPropertiesCollectionRequestBuilder {
	bb := &EventMultiValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/multiValueExtendedProperties"
	return bb
}

// EventMultiValueExtendedPropertiesCollectionRequestBuilder is request builder for MultiValueLegacyExtendedProperty collection
type EventMultiValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MultiValueLegacyExtendedProperty collection
func (b *EventMultiValueExtendedPropertiesCollectionRequestBuilder) Request() *EventMultiValueExtendedPropertiesCollectionRequest {
	return &EventMultiValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MultiValueLegacyExtendedProperty item
func (b *EventMultiValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *MultiValueLegacyExtendedPropertyRequestBuilder {
	bb := &MultiValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EventMultiValueExtendedPropertiesCollectionRequest is request for MultiValueLegacyExtendedProperty collection
type EventMultiValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for MultiValueLegacyExtendedProperty collection
func (r *EventMultiValueExtendedPropertiesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *MultiValueLegacyExtendedProperty, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for MultiValueLegacyExtendedProperty collection
func (r *EventMultiValueExtendedPropertiesCollectionRequest) Paging(method, path string, obj interface{}) ([]MultiValueLegacyExtendedProperty, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MultiValueLegacyExtendedProperty
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []MultiValueLegacyExtendedProperty
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

// Get performs GET request for MultiValueLegacyExtendedProperty collection
func (r *EventMultiValueExtendedPropertiesCollectionRequest) Get() ([]MultiValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for MultiValueLegacyExtendedProperty collection
func (r *EventMultiValueExtendedPropertiesCollectionRequest) Add(reqObj *MultiValueLegacyExtendedProperty) (*MultiValueLegacyExtendedProperty, error) {
	return r.Do("POST", "", reqObj)
}

// SingleValueExtendedProperties returns request builder for SingleValueLegacyExtendedProperty collection
func (b *EventRequestBuilder) SingleValueExtendedProperties() *EventSingleValueExtendedPropertiesCollectionRequestBuilder {
	bb := &EventSingleValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/singleValueExtendedProperties"
	return bb
}

// EventSingleValueExtendedPropertiesCollectionRequestBuilder is request builder for SingleValueLegacyExtendedProperty collection
type EventSingleValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SingleValueLegacyExtendedProperty collection
func (b *EventSingleValueExtendedPropertiesCollectionRequestBuilder) Request() *EventSingleValueExtendedPropertiesCollectionRequest {
	return &EventSingleValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SingleValueLegacyExtendedProperty item
func (b *EventSingleValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *SingleValueLegacyExtendedPropertyRequestBuilder {
	bb := &SingleValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EventSingleValueExtendedPropertiesCollectionRequest is request for SingleValueLegacyExtendedProperty collection
type EventSingleValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for SingleValueLegacyExtendedProperty collection
func (r *EventSingleValueExtendedPropertiesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *SingleValueLegacyExtendedProperty, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for SingleValueLegacyExtendedProperty collection
func (r *EventSingleValueExtendedPropertiesCollectionRequest) Paging(method, path string, obj interface{}) ([]SingleValueLegacyExtendedProperty, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SingleValueLegacyExtendedProperty
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []SingleValueLegacyExtendedProperty
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

// Get performs GET request for SingleValueLegacyExtendedProperty collection
func (r *EventSingleValueExtendedPropertiesCollectionRequest) Get() ([]SingleValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SingleValueLegacyExtendedProperty collection
func (r *EventSingleValueExtendedPropertiesCollectionRequest) Add(reqObj *SingleValueLegacyExtendedProperty) (*SingleValueLegacyExtendedProperty, error) {
	return r.Do("POST", "", reqObj)
}
