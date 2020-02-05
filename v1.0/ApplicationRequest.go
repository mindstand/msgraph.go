// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// ApplicationRequestBuilder is request builder for Application
type ApplicationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ApplicationRequest
func (b *ApplicationRequestBuilder) Request() *ApplicationRequest {
	return &ApplicationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *ApplicationRequestBuilder) Delta() *ApplicationRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// ApplicationRequest is request for Application
type ApplicationRequest struct{ BaseRequest }

// Get performs GET request for Application
func (r *ApplicationRequest) Get(ctx context.Context) (resObj *Application, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Application
func (r *ApplicationRequest) Update(ctx context.Context, reqObj *Application) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Application
func (r *ApplicationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CreatedOnBehalfOf is navigation property
func (b *ApplicationRequestBuilder) CreatedOnBehalfOf() *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/createdOnBehalfOf"
	return bb
}

// ExtensionProperties returns request builder for ExtensionProperty collection
func (b *ApplicationRequestBuilder) ExtensionProperties() *ApplicationExtensionPropertiesCollectionRequestBuilder {
	bb := &ApplicationExtensionPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/extensionProperties"
	return bb
}

// ApplicationExtensionPropertiesCollectionRequestBuilder is request builder for ExtensionProperty collection
type ApplicationExtensionPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ExtensionProperty collection
func (b *ApplicationExtensionPropertiesCollectionRequestBuilder) Request() *ApplicationExtensionPropertiesCollectionRequest {
	return &ApplicationExtensionPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ExtensionProperty item
func (b *ApplicationExtensionPropertiesCollectionRequestBuilder) ID(id string) *ExtensionPropertyRequestBuilder {
	bb := &ExtensionPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApplicationExtensionPropertiesCollectionRequest is request for ExtensionProperty collection
type ApplicationExtensionPropertiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ExtensionProperty collection
func (r *ApplicationExtensionPropertiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ExtensionProperty, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ExtensionProperty
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ExtensionProperty
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for ExtensionProperty collection
func (r *ApplicationExtensionPropertiesCollectionRequest) Get(ctx context.Context) ([]ExtensionProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ExtensionProperty collection
func (r *ApplicationExtensionPropertiesCollectionRequest) Add(ctx context.Context, reqObj *ExtensionProperty) (resObj *ExtensionProperty, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Owners returns request builder for DirectoryObject collection
func (b *ApplicationRequestBuilder) Owners() *ApplicationOwnersCollectionRequestBuilder {
	bb := &ApplicationOwnersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/owners"
	return bb
}

// ApplicationOwnersCollectionRequestBuilder is request builder for DirectoryObject collection
type ApplicationOwnersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ApplicationOwnersCollectionRequestBuilder) Request() *ApplicationOwnersCollectionRequest {
	return &ApplicationOwnersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ApplicationOwnersCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApplicationOwnersCollectionRequest is request for DirectoryObject collection
type ApplicationOwnersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ApplicationOwnersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []DirectoryObject
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for DirectoryObject collection
func (r *ApplicationOwnersCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *ApplicationOwnersCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
