// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CustomerPaymentJournalRequestBuilder is request builder for CustomerPaymentJournal
type CustomerPaymentJournalRequestBuilder struct{ BaseRequestBuilder }

// Request returns CustomerPaymentJournalRequest
func (b *CustomerPaymentJournalRequestBuilder) Request() *CustomerPaymentJournalRequest {
	return &CustomerPaymentJournalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CustomerPaymentJournalRequest is request for CustomerPaymentJournal
type CustomerPaymentJournalRequest struct{ BaseRequest }

// Do performs HTTP request for CustomerPaymentJournal
func (r *CustomerPaymentJournalRequest) Do(method, path string, reqObj interface{}) (resObj *CustomerPaymentJournal, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for CustomerPaymentJournal
func (r *CustomerPaymentJournalRequest) Get() (*CustomerPaymentJournal, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for CustomerPaymentJournal
func (r *CustomerPaymentJournalRequest) Update(reqObj *CustomerPaymentJournal) (*CustomerPaymentJournal, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for CustomerPaymentJournal
func (r *CustomerPaymentJournalRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Account is navigation property
func (b *CustomerPaymentJournalRequestBuilder) Account() *AccountRequestBuilder {
	bb := &AccountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/account"
	return bb
}

// CustomerPayments returns request builder for CustomerPayment collection
func (b *CustomerPaymentJournalRequestBuilder) CustomerPayments() *CustomerPaymentJournalCustomerPaymentsCollectionRequestBuilder {
	bb := &CustomerPaymentJournalCustomerPaymentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/customerPayments"
	return bb
}

// CustomerPaymentJournalCustomerPaymentsCollectionRequestBuilder is request builder for CustomerPayment collection
type CustomerPaymentJournalCustomerPaymentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CustomerPayment collection
func (b *CustomerPaymentJournalCustomerPaymentsCollectionRequestBuilder) Request() *CustomerPaymentJournalCustomerPaymentsCollectionRequest {
	return &CustomerPaymentJournalCustomerPaymentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CustomerPayment item
func (b *CustomerPaymentJournalCustomerPaymentsCollectionRequestBuilder) ID(id string) *CustomerPaymentRequestBuilder {
	bb := &CustomerPaymentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CustomerPaymentJournalCustomerPaymentsCollectionRequest is request for CustomerPayment collection
type CustomerPaymentJournalCustomerPaymentsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for CustomerPayment collection
func (r *CustomerPaymentJournalCustomerPaymentsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *CustomerPayment, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for CustomerPayment collection
func (r *CustomerPaymentJournalCustomerPaymentsCollectionRequest) Paging(method, path string, obj interface{}) ([]CustomerPayment, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []CustomerPayment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []CustomerPayment
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

// Get performs GET request for CustomerPayment collection
func (r *CustomerPaymentJournalCustomerPaymentsCollectionRequest) Get() ([]CustomerPayment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for CustomerPayment collection
func (r *CustomerPaymentJournalCustomerPaymentsCollectionRequest) Add(reqObj *CustomerPayment) (*CustomerPayment, error) {
	return r.Do("POST", "", reqObj)
}
