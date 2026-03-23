package http_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Nerzal/gocloak/v13"
	"io/ioutil"
	"ms-tagpeak/pkg/dotenv"
	"ms-tagpeak/pkg/logster"
	"net/http"
	"strings"
)

type InternalHttpClientInterface interface {
	Get(url string, headers map[string]string, out interface{}) (*http.Response, error)
	PostJson(url string, headers map[string]string, in interface{}, out interface{}) (*http.Response, error)
	Post(url string, headers map[string]string, in interface{}, out interface{}) (*http.Response, error)
	PutJson(url string, headers map[string]string, in interface{}, out interface{}) (*http.Response, error)
	Put(url string, headers map[string]string, in interface{}, out interface{}) (*http.Response, error)
	Delete(url string, headers map[string]string) (*http.Response, error)
	DeleteWithBody(url string, headers map[string]string, in interface{}) (*http.Response, error)
}

// InternalHttpClient represents an http wrapper which reduces the boiler plate needed to marshall/un-marshall request/response
// bodies by providing friendly CRUD http operations that allow in/out interfaces
type InternalHttpClient struct {
	InternalHttpClient *http.Client
}

// Get issues a GET HTTP request to the specified URL including the headers passed in.
//
// The 'out' param interface is the un-marshall representation of the http response returned
//
// Example on how to invoke the method:
//
//		type Out struct {
//			Id string `json:"id"`
//			Name        string   `json:"name"`
//			Description string   `json:"description"`
//		}
//
//	 out := &Out{}
//	 headers := map[string]string{"header_example": "header_value"}
//	 InternalHttpClient.Get("http://api.com/resource", headers, out)
func (httpClient *InternalHttpClient) Get(url string, headers map[string]string, out interface{}) (*http.Response, error) {
	headers = map[string]string{}
	headers["Authorization"] = "Bearer " + GetInternalKeycloakToken()
	if req, err := httpClient.prepareRequest(http.MethodGet, url, headers, nil); err != nil {
		return nil, err
	} else {
		return httpClient.performRequest(req, out)
	}
}

// PostJson issues a POST to the specified URL including the headers passed in.
// The content type of the body is set to application/json so it doesn't need to be added to the headers passed in
//
// The 'in' param interface is marshall and added to the htp request body.
// The 'out' param interface is the un-marshall representation of the http response returned
//
// Example on how to invoke the method:
//
//		type In struct {
//			Name        string   `json:"name"`
//			Description string   `json:"description"`
//		}
//
//		type Out struct {
//			Id string `json:"id"`
//			Name        string   `json:"name"`
//			Description string   `json:"description"`
//		}
//	 in := &In{}
//	 out := &Out{}
//	 headers := map[string]string{"header_example": "header_value"}
//	 InternalHttpClient.PostJson("http://api.com/resource", headers, in, out)
func (httpClient *InternalHttpClient) PostJson(url string, headers map[string]string, in interface{}, out interface{}) (*http.Response, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	httpClient.addJsonHeader(headers)
	return httpClient.Post(url, headers, in, out)
}

// Post issues a POST HTTP request to the specified URL including the headers passed in.
//
// The in interface is marshall and added to the htp request body.
// The out interface is the un-marshall representation of the http response returned
//
// Example on how to invoke the method:
//
//		type In struct {
//			Name        string   `json:"name"`
//			Description string   `json:"description"`
//		}
//
//		type Out struct {
//			Id string `json:"id"`
//			Name        string   `json:"name"`
//			Description string   `json:"description"`
//		}
//	 in := &In{}
//	 out := &Out{}
//	 headers := map[string]string{"header_example": "header_value"}
//	 InternalHttpClient.Post("http://api.com/resource", headers, in, out)
func (httpClient *InternalHttpClient) Post(url string, headers map[string]string, in interface{}, out interface{}) (*http.Response, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Authorization"] = "Bearer " + GetInternalKeycloakToken()
	if req, err := httpClient.prepareRequest(http.MethodPost, url, headers, in); err != nil {
		return nil, err
	} else {
		return httpClient.performRequest(req, out)
	}
}

// PutJson issues a PUT HTTP request to the specified URL including the headers passed in.
// The content type of the body is set to application/json
//
// The 'in' param interface is marshall and added to the htp request body.
// The 'out' param interface is the un-marshall representation of the http response returned
//
// Example on how to invoke the method:
//
//		type In struct {
//			Name        string   `json:"name"`
//			Description string   `json:"description"`
//		}
//
//		type Out struct {
//			Id string `json:"id"`
//			Name        string   `json:"name"`
//			Description string   `json:"description"`
//		}
//	 in := &In{}
//	 out := &Out{}
//	 headers := map[string]string{"header_example": "header_value"}
//	 InternalHttpClient.PutJson("http://api.com/resource", headers, in, out)
func (httpClient *InternalHttpClient) PutJson(url string, headers map[string]string, in interface{}, out interface{}) (*http.Response, error) {
	headers = map[string]string{}
	headers["Authorization"] = "Bearer " + GetInternalKeycloakToken()
	httpClient.addJsonHeader(headers)
	return httpClient.Put(url, headers, in, out)
}

