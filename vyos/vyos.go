package vyos

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"sync"
)

const (

	// defaultUserAgent is the default user agent used by the VyOS API client.
	defaultUserAgent        = "go-vyos"

	// OPMode constants
	OPModeShow       OPMode = "show"      // OPModeShow is the show operational mode.
	OPModeSet        OPMode = "set"       // OPModeSet is the set operational mode.
	OPModeComment    OPMode = "comment"   // OPModeComment is the comment operational mode.
	OPModeGenerate   OPMode = "generate"  // OPModeGenerate is the generate operational mode.
	OPModeConfigure  OPMode = "configure" // OPModeConfigure is the configure operational mode.
)

// Vyos represents a VyOS API client.
type Client struct {
	mu     sync.Mutex   // Mutex used to synchronize API requests.
	client *http.Client // HTTP client used to communicate with the API.

	BaseURL   string // Base URL for API requests.
	UserAgent string // User agent used when communicating with the API.

	Token string // Token used for authentication.

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the VyOS API.
	Show    *ShowService
	// Set     *SetService
	// Comment *CommentService


	Power *PowerService
	Image *ImageService
	Config *ConfigService
	
}

// Service represents a VyOS API service.
// These map to the different endponits in the VyOS API:
// /retrieve, /set, /delete, /comment, /commit, /discard, /rollback, /show, /run
type service struct {
	client *Client
}

// Request represents a request to the VyOS API.
type Response struct {
	*http.Response
}

type OPMode string // OPMode is a type for the different operational modes of the VyOS API.
type Path []string // Path is a type for the different paths of the VyOS API.

// Request represents a request to the VyOS API.
type Request struct {
	OPMode OPMode `json:"op,omitempty"`
	Path   Path   `json:"path,omitempty"`
}

// RawResponse represents a raw response from the VyOS API.
type RawResponse struct {
	Success bool        `json:"success,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// NewClient creates a new VyOS API client.
// If no HTTP client is provided, a new one is created.
func NewClient(httpClient *http.Client) *Client {

	// If no HTTP client is provided, create a new one.
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	// Create a new pointer to the http.Client struct.
	h := *httpClient

	// Create a new Vyos client.
	c := &Client{client: &h}
	c.init()

	return c
}

// Client returns a copy of the http.Client used by the VyOS API client.
func (c *Client) Client() *http.Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	clientCopy := *c.client
	return &clientCopy
}

// WithURL sets the base URL for the VyOS API client.
func (c *Client) WithURL(url string) *Client {
	newClient := c.copy()
	defer newClient.init()

	// Set the base URL for the API client.
	// TODO: Validate the URL. Parse() returns an error if the URL is invalid.
	newClient.BaseURL = url

	return newClient
}

// WithToken sets the token for the VyOS API client.
func (c *Client) WithToken(token string) *Client {
	newClient := c.copy()
	defer newClient.init()

	// Set the token for the API client.
	newClient.Token = token

	return newClient
}

// Insecure sets the InsecureSkipVerify field of the TLS configuration to true.
func (c *Client) Insecure() *Client {
	newClient := c.copy()
	defer newClient.init()

	// Type assertion to ensure transport is of type *http.Transport
	if t, ok := newClient.client.Transport.(*http.Transport); ok {
		t.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
		newClient.client.Transport = t
	} else {
		// Handle the case where the transport is not of type *http.Transport
		newClient.client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	return newClient
}

// init initializes the VyOS API client.
func (c *Client) init() {

	if c.client == nil {
		c.client = &http.Client{}
	}

	if c.UserAgent == "" {
		c.UserAgent = defaultUserAgent
	}

	c.common.client = c
	c.Show = (*ShowService)(&c.common)


	c.Power = (*PowerService)(&c.common)
	c.Image = (*ImageService)(&c.common)



}

// copy returns a copy of the VyOS API client.
func (c *Client) copy() *Client {

	// Create a new Vyos client.
	c.mu.Lock()
	defer c.mu.Unlock()

	return &Client{
		client:    &http.Client{},
		BaseURL:   c.BaseURL,
		Token:     c.Token,
		UserAgent: c.UserAgent,
	}
}

// NewRequest creates a new HTTP request for the VyOS API.
func (c *Client) NewRequest(urlStr string, request interface{}) (*http.Request, error) {

	// The method must always be POST. The VyOS API only supports POST requests.
	method := http.MethodPost

	// Resolve the URL.
	u := c.BaseURL + urlStr

	// Marshal the struct into JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	// Create a bytes.Buffer from the JSON data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("data", string(jsonData))
	writer.WriteField("key", c.Token)
	writer.Close()

	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}

	// Set the content type to multipart/form-data
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Set the content length header
	req.Header.Set("Content-Length", fmt.Sprintf("%d", body.Len()))

	// Set the user agent if it is provided.
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Do sends an API request and returns the API response.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {

	// Check if the context is nil. We always need a context - so they can be cancelled.
	if ctx == nil {
		return nil, ErrContextNil
	}

	if v == nil {
		return nil, ErrInterfaceNil
	}

	r, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	resp := &Response{Response: r}
	defer resp.Body.Close()

	// Decode the response body into the provided interface.
	errDecode := json.NewDecoder(resp.Body).Decode(v)
	if errDecode != nil {
		return nil, errDecode
	}

	return resp, nil

}
