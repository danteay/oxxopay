package oxxopay

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// RequestData
//
// Map data por post in the request
type RequestData map[string]interface{}

// RequestHeaders
//
// Map extra request headers
type RequestHeaders map[string]string

// RestClient
//
// Main structure for rest request
type RestClient struct{}

// Execute
//
// Make request and return the response
func (rest *RestClient) Execute(method string, url string, data RequestData, headers RequestHeaders) ([]byte, error) {
	if method != "GET" && method != "POST" && method != "PUT" && method != "DELETE" {
		return nil, errors.New("rest: Not supported method")
	}

	var payload *strings.Reader
	if data != nil {
		payloadByte, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		payload = strings.NewReader(string(payloadByte))
	} else {
		payload = strings.NewReader("")
	}

	req, _ := http.NewRequest(method, url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "application/json")

	if headers != nil {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}

	res, errReq := http.DefaultClient.Do(req)
	if errReq != nil {
		return nil, errReq
	}

	defer res.Body.Close()
	body, errRead := ioutil.ReadAll(res.Body)
	return body, errRead
}

// Get
//
// Execute only Get request
func (rest *RestClient) Get(url string, data RequestData, headers RequestHeaders) ([]byte, error) {
	return rest.Execute("GET", url, data, headers)
}

// Post
//
// Execute only Post request
func (rest *RestClient) Post(url string, data RequestData, headers RequestHeaders) ([]byte, error) {
	return rest.Execute("POST", url, data, headers)
}

// Put
//
// Execute only put request
func (rest *RestClient) Put(url string, data RequestData, headers RequestHeaders) ([]byte, error) {
	return rest.Execute("PUT", url, data, headers)
}

// Delete
//
// Execute only delete request
func (rest *RestClient) Delete(url string, data RequestData, headers RequestHeaders) ([]byte, error) {
	return rest.Execute("DELETE", url, data, headers)
}

