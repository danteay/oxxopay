package oxxopay

import (
	"encoding/base64"
)

const apiBaseUri = "https://api.conekta.io"

// OpClient
//
// Main structure for OxxoPay
type OpClient struct {
	apiKey string
	rest   *RestClient
}

// Init
//
// Init data for OxxoPay
func (op *OpClient) Init(apiKey string) {
	op.apiKey = apiKey
	op.rest = new(RestClient)
}

// CreateOrder
//
// Send an order creation with oxxopay
func (op *OpClient) CreateOrder(data RequestData) ([]byte, error) {
	headers := op.getHeaders()

	url := apiBaseUri + "/orders"
	return op.rest.Post(url, data, headers)
}

// GetOrder
//
// Get order information for precreated order
func (op *OpClient) GetOrder(id string) ([]byte, error) {
	headers := op.getHeaders()

	url := apiBaseUri + "/orders/" + id
	return op.rest.Get(url, nil, headers)
}

// GetListOrders
func (op *OpClient) GetListOrders() ([]byte, error) {
	headers := op.getHeaders()

	url := apiBaseUri + "/orders"
	return op.rest.Get(url, nil, headers)
}

// getHeaders
//
// return generic request headers
func (op *OpClient) getHeaders() RequestHeaders {
	auth := base64.StdEncoding.EncodeToString([]byte(op.apiKey + ":"))

	headers := RequestHeaders{
		"accept":        "application/vnd.conekta-v2.0.0+json",
		"content-type":  "application/json",
		"authorization": "Basic " + auth,
	}

	return headers
}
