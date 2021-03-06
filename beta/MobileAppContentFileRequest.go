// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// MobileAppContentFileRequestBuilder is request builder for MobileAppContentFile
type MobileAppContentFileRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppContentFileRequest
func (b *MobileAppContentFileRequestBuilder) Request() *MobileAppContentFileRequest {
	return &MobileAppContentFileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppContentFileRequest is request for MobileAppContentFile
type MobileAppContentFileRequest struct{ BaseRequest }

// Get performs GET request for MobileAppContentFile
func (r *MobileAppContentFileRequest) Get(ctx context.Context) (resObj *MobileAppContentFile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileAppContentFile
func (r *MobileAppContentFileRequest) Update(ctx context.Context, reqObj *MobileAppContentFile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileAppContentFile
func (r *MobileAppContentFileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
