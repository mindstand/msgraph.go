// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// SymantecCodeSigningCertificateRequestBuilder is request builder for SymantecCodeSigningCertificate
type SymantecCodeSigningCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns SymantecCodeSigningCertificateRequest
func (b *SymantecCodeSigningCertificateRequestBuilder) Request() *SymantecCodeSigningCertificateRequest {
	return &SymantecCodeSigningCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *SymantecCodeSigningCertificateRequestBuilder) Delta() *SymantecCodeSigningCertificateRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// SymantecCodeSigningCertificateRequest is request for SymantecCodeSigningCertificate
type SymantecCodeSigningCertificateRequest struct{ BaseRequest }

// Get performs GET request for SymantecCodeSigningCertificate
func (r *SymantecCodeSigningCertificateRequest) Get(ctx context.Context) (resObj *SymantecCodeSigningCertificate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SymantecCodeSigningCertificate
func (r *SymantecCodeSigningCertificateRequest) Update(ctx context.Context, reqObj *SymantecCodeSigningCertificate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SymantecCodeSigningCertificate
func (r *SymantecCodeSigningCertificateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
