// Package boilerplate provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package boilerplate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// ServeIndex request
	ServeIndex(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Healthcheck request
	Healthcheck(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateNewRoomWithBody request with any body
	CreateNewRoomWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateNewRoom(ctx context.Context, body CreateNewRoomJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetRoomInfo request
	GetRoomInfo(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetApi request
	GetApi(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ConnectRoom request
	ConnectRoom(ctx context.Context, id string, params *ConnectRoomParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ServeIndex(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewServeIndexRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Healthcheck(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewHealthcheckRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateNewRoomWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateNewRoomRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateNewRoom(ctx context.Context, body CreateNewRoomJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateNewRoomRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetRoomInfo(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRoomInfoRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetApi(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ConnectRoom(ctx context.Context, id string, params *ConnectRoomParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewConnectRoomRequest(c.Server, id, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewServeIndexRequest generates requests for ServeIndex
func NewServeIndexRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewHealthcheckRequest generates requests for Healthcheck
func NewHealthcheckRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/healthcheck")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateNewRoomRequest calls the generic CreateNewRoom builder with application/json body
func NewCreateNewRoomRequest(server string, body CreateNewRoomJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateNewRoomRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateNewRoomRequestWithBody generates requests for CreateNewRoom with any type of body
func NewCreateNewRoomRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/room")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetRoomInfoRequest generates requests for GetRoomInfo
func NewGetRoomInfoRequest(server string, id string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/room/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetApiRequest generates requests for GetApi
func NewGetApiRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/swagger")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewConnectRoomRequest generates requests for ConnectRoom
func NewConnectRoomRequest(server string, id string, params *ConnectRoomParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/ws/room/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Username != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "username", runtime.ParamLocationQuery, *params.Username); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// ServeIndexWithResponse request
	ServeIndexWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ServeIndexResponse, error)

	// HealthcheckWithResponse request
	HealthcheckWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*HealthcheckResponse, error)

	// CreateNewRoomWithBodyWithResponse request with any body
	CreateNewRoomWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateNewRoomResponse, error)

	CreateNewRoomWithResponse(ctx context.Context, body CreateNewRoomJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateNewRoomResponse, error)

	// GetRoomInfoWithResponse request
	GetRoomInfoWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetRoomInfoResponse, error)

	// GetApiWithResponse request
	GetApiWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiResponse, error)

	// ConnectRoomWithResponse request
	ConnectRoomWithResponse(ctx context.Context, id string, params *ConnectRoomParams, reqEditors ...RequestEditorFn) (*ConnectRoomResponse, error)
}

type ServeIndexResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r ServeIndexResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ServeIndexResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type HealthcheckResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r HealthcheckResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r HealthcheckResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateNewRoomResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *RoomInfo
	JSON500      *Error
}

// Status returns HTTPResponse.Status
func (r CreateNewRoomResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateNewRoomResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetRoomInfoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *RoomInfo
	JSON404      *Error
	JSON500      *Error
}

// Status returns HTTPResponse.Status
func (r GetRoomInfoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetRoomInfoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetApiResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *API
}

// Status returns HTTPResponse.Status
func (r GetApiResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetApiResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ConnectRoomResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON404      *Error
	JSON500      *Error
}

// Status returns HTTPResponse.Status
func (r ConnectRoomResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ConnectRoomResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ServeIndexWithResponse request returning *ServeIndexResponse
func (c *ClientWithResponses) ServeIndexWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ServeIndexResponse, error) {
	rsp, err := c.ServeIndex(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseServeIndexResponse(rsp)
}

// HealthcheckWithResponse request returning *HealthcheckResponse
func (c *ClientWithResponses) HealthcheckWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*HealthcheckResponse, error) {
	rsp, err := c.Healthcheck(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseHealthcheckResponse(rsp)
}

// CreateNewRoomWithBodyWithResponse request with arbitrary body returning *CreateNewRoomResponse
func (c *ClientWithResponses) CreateNewRoomWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateNewRoomResponse, error) {
	rsp, err := c.CreateNewRoomWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateNewRoomResponse(rsp)
}

func (c *ClientWithResponses) CreateNewRoomWithResponse(ctx context.Context, body CreateNewRoomJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateNewRoomResponse, error) {
	rsp, err := c.CreateNewRoom(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateNewRoomResponse(rsp)
}

// GetRoomInfoWithResponse request returning *GetRoomInfoResponse
func (c *ClientWithResponses) GetRoomInfoWithResponse(ctx context.Context, id string, reqEditors ...RequestEditorFn) (*GetRoomInfoResponse, error) {
	rsp, err := c.GetRoomInfo(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRoomInfoResponse(rsp)
}

// GetApiWithResponse request returning *GetApiResponse
func (c *ClientWithResponses) GetApiWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiResponse, error) {
	rsp, err := c.GetApi(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiResponse(rsp)
}

// ConnectRoomWithResponse request returning *ConnectRoomResponse
func (c *ClientWithResponses) ConnectRoomWithResponse(ctx context.Context, id string, params *ConnectRoomParams, reqEditors ...RequestEditorFn) (*ConnectRoomResponse, error) {
	rsp, err := c.ConnectRoom(ctx, id, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseConnectRoomResponse(rsp)
}

// ParseServeIndexResponse parses an HTTP response from a ServeIndexWithResponse call
func ParseServeIndexResponse(rsp *http.Response) (*ServeIndexResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ServeIndexResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseHealthcheckResponse parses an HTTP response from a HealthcheckWithResponse call
func ParseHealthcheckResponse(rsp *http.Response) (*HealthcheckResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &HealthcheckResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseCreateNewRoomResponse parses an HTTP response from a CreateNewRoomWithResponse call
func ParseCreateNewRoomResponse(rsp *http.Response) (*CreateNewRoomResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateNewRoomResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest RoomInfo
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetRoomInfoResponse parses an HTTP response from a GetRoomInfoWithResponse call
func ParseGetRoomInfoResponse(rsp *http.Response) (*GetRoomInfoResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetRoomInfoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest RoomInfo
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetApiResponse parses an HTTP response from a GetApiWithResponse call
func ParseGetApiResponse(rsp *http.Response) (*GetApiResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest API
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseConnectRoomResponse parses an HTTP response from a ConnectRoomWithResponse call
func ParseConnectRoomResponse(rsp *http.Response) (*ConnectRoomResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ConnectRoomResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