// Put issues a PUT HTTP request to the specified URL including the headers passed in.
//
// The in interface is marshall and added to the http request body.
// The out interface is the un-marshall representation of the http response returned
//
// Example on how to invoke the method:
//
//		type In struct {
//			Name        string   `json:"name"`
//			Description string   `json:"description"`
//		}
//
//		type Out struct {
//			Id string `json:"id"`
//			Name        string   `json:"name"`
//			Description string   `json:"description"`
//		}
//	 in := &In{}
//	 out := &Out{}
//	 headers := map[string]string{"header_example": "header_value"}
//	 InternalHttpClient.Put("http://api.com/resource", headers, in, out)
func (httpClient *InternalHttpClient) Put(url string, headers map[string]string, in interface{}, out interface{}) (*http.Response, error) {
	headers = map[string]string{}
	headers["Authorization"] = "Bearer " + GetInternalKeycloakToken()

	if req, err := httpClient.prepareRequest(http.MethodPut, url, headers, in); err != nil {
		return nil, err
	} else {
		return httpClient.performRequest(req, out)
	}
}

// Delete issues a DELETE HTTP request to the specified URL including the headers passed in.
func (httpClient *InternalHttpClient) Delete(url string, headers map[string]string) (*http.Response, error) {
	headers = map[string]string{}
	headers["Authorization"] = "Bearer " + GetInternalKeycloakToken()

	if req, err := httpClient.prepareRequest(http.MethodDelete, url, headers, nil); err != nil {
		return nil, err
	} else {
		return httpClient.performRequest(req, nil)
	}
}

func (httpClient *InternalHttpClient) DeleteWithBody(url string, headers map[string]string, in interface{}) (*http.Response, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Authorization"] = "Bearer " + GetInternalKeycloakToken()
	httpClient.addJsonHeader(headers)

	if req, err := httpClient.prepareRequest(http.MethodDelete, url, headers, in); err != nil {
		return nil, err
	} else {
		return httpClient.performRequest(req, nil)
	}
}

func (httpClient *InternalHttpClient) prepareRequest(method, url string, headers map[string]string, in interface{}) (*http.Request, error) {
	var body []byte
	var err error
	if in != nil {
		body, err = json.Marshal(in)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return req, nil
}

func (httpClient *InternalHttpClient) performRequest(req *http.Request, out interface{}) (*http.Response, error) {
	logster.Info(fmt.Sprintf("Performing request %s %s %s\n", req.Method, req.URL, req.Proto))
	resp, err := httpClient.InternalHttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request %s %s %s failed. Response Error: '%s'", req.Method, req.URL, req.Proto, err.Error())
	}
	if out != nil {
		var body []byte
		if body, err = ioutil.ReadAll(resp.Body); err != nil {
			return nil, err
		}
		resp.Body.Close()                                   // close stream so connection is closed gracefully
		resp.Body = ioutil.NopCloser(bytes.NewReader(body)) // create a new reader from bytes read in the response and set the response body (allowing the client to still be able to do res.Body afterwards)
		if len(body) > 0 {
			if err = json.Unmarshal(body, &out); err != nil {
				return resp, fmt.Errorf("unable to unmarshal response body ['%s'] for request = '%s %s %s'. Response = '%s'", err.Error(), req.Method, req.URL, req.Proto, resp.Status)
			}
		} else {
			return resp, fmt.Errorf("expected a response body but response body received was empty for request = '%s %s %s'. Response = '%s'", req.Method, req.URL, req.Proto, resp.Status)
		}
	}
	return resp, nil
}

func (httpClient *InternalHttpClient) addJsonHeader(headers map[string]string) {
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Content-Type"] = "application/json"
}

func GetInternalKeycloakToken() string {

	client := gocloak.NewClient(dotenv.GetEnv("KEYCLOAK_URL"))

	token, err := client.LoginClient(
		context.Background(),
		dotenv.GetEnv("INTERNAL_KEYCLOAK_CLIENT_ID"),
		dotenv.GetEnv("INTERNAL_KEYCLOAK_HOST_SECRET"),
		dotenv.GetEnv("KEYCLOAK_REALM"),
	)

	if err != nil {
		fmt.Println("Error getting Internal Keycloak token: ", err)
		return ""
	}

	if token == nil {
		return ""
	}

	return token.AccessToken
}
