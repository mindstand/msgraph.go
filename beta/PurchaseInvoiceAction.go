// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PurchaseInvoicePostRequestParameter undocumented
type PurchaseInvoicePostRequestParameter struct {
}

//
type PurchaseInvoicePostRequestBuilder struct{ BaseRequestBuilder }

// Post action undocumented
func (b *PurchaseInvoiceRequestBuilder) Post(reqObj *PurchaseInvoicePostRequestParameter) *PurchaseInvoicePostRequestBuilder {
	bb := &PurchaseInvoicePostRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/post"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PurchaseInvoicePostRequest struct{ BaseRequest }

//
func (b *PurchaseInvoicePostRequestBuilder) Request() *PurchaseInvoicePostRequest {
	return &PurchaseInvoicePostRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PurchaseInvoicePostRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *PurchaseInvoicePostRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}
