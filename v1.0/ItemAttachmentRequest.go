// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ItemAttachmentRequestBuilder is request builder for ItemAttachment
type ItemAttachmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns ItemAttachmentRequest
func (b *ItemAttachmentRequestBuilder) Request() *ItemAttachmentRequest {
	return &ItemAttachmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ItemAttachmentRequest is request for ItemAttachment
type ItemAttachmentRequest struct{ BaseRequest }

// Do performs HTTP request for ItemAttachment
func (r *ItemAttachmentRequest) Do(method, path string, reqObj interface{}) (resObj *ItemAttachment, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ItemAttachment
func (r *ItemAttachmentRequest) Get() (*ItemAttachment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ItemAttachment
func (r *ItemAttachmentRequest) Update(reqObj *ItemAttachment) (*ItemAttachment, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ItemAttachment
func (r *ItemAttachmentRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Item is navigation property
func (b *ItemAttachmentRequestBuilder) Item() *OutlookItemRequestBuilder {
	bb := &OutlookItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/item"
	return bb
}
