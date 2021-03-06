// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// GroupValidatePropertiesRequestParameter undocumented
type GroupValidatePropertiesRequestParameter struct {
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// OnBehalfOfUserID undocumented
	OnBehalfOfUserID *UUID `json:"onBehalfOfUserId,omitempty"`
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
func (r *GroupValidatePropertiesRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
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
func (r *GroupSubscribeByMailRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
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
func (r *GroupUnsubscribeByMailRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
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
func (r *GroupAddFavoriteRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
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
func (r *GroupRemoveFavoriteRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
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
func (r *GroupResetUnseenCountRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
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
func (r *GroupRenewRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
