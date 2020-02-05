// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// EnrollmentConfigurationAssignmentRequestBuilder is request builder for EnrollmentConfigurationAssignment
type EnrollmentConfigurationAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns EnrollmentConfigurationAssignmentRequest
func (b *EnrollmentConfigurationAssignmentRequestBuilder) Request() *EnrollmentConfigurationAssignmentRequest {
	return &EnrollmentConfigurationAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *EnrollmentConfigurationAssignmentRequestBuilder) Delta() *EnrollmentConfigurationAssignmentRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// EnrollmentConfigurationAssignmentRequest is request for EnrollmentConfigurationAssignment
type EnrollmentConfigurationAssignmentRequest struct{ BaseRequest }

// Get performs GET request for EnrollmentConfigurationAssignment
func (r *EnrollmentConfigurationAssignmentRequest) Get(ctx context.Context) (resObj *EnrollmentConfigurationAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EnrollmentConfigurationAssignment
func (r *EnrollmentConfigurationAssignmentRequest) Update(ctx context.Context, reqObj *EnrollmentConfigurationAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EnrollmentConfigurationAssignment
func (r *EnrollmentConfigurationAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
