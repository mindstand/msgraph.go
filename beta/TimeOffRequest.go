// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// TimeOffRequestBuilder is request builder for TimeOff
type TimeOffRequestBuilder struct{ BaseRequestBuilder }

// Request returns TimeOffRequest
func (b *TimeOffRequestBuilder) Request() *TimeOffRequest {
	return &TimeOffRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *TimeOffRequestBuilder) Delta() *TimeOffRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// TimeOffRequest is request for TimeOff
type TimeOffRequest struct{ BaseRequest }

// Get performs GET request for TimeOff
func (r *TimeOffRequest) Get(ctx context.Context) (resObj *TimeOff, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TimeOff
func (r *TimeOffRequest) Update(ctx context.Context, reqObj *TimeOff) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TimeOff
func (r *TimeOffRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
