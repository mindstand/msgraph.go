// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AndroidForWorkCertificateProfileBaseRequestBuilder is request builder for AndroidForWorkCertificateProfileBase
type AndroidForWorkCertificateProfileBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidForWorkCertificateProfileBaseRequest
func (b *AndroidForWorkCertificateProfileBaseRequestBuilder) Request() *AndroidForWorkCertificateProfileBaseRequest {
	return &AndroidForWorkCertificateProfileBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *AndroidForWorkCertificateProfileBaseRequestBuilder) Delta() *AndroidForWorkCertificateProfileBaseRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// AndroidForWorkCertificateProfileBaseRequest is request for AndroidForWorkCertificateProfileBase
type AndroidForWorkCertificateProfileBaseRequest struct{ BaseRequest }

// Get performs GET request for AndroidForWorkCertificateProfileBase
func (r *AndroidForWorkCertificateProfileBaseRequest) Get(ctx context.Context) (resObj *AndroidForWorkCertificateProfileBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidForWorkCertificateProfileBase
func (r *AndroidForWorkCertificateProfileBaseRequest) Update(ctx context.Context, reqObj *AndroidForWorkCertificateProfileBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidForWorkCertificateProfileBase
func (r *AndroidForWorkCertificateProfileBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RootCertificate is navigation property
func (b *AndroidForWorkCertificateProfileBaseRequestBuilder) RootCertificate() *AndroidForWorkTrustedRootCertificateRequestBuilder {
	bb := &AndroidForWorkTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificate"
	return bb
}
