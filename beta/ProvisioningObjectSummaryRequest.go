// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ProvisioningObjectSummaryRequestBuilder is request builder for ProvisioningObjectSummary
type ProvisioningObjectSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProvisioningObjectSummaryRequest
func (b *ProvisioningObjectSummaryRequestBuilder) Request() *ProvisioningObjectSummaryRequest {
	return &ProvisioningObjectSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *ProvisioningObjectSummaryRequestBuilder) Delta() *ProvisioningObjectSummaryRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// ProvisioningObjectSummaryRequest is request for ProvisioningObjectSummary
type ProvisioningObjectSummaryRequest struct{ BaseRequest }

// Get performs GET request for ProvisioningObjectSummary
func (r *ProvisioningObjectSummaryRequest) Get(ctx context.Context) (resObj *ProvisioningObjectSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ProvisioningObjectSummary
func (r *ProvisioningObjectSummaryRequest) Update(ctx context.Context, reqObj *ProvisioningObjectSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ProvisioningObjectSummary
func (r *ProvisioningObjectSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
