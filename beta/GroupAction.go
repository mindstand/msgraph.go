// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GroupValidatePropertiesRequestParameter undocumented
type GroupValidatePropertiesRequestParameter struct {
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// OnBehalfOfUserID undocumented
	OnBehalfOfUserID *UUID `json:"onBehalfOfUserId,omitempty"`
}

// GroupCheckGrantedPermissionsForAppRequestParameter undocumented
type GroupCheckGrantedPermissionsForAppRequestParameter struct {
}

// GroupAssignLicenseRequestParameter undocumented
type GroupAssignLicenseRequestParameter struct {
	// AddLicenses undocumented
	AddLicenses []AssignedLicense `json:"addLicenses,omitempty"`
	// RemoveLicenses undocumented
	RemoveLicenses []UUID `json:"removeLicenses,omitempty"`
}

// GroupSubscribeByMailRequestParameter undocumented
type GroupSubscribeByMailRequestParameter struct {
}

// GroupUnsubscribeByMailRequestParameter undocumented
type GroupUnsubscribeByMailRequestParameter struct {
}

// GroupAddFavoriteRequestParameter undocumented
type GroupAddFavoriteRequestParameter struct {
}

// GroupRemoveFavoriteRequestParameter undocumented
type GroupRemoveFavoriteRequestParameter struct {
}

// GroupResetUnseenCountRequestParameter undocumented
type GroupResetUnseenCountRequestParameter struct {
}

// GroupRenewRequestParameter undocumented
type GroupRenewRequestParameter struct {
}

//
type GroupValidatePropertiesRequestBuilder struct{ BaseRequestBuilder }

