// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsAutopilotSettingsSyncRequestParameter undocumented
type WindowsAutopilotSettingsSyncRequestParameter struct {
}

//
type WindowsAutopilotSettingsSyncRequestBuilder struct{ BaseRequestBuilder }

// Sync action undocumented
func (b *WindowsAutopilotSettingsRequestBuilder) Sync(reqObj *WindowsAutopilotSettingsSyncRequestParameter) *WindowsAutopilotSettingsSyncRequestBuilder {
	bb := &WindowsAutopilotSettingsSyncRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/sync"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WindowsAutopilotSettingsSyncRequest struct{ BaseRequest }

//
func (b *WindowsAutopilotSettingsSyncRequestBuilder) Request() *WindowsAutopilotSettingsSyncRequest {
	return &WindowsAutopilotSettingsSyncRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WindowsAutopilotSettingsSyncRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequest(method, path, reqObj, nil)
}

//
func (r *WindowsAutopilotSettingsSyncRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}