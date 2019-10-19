// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// MailFolderRequestBuilder is request builder for MailFolder
type MailFolderRequestBuilder struct{ BaseRequestBuilder }

// Request returns MailFolderRequest
func (b *MailFolderRequestBuilder) Request() *MailFolderRequest {
	return &MailFolderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MailFolderRequest is request for MailFolder
type MailFolderRequest struct{ BaseRequest }

// Do performs HTTP request for MailFolder
func (r *MailFolderRequest) Do(method, path string, reqObj interface{}) (resObj *MailFolder, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for MailFolder
func (r *MailFolderRequest) Get() (*MailFolder, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for MailFolder
func (r *MailFolderRequest) Update(reqObj *MailFolder) (*MailFolder, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for MailFolder
func (r *MailFolderRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// ChildFolders returns request builder for MailFolder collection
func (b *MailFolderRequestBuilder) ChildFolders() *MailFolderChildFoldersCollectionRequestBuilder {
	bb := &MailFolderChildFoldersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/childFolders"
	return bb
}

// MailFolderChildFoldersCollectionRequestBuilder is request builder for MailFolder collection
type MailFolderChildFoldersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MailFolder collection
func (b *MailFolderChildFoldersCollectionRequestBuilder) Request() *MailFolderChildFoldersCollectionRequest {
	return &MailFolderChildFoldersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MailFolder item
func (b *MailFolderChildFoldersCollectionRequestBuilder) ID(id string) *MailFolderRequestBuilder {
	bb := &MailFolderRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MailFolderChildFoldersCollectionRequest is request for MailFolder collection
type MailFolderChildFoldersCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for MailFolder collection
func (r *MailFolderChildFoldersCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *MailFolder, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for MailFolder collection
func (r *MailFolderChildFoldersCollectionRequest) Paging(method, path string, obj interface{}) ([]MailFolder, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MailFolder
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []MailFolder
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

// Get performs GET request for MailFolder collection
func (r *MailFolderChildFoldersCollectionRequest) Get() ([]MailFolder, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for MailFolder collection
func (r *MailFolderChildFoldersCollectionRequest) Add(reqObj *MailFolder) (*MailFolder, error) {
	return r.Do("POST", "", reqObj)
}

// MessageRules returns request builder for MessageRule collection
func (b *MailFolderRequestBuilder) MessageRules() *MailFolderMessageRulesCollectionRequestBuilder {
	bb := &MailFolderMessageRulesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/messageRules"
	return bb
}

// MailFolderMessageRulesCollectionRequestBuilder is request builder for MessageRule collection
type MailFolderMessageRulesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MessageRule collection
func (b *MailFolderMessageRulesCollectionRequestBuilder) Request() *MailFolderMessageRulesCollectionRequest {
	return &MailFolderMessageRulesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MessageRule item
func (b *MailFolderMessageRulesCollectionRequestBuilder) ID(id string) *MessageRuleRequestBuilder {
	bb := &MessageRuleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MailFolderMessageRulesCollectionRequest is request for MessageRule collection
type MailFolderMessageRulesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for MessageRule collection
func (r *MailFolderMessageRulesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *MessageRule, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for MessageRule collection
func (r *MailFolderMessageRulesCollectionRequest) Paging(method, path string, obj interface{}) ([]MessageRule, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MessageRule
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []MessageRule
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

// Get performs GET request for MessageRule collection
func (r *MailFolderMessageRulesCollectionRequest) Get() ([]MessageRule, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for MessageRule collection
func (r *MailFolderMessageRulesCollectionRequest) Add(reqObj *MessageRule) (*MessageRule, error) {
	return r.Do("POST", "", reqObj)
}

// Messages returns request builder for Message collection
func (b *MailFolderRequestBuilder) Messages() *MailFolderMessagesCollectionRequestBuilder {
	bb := &MailFolderMessagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/messages"
	return bb
}

// MailFolderMessagesCollectionRequestBuilder is request builder for Message collection
type MailFolderMessagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Message collection
func (b *MailFolderMessagesCollectionRequestBuilder) Request() *MailFolderMessagesCollectionRequest {
	return &MailFolderMessagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Message item
func (b *MailFolderMessagesCollectionRequestBuilder) ID(id string) *MessageRequestBuilder {
	bb := &MessageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MailFolderMessagesCollectionRequest is request for Message collection
type MailFolderMessagesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for Message collection
func (r *MailFolderMessagesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Message, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Message collection
func (r *MailFolderMessagesCollectionRequest) Paging(method, path string, obj interface{}) ([]Message, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Message
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Message
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

// Get performs GET request for Message collection
func (r *MailFolderMessagesCollectionRequest) Get() ([]Message, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Message collection
func (r *MailFolderMessagesCollectionRequest) Add(reqObj *Message) (*Message, error) {
	return r.Do("POST", "", reqObj)
}

// MultiValueExtendedProperties returns request builder for MultiValueLegacyExtendedProperty collection
func (b *MailFolderRequestBuilder) MultiValueExtendedProperties() *MailFolderMultiValueExtendedPropertiesCollectionRequestBuilder {
	bb := &MailFolderMultiValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/multiValueExtendedProperties"
	return bb
}

// MailFolderMultiValueExtendedPropertiesCollectionRequestBuilder is request builder for MultiValueLegacyExtendedProperty collection
type MailFolderMultiValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MultiValueLegacyExtendedProperty collection
func (b *MailFolderMultiValueExtendedPropertiesCollectionRequestBuilder) Request() *MailFolderMultiValueExtendedPropertiesCollectionRequest {
	return &MailFolderMultiValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MultiValueLegacyExtendedProperty item
func (b *MailFolderMultiValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *MultiValueLegacyExtendedPropertyRequestBuilder {
	bb := &MultiValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MailFolderMultiValueExtendedPropertiesCollectionRequest is request for MultiValueLegacyExtendedProperty collection
type MailFolderMultiValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for MultiValueLegacyExtendedProperty collection
func (r *MailFolderMultiValueExtendedPropertiesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *MultiValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for MultiValueLegacyExtendedProperty collection
func (r *MailFolderMultiValueExtendedPropertiesCollectionRequest) Paging(method, path string, obj interface{}) ([]MultiValueLegacyExtendedProperty, error) {
	req, err := r.NewJSONRequest(method, path, obj)
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
func (r *MailFolderMultiValueExtendedPropertiesCollectionRequest) Get() ([]MultiValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for MultiValueLegacyExtendedProperty collection
func (r *MailFolderMultiValueExtendedPropertiesCollectionRequest) Add(reqObj *MultiValueLegacyExtendedProperty) (*MultiValueLegacyExtendedProperty, error) {
	return r.Do("POST", "", reqObj)
}

// SingleValueExtendedProperties returns request builder for SingleValueLegacyExtendedProperty collection
func (b *MailFolderRequestBuilder) SingleValueExtendedProperties() *MailFolderSingleValueExtendedPropertiesCollectionRequestBuilder {
	bb := &MailFolderSingleValueExtendedPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/singleValueExtendedProperties"
	return bb
}

// MailFolderSingleValueExtendedPropertiesCollectionRequestBuilder is request builder for SingleValueLegacyExtendedProperty collection
type MailFolderSingleValueExtendedPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SingleValueLegacyExtendedProperty collection
func (b *MailFolderSingleValueExtendedPropertiesCollectionRequestBuilder) Request() *MailFolderSingleValueExtendedPropertiesCollectionRequest {
	return &MailFolderSingleValueExtendedPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SingleValueLegacyExtendedProperty item
func (b *MailFolderSingleValueExtendedPropertiesCollectionRequestBuilder) ID(id string) *SingleValueLegacyExtendedPropertyRequestBuilder {
	bb := &SingleValueLegacyExtendedPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MailFolderSingleValueExtendedPropertiesCollectionRequest is request for SingleValueLegacyExtendedProperty collection
type MailFolderSingleValueExtendedPropertiesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for SingleValueLegacyExtendedProperty collection
func (r *MailFolderSingleValueExtendedPropertiesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *SingleValueLegacyExtendedProperty, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for SingleValueLegacyExtendedProperty collection
func (r *MailFolderSingleValueExtendedPropertiesCollectionRequest) Paging(method, path string, obj interface{}) ([]SingleValueLegacyExtendedProperty, error) {
	req, err := r.NewJSONRequest(method, path, obj)
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
func (r *MailFolderSingleValueExtendedPropertiesCollectionRequest) Get() ([]SingleValueLegacyExtendedProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SingleValueLegacyExtendedProperty collection
func (r *MailFolderSingleValueExtendedPropertiesCollectionRequest) Add(reqObj *SingleValueLegacyExtendedProperty) (*SingleValueLegacyExtendedProperty, error) {
	return r.Do("POST", "", reqObj)
}