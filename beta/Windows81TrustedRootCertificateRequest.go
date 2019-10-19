// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Windows81TrustedRootCertificateRequestBuilder is request builder for Windows81TrustedRootCertificate
type Windows81TrustedRootCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns Windows81TrustedRootCertificateRequest
func (b *Windows81TrustedRootCertificateRequestBuilder) Request() *Windows81TrustedRootCertificateRequest {
	return &Windows81TrustedRootCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Windows81TrustedRootCertificateRequest is request for Windows81TrustedRootCertificate
type Windows81TrustedRootCertificateRequest struct{ BaseRequest }

// Do performs HTTP request for Windows81TrustedRootCertificate
func (r *Windows81TrustedRootCertificateRequest) Do(method, path string, reqObj interface{}) (resObj *Windows81TrustedRootCertificate, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Windows81TrustedRootCertificate
func (r *Windows81TrustedRootCertificateRequest) Get() (*Windows81TrustedRootCertificate, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Windows81TrustedRootCertificate
func (r *Windows81TrustedRootCertificateRequest) Update(reqObj *Windows81TrustedRootCertificate) (*Windows81TrustedRootCertificate, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Windows81TrustedRootCertificate
func (r *Windows81TrustedRootCertificateRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}
