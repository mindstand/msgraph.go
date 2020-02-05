// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DeviceCategoryRequestBuilder is request builder for DeviceCategory
type DeviceCategoryRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceCategoryRequest
func (b *DeviceCategoryRequestBuilder) Request() *DeviceCategoryRequest {
	return &DeviceCategoryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *DeviceCategoryRequestBuilder) Delta() *DeviceCategoryRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// DeviceCategoryRequest is request for DeviceCategory
type DeviceCategoryRequest struct{ BaseRequest }

// Get performs GET request for DeviceCategory
func (r *DeviceCategoryRequest) Get(ctx context.Context) (resObj *DeviceCategory, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceCategory
func (r *DeviceCategoryRequest) Update(ctx context.Context, reqObj *DeviceCategory) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceCategory
func (r *DeviceCategoryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
