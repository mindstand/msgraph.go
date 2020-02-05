// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// SalesQuoteRequestBuilder is request builder for SalesQuote
type SalesQuoteRequestBuilder struct{ BaseRequestBuilder }

// Request returns SalesQuoteRequest
func (b *SalesQuoteRequestBuilder) Request() *SalesQuoteRequest {
	return &SalesQuoteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *SalesQuoteRequestBuilder) Delta() *SalesQuoteRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// SalesQuoteRequest is request for SalesQuote
type SalesQuoteRequest struct{ BaseRequest }

// Get performs GET request for SalesQuote
func (r *SalesQuoteRequest) Get(ctx context.Context) (resObj *SalesQuote, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SalesQuote
func (r *SalesQuoteRequest) Update(ctx context.Context, reqObj *SalesQuote) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SalesQuote
func (r *SalesQuoteRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Currency is navigation property
func (b *SalesQuoteRequestBuilder) Currency() *CurrencyRequestBuilder {
	bb := &CurrencyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/currency"
	return bb
}

// Customer is navigation property
func (b *SalesQuoteRequestBuilder) Customer() *CustomerRequestBuilder {
	bb := &CustomerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/customer"
	return bb
}

// PaymentTerm is navigation property
func (b *SalesQuoteRequestBuilder) PaymentTerm() *PaymentTermRequestBuilder {
	bb := &PaymentTermRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/paymentTerm"
	return bb
}

// SalesQuoteLines returns request builder for SalesQuoteLine collection
func (b *SalesQuoteRequestBuilder) SalesQuoteLines() *SalesQuoteSalesQuoteLinesCollectionRequestBuilder {
	bb := &SalesQuoteSalesQuoteLinesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/salesQuoteLines"
	return bb
}

// SalesQuoteSalesQuoteLinesCollectionRequestBuilder is request builder for SalesQuoteLine collection
type SalesQuoteSalesQuoteLinesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SalesQuoteLine collection
func (b *SalesQuoteSalesQuoteLinesCollectionRequestBuilder) Request() *SalesQuoteSalesQuoteLinesCollectionRequest {
	return &SalesQuoteSalesQuoteLinesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SalesQuoteLine item
func (b *SalesQuoteSalesQuoteLinesCollectionRequestBuilder) ID(id string) *SalesQuoteLineRequestBuilder {
	bb := &SalesQuoteLineRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SalesQuoteSalesQuoteLinesCollectionRequest is request for SalesQuoteLine collection
type SalesQuoteSalesQuoteLinesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SalesQuoteLine collection
func (r *SalesQuoteSalesQuoteLinesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]SalesQuoteLine, error) {
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
	var values []SalesQuoteLine
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
			value  []SalesQuoteLine
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

// Get performs GET request for SalesQuoteLine collection
func (r *SalesQuoteSalesQuoteLinesCollectionRequest) Get(ctx context.Context) ([]SalesQuoteLine, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for SalesQuoteLine collection
func (r *SalesQuoteSalesQuoteLinesCollectionRequest) Add(ctx context.Context, reqObj *SalesQuoteLine) (resObj *SalesQuoteLine, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ShipmentMethod is navigation property
func (b *SalesQuoteRequestBuilder) ShipmentMethod() *ShipmentMethodRequestBuilder {
	bb := &ShipmentMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/shipmentMethod"
	return bb
}
