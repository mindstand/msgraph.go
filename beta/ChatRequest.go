// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ChatRequestBuilder is request builder for Chat
type ChatRequestBuilder struct{ BaseRequestBuilder }

// Request returns ChatRequest
func (b *ChatRequestBuilder) Request() *ChatRequest {
	return &ChatRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ChatRequest is request for Chat
type ChatRequest struct{ BaseRequest }

// Do performs HTTP request for Chat
func (r *ChatRequest) Do(method, path string, reqObj interface{}) (resObj *Chat, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Chat
func (r *ChatRequest) Get() (*Chat, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Chat
func (r *ChatRequest) Update(reqObj *Chat) (*Chat, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Chat
func (r *ChatRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// InstalledApps returns request builder for TeamsAppInstallation collection
func (b *ChatRequestBuilder) InstalledApps() *ChatInstalledAppsCollectionRequestBuilder {
	bb := &ChatInstalledAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/installedApps"
	return bb
}

// ChatInstalledAppsCollectionRequestBuilder is request builder for TeamsAppInstallation collection
type ChatInstalledAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamsAppInstallation collection
func (b *ChatInstalledAppsCollectionRequestBuilder) Request() *ChatInstalledAppsCollectionRequest {
	return &ChatInstalledAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamsAppInstallation item
func (b *ChatInstalledAppsCollectionRequestBuilder) ID(id string) *TeamsAppInstallationRequestBuilder {
	bb := &TeamsAppInstallationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChatInstalledAppsCollectionRequest is request for TeamsAppInstallation collection
type ChatInstalledAppsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for TeamsAppInstallation collection
func (r *ChatInstalledAppsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *TeamsAppInstallation, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for TeamsAppInstallation collection
func (r *ChatInstalledAppsCollectionRequest) Paging(method, path string, obj interface{}) ([]TeamsAppInstallation, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TeamsAppInstallation
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []TeamsAppInstallation
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

// Get performs GET request for TeamsAppInstallation collection
func (r *ChatInstalledAppsCollectionRequest) Get() ([]TeamsAppInstallation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for TeamsAppInstallation collection
func (r *ChatInstalledAppsCollectionRequest) Add(reqObj *TeamsAppInstallation) (*TeamsAppInstallation, error) {
	return r.Do("POST", "", reqObj)
}

// Members returns request builder for ConversationMember collection
func (b *ChatRequestBuilder) Members() *ChatMembersCollectionRequestBuilder {
	bb := &ChatMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/members"
	return bb
}

// ChatMembersCollectionRequestBuilder is request builder for ConversationMember collection
type ChatMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ConversationMember collection
func (b *ChatMembersCollectionRequestBuilder) Request() *ChatMembersCollectionRequest {
	return &ChatMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ConversationMember item
func (b *ChatMembersCollectionRequestBuilder) ID(id string) *ConversationMemberRequestBuilder {
	bb := &ConversationMemberRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChatMembersCollectionRequest is request for ConversationMember collection
type ChatMembersCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ConversationMember collection
func (r *ChatMembersCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ConversationMember, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ConversationMember collection
func (r *ChatMembersCollectionRequest) Paging(method, path string, obj interface{}) ([]ConversationMember, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ConversationMember
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ConversationMember
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

// Get performs GET request for ConversationMember collection
func (r *ChatMembersCollectionRequest) Get() ([]ConversationMember, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ConversationMember collection
func (r *ChatMembersCollectionRequest) Add(reqObj *ConversationMember) (*ConversationMember, error) {
	return r.Do("POST", "", reqObj)
}

// Messages returns request builder for ChatMessage collection
func (b *ChatRequestBuilder) Messages() *ChatMessagesCollectionRequestBuilder {
	bb := &ChatMessagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/messages"
	return bb
}

// ChatMessagesCollectionRequestBuilder is request builder for ChatMessage collection
type ChatMessagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ChatMessage collection
func (b *ChatMessagesCollectionRequestBuilder) Request() *ChatMessagesCollectionRequest {
	return &ChatMessagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ChatMessage item
func (b *ChatMessagesCollectionRequestBuilder) ID(id string) *ChatMessageRequestBuilder {
	bb := &ChatMessageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChatMessagesCollectionRequest is request for ChatMessage collection
type ChatMessagesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ChatMessage collection
func (r *ChatMessagesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ChatMessage, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ChatMessage collection
func (r *ChatMessagesCollectionRequest) Paging(method, path string, obj interface{}) ([]ChatMessage, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ChatMessage
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ChatMessage
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

// Get performs GET request for ChatMessage collection
func (r *ChatMessagesCollectionRequest) Get() ([]ChatMessage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ChatMessage collection
func (r *ChatMessagesCollectionRequest) Add(reqObj *ChatMessage) (*ChatMessage, error) {
	return r.Do("POST", "", reqObj)
}