// ValidateProperties action undocumented
func (b *GroupRequestBuilder) ValidateProperties(reqObj *GroupValidatePropertiesRequestParameter) *GroupValidatePropertiesRequestBuilder {
	bb := &GroupValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupValidatePropertiesRequest struct{ BaseRequest }

//
func (b *GroupValidatePropertiesRequestBuilder) Request() *GroupValidatePropertiesRequest {
	return &GroupValidatePropertiesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupValidatePropertiesRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *GroupValidatePropertiesRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type GroupCheckGrantedPermissionsForAppRequestBuilder struct{ BaseRequestBuilder }

// CheckGrantedPermissionsForApp action undocumented
func (b *GroupRequestBuilder) CheckGrantedPermissionsForApp(reqObj *GroupCheckGrantedPermissionsForAppRequestParameter) *GroupCheckGrantedPermissionsForAppRequestBuilder {
	bb := &GroupCheckGrantedPermissionsForAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/checkGrantedPermissionsForApp"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupCheckGrantedPermissionsForAppRequest struct{ BaseRequest }

//
func (b *GroupCheckGrantedPermissionsForAppRequestBuilder) Request() *GroupCheckGrantedPermissionsForAppRequest {
	return &GroupCheckGrantedPermissionsForAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupCheckGrantedPermissionsForAppRequest) Do(method, path string, reqObj interface{}) (resObj *[]ResourceSpecificPermissionGrant, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *GroupCheckGrantedPermissionsForAppRequest) Paging(method, path string, obj interface{}) ([][]ResourceSpecificPermissionGrant, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]ResourceSpecificPermissionGrant
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]ResourceSpecificPermissionGrant
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
func (r *GroupCheckGrantedPermissionsForAppRequest) Get() ([][]ResourceSpecificPermissionGrant, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

//
type GroupAssignLicenseRequestBuilder struct{ BaseRequestBuilder }

// AssignLicense action undocumented
func (b *GroupRequestBuilder) AssignLicense(reqObj *GroupAssignLicenseRequestParameter) *GroupAssignLicenseRequestBuilder {
	bb := &GroupAssignLicenseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assignLicense"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupAssignLicenseRequest struct{ BaseRequest }

//
func (b *GroupAssignLicenseRequestBuilder) Request() *GroupAssignLicenseRequest {
	return &GroupAssignLicenseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupAssignLicenseRequest) Do(method, path string, reqObj interface{}) (resObj *Group, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *GroupAssignLicenseRequest) Post() (*Group, error) {
	return r.Do("POST", "", r.requestObject)
}

//
type GroupSubscribeByMailRequestBuilder struct{ BaseRequestBuilder }

// SubscribeByMail action undocumented
func (b *GroupRequestBuilder) SubscribeByMail(reqObj *GroupSubscribeByMailRequestParameter) *GroupSubscribeByMailRequestBuilder {
	bb := &GroupSubscribeByMailRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/subscribeByMail"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupSubscribeByMailRequest struct{ BaseRequest }

//
func (b *GroupSubscribeByMailRequestBuilder) Request() *GroupSubscribeByMailRequest {
	return &GroupSubscribeByMailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupSubscribeByMailRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *GroupSubscribeByMailRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type GroupUnsubscribeByMailRequestBuilder struct{ BaseRequestBuilder }

// UnsubscribeByMail action undocumented
func (b *GroupRequestBuilder) UnsubscribeByMail(reqObj *GroupUnsubscribeByMailRequestParameter) *GroupUnsubscribeByMailRequestBuilder {
	bb := &GroupUnsubscribeByMailRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unsubscribeByMail"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupUnsubscribeByMailRequest struct{ BaseRequest }

//
func (b *GroupUnsubscribeByMailRequestBuilder) Request() *GroupUnsubscribeByMailRequest {
	return &GroupUnsubscribeByMailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupUnsubscribeByMailRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *GroupUnsubscribeByMailRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type GroupAddFavoriteRequestBuilder struct{ BaseRequestBuilder }

// AddFavorite action undocumented
func (b *GroupRequestBuilder) AddFavorite(reqObj *GroupAddFavoriteRequestParameter) *GroupAddFavoriteRequestBuilder {
	bb := &GroupAddFavoriteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/addFavorite"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupAddFavoriteRequest struct{ BaseRequest }

//
func (b *GroupAddFavoriteRequestBuilder) Request() *GroupAddFavoriteRequest {
	return &GroupAddFavoriteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupAddFavoriteRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *GroupAddFavoriteRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type GroupRemoveFavoriteRequestBuilder struct{ BaseRequestBuilder }

// RemoveFavorite action undocumented
func (b *GroupRequestBuilder) RemoveFavorite(reqObj *GroupRemoveFavoriteRequestParameter) *GroupRemoveFavoriteRequestBuilder {
	bb := &GroupRemoveFavoriteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/removeFavorite"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupRemoveFavoriteRequest struct{ BaseRequest }

//
func (b *GroupRemoveFavoriteRequestBuilder) Request() *GroupRemoveFavoriteRequest {
	return &GroupRemoveFavoriteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupRemoveFavoriteRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *GroupRemoveFavoriteRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type GroupResetUnseenCountRequestBuilder struct{ BaseRequestBuilder }

// ResetUnseenCount action undocumented
func (b *GroupRequestBuilder) ResetUnseenCount(reqObj *GroupResetUnseenCountRequestParameter) *GroupResetUnseenCountRequestBuilder {
	bb := &GroupResetUnseenCountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/resetUnseenCount"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupResetUnseenCountRequest struct{ BaseRequest }

//
func (b *GroupResetUnseenCountRequestBuilder) Request() *GroupResetUnseenCountRequest {
	return &GroupResetUnseenCountRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupResetUnseenCountRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *GroupResetUnseenCountRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type GroupRenewRequestBuilder struct{ BaseRequestBuilder }

// Renew action undocumented
func (b *GroupRequestBuilder) Renew(reqObj *GroupRenewRequestParameter) *GroupRenewRequestBuilder {
	bb := &GroupRenewRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/renew"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type GroupRenewRequest struct{ BaseRequest }

//
func (b *GroupRenewRequestBuilder) Request() *GroupRenewRequest {
	return &GroupRenewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *GroupRenewRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *GroupRenewRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}
