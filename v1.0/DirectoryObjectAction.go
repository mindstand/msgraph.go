// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DirectoryObjectCollectionGetByIDsRequestParameter undocumented
type DirectoryObjectCollectionGetByIDsRequestParameter struct {
	// IDs undocumented
	IDs []string `json:"ids,omitempty"`
	// Types undocumented
	Types []string `json:"types,omitempty"`
}

// DirectoryObjectCollectionValidatePropertiesRequestParameter undocumented
type DirectoryObjectCollectionValidatePropertiesRequestParameter struct {
	// EntityType undocumented
	EntityType *string `json:"entityType,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// OnBehalfOfUserID undocumented
	OnBehalfOfUserID *UUID `json:"onBehalfOfUserId,omitempty"`
}

// DirectoryObjectCheckMemberGroupsRequestParameter undocumented
type DirectoryObjectCheckMemberGroupsRequestParameter struct {
	// GroupIDs undocumented
	GroupIDs []string `json:"groupIds,omitempty"`
}

// DirectoryObjectCheckMemberObjectsRequestParameter undocumented
type DirectoryObjectCheckMemberObjectsRequestParameter struct {
	// IDs undocumented
	IDs []string `json:"ids,omitempty"`
}

// DirectoryObjectGetMemberGroupsRequestParameter undocumented
type DirectoryObjectGetMemberGroupsRequestParameter struct {
	// SecurityEnabledOnly undocumented
	SecurityEnabledOnly *bool `json:"securityEnabledOnly,omitempty"`
}

// DirectoryObjectGetMemberObjectsRequestParameter undocumented
type DirectoryObjectGetMemberObjectsRequestParameter struct {
	// SecurityEnabledOnly undocumented
	SecurityEnabledOnly *bool `json:"securityEnabledOnly,omitempty"`
}

// DirectoryObjectRestoreRequestParameter undocumented
type DirectoryObjectRestoreRequestParameter struct {
}

//
type DirectoryObjectCollectionGetByIDsRequestBuilder struct{ BaseRequestBuilder }

// GetByIDs action undocumented
func (b *DeviceMemberOfCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *DeviceRegisteredOwnersCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *DeviceRegisteredUsersCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *DeviceTransitiveMemberOfCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *DirectoryDeletedItemsCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *DirectoryRoleMembersCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *DomainDomainNameReferencesCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *GroupAcceptedSendersCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *GroupMemberOfCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *GroupMembersCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *GroupMembersWithLicenseErrorsCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *GroupOwnersCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *GroupRejectedSendersCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *GroupTransitiveMemberOfCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *GroupTransitiveMembersCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *UserCreatedObjectsCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *UserDirectReportsCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *UserMemberOfCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *UserOwnedDevicesCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *UserOwnedObjectsCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *UserRegisteredDevicesCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetByIDs action undocumented
func (b *UserTransitiveMemberOfCollectionRequestBuilder) GetByIDs(reqObj *DirectoryObjectCollectionGetByIDsRequestParameter) *DirectoryObjectCollectionGetByIDsRequestBuilder {
	bb := &DirectoryObjectCollectionGetByIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getByIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DirectoryObjectCollectionGetByIDsRequest struct{ BaseRequest }

//
func (b *DirectoryObjectCollectionGetByIDsRequestBuilder) Request() *DirectoryObjectCollectionGetByIDsRequest {
	return &DirectoryObjectCollectionGetByIDsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DirectoryObjectCollectionGetByIDsRequest) Do(method, path string, reqObj interface{}) (resObj *[]DirectoryObject, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *DirectoryObjectCollectionGetByIDsRequest) Paging(method, path string, obj interface{}) ([][]DirectoryObject, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]DirectoryObject
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]DirectoryObject
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
func (r *DirectoryObjectCollectionGetByIDsRequest) Get() ([][]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

//
type DirectoryObjectCollectionValidatePropertiesRequestBuilder struct{ BaseRequestBuilder }

// ValidateProperties action undocumented
func (b *DeviceMemberOfCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *DeviceRegisteredOwnersCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *DeviceRegisteredUsersCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *DeviceTransitiveMemberOfCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *DirectoryDeletedItemsCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *DirectoryRoleMembersCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *DomainDomainNameReferencesCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *GroupAcceptedSendersCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *GroupMemberOfCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *GroupMembersCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *GroupMembersWithLicenseErrorsCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *GroupOwnersCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *GroupRejectedSendersCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *GroupTransitiveMemberOfCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *GroupTransitiveMembersCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *UserCreatedObjectsCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *UserDirectReportsCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *UserMemberOfCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *UserOwnedDevicesCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *UserOwnedObjectsCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *UserRegisteredDevicesCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// ValidateProperties action undocumented
func (b *UserTransitiveMemberOfCollectionRequestBuilder) ValidateProperties(reqObj *DirectoryObjectCollectionValidatePropertiesRequestParameter) *DirectoryObjectCollectionValidatePropertiesRequestBuilder {
	bb := &DirectoryObjectCollectionValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/validateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DirectoryObjectCollectionValidatePropertiesRequest struct{ BaseRequest }

//
func (b *DirectoryObjectCollectionValidatePropertiesRequestBuilder) Request() *DirectoryObjectCollectionValidatePropertiesRequest {
	return &DirectoryObjectCollectionValidatePropertiesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DirectoryObjectCollectionValidatePropertiesRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *DirectoryObjectCollectionValidatePropertiesRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type DirectoryObjectCheckMemberGroupsRequestBuilder struct{ BaseRequestBuilder }

// CheckMemberGroups action undocumented
func (b *DirectoryObjectRequestBuilder) CheckMemberGroups(reqObj *DirectoryObjectCheckMemberGroupsRequestParameter) *DirectoryObjectCheckMemberGroupsRequestBuilder {
	bb := &DirectoryObjectCheckMemberGroupsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/checkMemberGroups"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DirectoryObjectCheckMemberGroupsRequest struct{ BaseRequest }

//
func (b *DirectoryObjectCheckMemberGroupsRequestBuilder) Request() *DirectoryObjectCheckMemberGroupsRequest {
	return &DirectoryObjectCheckMemberGroupsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DirectoryObjectCheckMemberGroupsRequest) Do(method, path string, reqObj interface{}) (resObj *[]string, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *DirectoryObjectCheckMemberGroupsRequest) Paging(method, path string, obj interface{}) ([][]string, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]string
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]string
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
func (r *DirectoryObjectCheckMemberGroupsRequest) Get() ([][]string, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

//
type DirectoryObjectCheckMemberObjectsRequestBuilder struct{ BaseRequestBuilder }

// CheckMemberObjects action undocumented
func (b *DirectoryObjectRequestBuilder) CheckMemberObjects(reqObj *DirectoryObjectCheckMemberObjectsRequestParameter) *DirectoryObjectCheckMemberObjectsRequestBuilder {
	bb := &DirectoryObjectCheckMemberObjectsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/checkMemberObjects"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DirectoryObjectCheckMemberObjectsRequest struct{ BaseRequest }

//
func (b *DirectoryObjectCheckMemberObjectsRequestBuilder) Request() *DirectoryObjectCheckMemberObjectsRequest {
	return &DirectoryObjectCheckMemberObjectsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DirectoryObjectCheckMemberObjectsRequest) Do(method, path string, reqObj interface{}) (resObj *[]string, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *DirectoryObjectCheckMemberObjectsRequest) Paging(method, path string, obj interface{}) ([][]string, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]string
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]string
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
func (r *DirectoryObjectCheckMemberObjectsRequest) Get() ([][]string, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

//
type DirectoryObjectGetMemberGroupsRequestBuilder struct{ BaseRequestBuilder }

// GetMemberGroups action undocumented
func (b *DirectoryObjectRequestBuilder) GetMemberGroups(reqObj *DirectoryObjectGetMemberGroupsRequestParameter) *DirectoryObjectGetMemberGroupsRequestBuilder {
	bb := &DirectoryObjectGetMemberGroupsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getMemberGroups"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DirectoryObjectGetMemberGroupsRequest struct{ BaseRequest }

//
func (b *DirectoryObjectGetMemberGroupsRequestBuilder) Request() *DirectoryObjectGetMemberGroupsRequest {
	return &DirectoryObjectGetMemberGroupsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DirectoryObjectGetMemberGroupsRequest) Do(method, path string, reqObj interface{}) (resObj *[]string, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *DirectoryObjectGetMemberGroupsRequest) Paging(method, path string, obj interface{}) ([][]string, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]string
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]string
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
func (r *DirectoryObjectGetMemberGroupsRequest) Get() ([][]string, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

//
type DirectoryObjectGetMemberObjectsRequestBuilder struct{ BaseRequestBuilder }

// GetMemberObjects action undocumented
func (b *DirectoryObjectRequestBuilder) GetMemberObjects(reqObj *DirectoryObjectGetMemberObjectsRequestParameter) *DirectoryObjectGetMemberObjectsRequestBuilder {
	bb := &DirectoryObjectGetMemberObjectsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getMemberObjects"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DirectoryObjectGetMemberObjectsRequest struct{ BaseRequest }

//
func (b *DirectoryObjectGetMemberObjectsRequestBuilder) Request() *DirectoryObjectGetMemberObjectsRequest {
	return &DirectoryObjectGetMemberObjectsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DirectoryObjectGetMemberObjectsRequest) Do(method, path string, reqObj interface{}) (resObj *[]string, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *DirectoryObjectGetMemberObjectsRequest) Paging(method, path string, obj interface{}) ([][]string, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]string
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]string
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
func (r *DirectoryObjectGetMemberObjectsRequest) Get() ([][]string, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

//
type DirectoryObjectRestoreRequestBuilder struct{ BaseRequestBuilder }

// Restore action undocumented
func (b *DirectoryObjectRequestBuilder) Restore(reqObj *DirectoryObjectRestoreRequestParameter) *DirectoryObjectRestoreRequestBuilder {
	bb := &DirectoryObjectRestoreRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/restore"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DirectoryObjectRestoreRequest struct{ BaseRequest }

//
func (b *DirectoryObjectRestoreRequestBuilder) Request() *DirectoryObjectRestoreRequest {
	return &DirectoryObjectRestoreRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DirectoryObjectRestoreRequest) Do(method, path string, reqObj interface{}) (resObj *DirectoryObject, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *DirectoryObjectRestoreRequest) Post() (*DirectoryObject, error) {
	return r.Do("POST", "", r.requestObject)
}
