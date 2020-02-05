// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// EducationAssignmentResourceRequestBuilder is request builder for EducationAssignmentResource
type EducationAssignmentResourceRequestBuilder struct{ BaseRequestBuilder }

// Request returns EducationAssignmentResourceRequest
func (b *EducationAssignmentResourceRequestBuilder) Request() *EducationAssignmentResourceRequest {
	return &EducationAssignmentResourceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *EducationAssignmentResourceRequestBuilder) Delta() *EducationAssignmentResourceRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// EducationAssignmentResourceRequest is request for EducationAssignmentResource
type EducationAssignmentResourceRequest struct{ BaseRequest }

// Get performs GET request for EducationAssignmentResource
func (r *EducationAssignmentResourceRequest) Get(ctx context.Context) (resObj *EducationAssignmentResource, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EducationAssignmentResource
func (r *EducationAssignmentResourceRequest) Update(ctx context.Context, reqObj *EducationAssignmentResource) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EducationAssignmentResource
func (r *EducationAssignmentResourceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
