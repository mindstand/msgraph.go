// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ProjectParticipationRequestBuilder is request builder for ProjectParticipation
type ProjectParticipationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProjectParticipationRequest
func (b *ProjectParticipationRequestBuilder) Request() *ProjectParticipationRequest {
	return &ProjectParticipationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *ProjectParticipationRequestBuilder) Delta() *ProjectParticipationRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// ProjectParticipationRequest is request for ProjectParticipation
type ProjectParticipationRequest struct{ BaseRequest }

// Get performs GET request for ProjectParticipation
func (r *ProjectParticipationRequest) Get(ctx context.Context) (resObj *ProjectParticipation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ProjectParticipation
func (r *ProjectParticipationRequest) Update(ctx context.Context, reqObj *ProjectParticipation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ProjectParticipation
func (r *ProjectParticipationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
