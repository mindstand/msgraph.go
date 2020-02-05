// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ContractRequestBuilder is request builder for Contract
type ContractRequestBuilder struct{ BaseRequestBuilder }

// Request returns ContractRequest
func (b *ContractRequestBuilder) Request() *ContractRequest {
	return &ContractRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *ContractRequestBuilder) Delta() *ContractRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// ContractRequest is request for Contract
type ContractRequest struct{ BaseRequest }

// Get performs GET request for Contract
func (r *ContractRequest) Get(ctx context.Context) (resObj *Contract, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Contract
func (r *ContractRequest) Update(ctx context.Context, reqObj *Contract) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Contract
func (r *ContractRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
