// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// PolicySetAssignmentRequestBuilder is request builder for PolicySetAssignment
type PolicySetAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns PolicySetAssignmentRequest
func (b *PolicySetAssignmentRequestBuilder) Request() *PolicySetAssignmentRequest {
	return &PolicySetAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *PolicySetAssignmentRequestBuilder) Delta() *PolicySetAssignmentRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// PolicySetAssignmentRequest is request for PolicySetAssignment
type PolicySetAssignmentRequest struct{ BaseRequest }

// Get performs GET request for PolicySetAssignment
func (r *PolicySetAssignmentRequest) Get(ctx context.Context) (resObj *PolicySetAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PolicySetAssignment
func (r *PolicySetAssignmentRequest) Update(ctx context.Context, reqObj *PolicySetAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PolicySetAssignment
func (r *PolicySetAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
