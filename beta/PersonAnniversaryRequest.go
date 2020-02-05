// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// PersonAnniversaryRequestBuilder is request builder for PersonAnniversary
type PersonAnniversaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersonAnniversaryRequest
func (b *PersonAnniversaryRequestBuilder) Request() *PersonAnniversaryRequest {
	return &PersonAnniversaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *PersonAnniversaryRequestBuilder) Delta() *PersonAnniversaryRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// PersonAnniversaryRequest is request for PersonAnniversary
type PersonAnniversaryRequest struct{ BaseRequest }

// Get performs GET request for PersonAnniversary
func (r *PersonAnniversaryRequest) Get(ctx context.Context) (resObj *PersonAnniversary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersonAnniversary
func (r *PersonAnniversaryRequest) Update(ctx context.Context, reqObj *PersonAnniversary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersonAnniversary
func (r *PersonAnniversaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
