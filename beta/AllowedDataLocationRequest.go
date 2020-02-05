// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AllowedDataLocationRequestBuilder is request builder for AllowedDataLocation
type AllowedDataLocationRequestBuilder struct{ BaseRequestBuilder }

// Request returns AllowedDataLocationRequest
func (b *AllowedDataLocationRequestBuilder) Request() *AllowedDataLocationRequest {
	return &AllowedDataLocationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *AllowedDataLocationRequestBuilder) Delta() *AllowedDataLocationRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// AllowedDataLocationRequest is request for AllowedDataLocation
type AllowedDataLocationRequest struct{ BaseRequest }

// Get performs GET request for AllowedDataLocation
func (r *AllowedDataLocationRequest) Get(ctx context.Context) (resObj *AllowedDataLocation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AllowedDataLocation
func (r *AllowedDataLocationRequest) Update(ctx context.Context, reqObj *AllowedDataLocation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AllowedDataLocation
func (r *AllowedDataLocationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
