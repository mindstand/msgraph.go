// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EducationAssignmentPublishRequestParameter undocumented
type EducationAssignmentPublishRequestParameter struct {
}

//
type EducationAssignmentPublishRequestBuilder struct{ BaseRequestBuilder }

// Publish action undocumented
func (b *EducationAssignmentRequestBuilder) Publish(reqObj *EducationAssignmentPublishRequestParameter) *EducationAssignmentPublishRequestBuilder {
	bb := &EducationAssignmentPublishRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/publish"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type EducationAssignmentPublishRequest struct{ BaseRequest }

//
func (b *EducationAssignmentPublishRequestBuilder) Request() *EducationAssignmentPublishRequest {
	return &EducationAssignmentPublishRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *EducationAssignmentPublishRequest) Do(method, path string, reqObj interface{}) (resObj *EducationAssignment, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *EducationAssignmentPublishRequest) Post() (*EducationAssignment, error) {
	return r.Do("POST", "", r.requestObject)
}
