// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// IPSecurityProfileRequestBuilder is request builder for IPSecurityProfile
type IPSecurityProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns IPSecurityProfileRequest
func (b *IPSecurityProfileRequestBuilder) Request() *IPSecurityProfileRequest {
	return &IPSecurityProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *IPSecurityProfileRequestBuilder) Delta() *IPSecurityProfileRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// IPSecurityProfileRequest is request for IPSecurityProfile
type IPSecurityProfileRequest struct{ BaseRequest }

// Get performs GET request for IPSecurityProfile
func (r *IPSecurityProfileRequest) Get(ctx context.Context) (resObj *IPSecurityProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IPSecurityProfile
func (r *IPSecurityProfileRequest) Update(ctx context.Context, reqObj *IPSecurityProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IPSecurityProfile
func (r *IPSecurityProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
