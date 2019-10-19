// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementEnableLegacyPcManagementRequestParameter undocumented
type DeviceManagementEnableLegacyPcManagementRequestParameter struct {
}

// DeviceManagementEnableAndroidDeviceAdministratorEnrollmentRequestParameter undocumented
type DeviceManagementEnableAndroidDeviceAdministratorEnrollmentRequestParameter struct {
}

// DeviceManagementSendCustomNotificationToCompanyPortalRequestParameter undocumented
type DeviceManagementSendCustomNotificationToCompanyPortalRequestParameter struct {
	// NotificationTitle undocumented
	NotificationTitle *string `json:"notificationTitle,omitempty"`
	// NotificationBody undocumented
	NotificationBody *string `json:"notificationBody,omitempty"`
	// GroupsToNotify undocumented
	GroupsToNotify []string `json:"groupsToNotify,omitempty"`
}

//
type DeviceManagementEnableLegacyPcManagementRequestBuilder struct{ BaseRequestBuilder }

// EnableLegacyPcManagement action undocumented
func (b *DeviceManagementRequestBuilder) EnableLegacyPcManagement(reqObj *DeviceManagementEnableLegacyPcManagementRequestParameter) *DeviceManagementEnableLegacyPcManagementRequestBuilder {
	bb := &DeviceManagementEnableLegacyPcManagementRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/enableLegacyPcManagement"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceManagementEnableLegacyPcManagementRequest struct{ BaseRequest }

//
func (b *DeviceManagementEnableLegacyPcManagementRequestBuilder) Request() *DeviceManagementEnableLegacyPcManagementRequest {
	return &DeviceManagementEnableLegacyPcManagementRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceManagementEnableLegacyPcManagementRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *DeviceManagementEnableLegacyPcManagementRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type DeviceManagementEnableAndroidDeviceAdministratorEnrollmentRequestBuilder struct{ BaseRequestBuilder }

// EnableAndroidDeviceAdministratorEnrollment action undocumented
func (b *DeviceManagementRequestBuilder) EnableAndroidDeviceAdministratorEnrollment(reqObj *DeviceManagementEnableAndroidDeviceAdministratorEnrollmentRequestParameter) *DeviceManagementEnableAndroidDeviceAdministratorEnrollmentRequestBuilder {
	bb := &DeviceManagementEnableAndroidDeviceAdministratorEnrollmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/enableAndroidDeviceAdministratorEnrollment"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceManagementEnableAndroidDeviceAdministratorEnrollmentRequest struct{ BaseRequest }

//
func (b *DeviceManagementEnableAndroidDeviceAdministratorEnrollmentRequestBuilder) Request() *DeviceManagementEnableAndroidDeviceAdministratorEnrollmentRequest {
	return &DeviceManagementEnableAndroidDeviceAdministratorEnrollmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceManagementEnableAndroidDeviceAdministratorEnrollmentRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *DeviceManagementEnableAndroidDeviceAdministratorEnrollmentRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type DeviceManagementSendCustomNotificationToCompanyPortalRequestBuilder struct{ BaseRequestBuilder }

// SendCustomNotificationToCompanyPortal action undocumented
func (b *DeviceManagementRequestBuilder) SendCustomNotificationToCompanyPortal(reqObj *DeviceManagementSendCustomNotificationToCompanyPortalRequestParameter) *DeviceManagementSendCustomNotificationToCompanyPortalRequestBuilder {
	bb := &DeviceManagementSendCustomNotificationToCompanyPortalRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/sendCustomNotificationToCompanyPortal"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceManagementSendCustomNotificationToCompanyPortalRequest struct{ BaseRequest }

//
func (b *DeviceManagementSendCustomNotificationToCompanyPortalRequestBuilder) Request() *DeviceManagementSendCustomNotificationToCompanyPortalRequest {
	return &DeviceManagementSendCustomNotificationToCompanyPortalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceManagementSendCustomNotificationToCompanyPortalRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *DeviceManagementSendCustomNotificationToCompanyPortalRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}
