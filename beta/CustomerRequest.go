// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CustomerRequestBuilder is request builder for Customer
type CustomerRequestBuilder struct{ BaseRequestBuilder }

// Request returns CustomerRequest
func (b *CustomerRequestBuilder) Request() *CustomerRequest {
	return &CustomerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CustomerRequest is request for Customer
type CustomerRequest struct{ BaseRequest }

// Do performs HTTP request for Customer
func (r *CustomerRequest) Do(method, path string, reqObj interface{}) (resObj *Customer, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Customer
func (r *CustomerRequest) Get() (*Customer, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Customer
func (r *CustomerRequest) Update(reqObj *Customer) (*Customer, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Customer
func (r *CustomerRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Currency is navigation property
func (b *CustomerRequestBuilder) Currency() *CurrencyRequestBuilder {
	bb := &CurrencyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/currency"
	return bb
}

// PaymentMethod is navigation property
func (b *CustomerRequestBuilder) PaymentMethod() *PaymentMethodRequestBuilder {
	bb := &PaymentMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/paymentMethod"
	return bb
}

// PaymentTerm is navigation property
func (b *CustomerRequestBuilder) PaymentTerm() *PaymentTermRequestBuilder {
	bb := &PaymentTermRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/paymentTerm"
	return bb
}

// Picture returns request builder for Picture collection
func (b *CustomerRequestBuilder) Picture() *CustomerPictureCollectionRequestBuilder {
	bb := &CustomerPictureCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/picture"
	return bb
}

// CustomerPictureCollectionRequestBuilder is request builder for Picture collection
type CustomerPictureCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Picture collection
func (b *CustomerPictureCollectionRequestBuilder) Request() *CustomerPictureCollectionRequest {
	return &CustomerPictureCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Picture item
func (b *CustomerPictureCollectionRequestBuilder) ID(id string) *PictureRequestBuilder {
	bb := &PictureRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CustomerPictureCollectionRequest is request for Picture collection
type CustomerPictureCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for Picture collection
func (r *CustomerPictureCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Picture, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Picture collection
func (r *CustomerPictureCollectionRequest) Paging(method, path string, obj interface{}) ([]Picture, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Picture
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Picture
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for Picture collection
func (r *CustomerPictureCollectionRequest) Get() ([]Picture, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Picture collection
func (r *CustomerPictureCollectionRequest) Add(reqObj *Picture) (*Picture, error) {
	return r.Do("POST", "", reqObj)
}

// ShipmentMethod is navigation property
func (b *CustomerRequestBuilder) ShipmentMethod() *ShipmentMethodRequestBuilder {
	bb := &ShipmentMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/shipmentMethod"
	return bb
}
