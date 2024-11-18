package vyos

import (
	"context"
	"errors"
	"strings"
)

// Response represents a response from the VyOS API.
type ConfigService service

// ShowResponse represents a response from the VyOS API.
type ConfigResponse struct {
	*RawResponse
}

type ConfigRequest struct {
	OPMode OPMode `json:"op,omitempty"`
	File   string `json:"file,omitempty"`
}

// type RetriveConfigRequest struct {
// 	OPMode OPMode `json:"op,omitempty"`
// 	Path   Path `json:"path"`
// }

type RetrieveOptions struct {
	MultiValue bool
}

var (
	ErrMustLoadFromFile = errors.New("file must not be empty or nil")
)

// endpoint: /retrive
// method: POST

// Options
// op: showConfig, path: ""- To get the whole configuration
// op: showConfig, path: []string{} - To only get a part of the configuration
// op: returnValues, path: []string{} &ShowOptions{multivalue: true} - To get the multi-valued node
// op: exists, path: []string{} - To check existence of a configuration path

// Get retrives the configuration from the VyOS API. If path is empty, it will return the entire configuration.
// Handles:
// - op: showConfig, path: ""
// - op: showConfig, path: []string{}
// - op: returnValues, path: []string{} &ShowOptions{multivalue: true}
func (s *ConfigService) Get(ctx context.Context, path string, options *RetrieveOptions) (*ConfigResponse, *Response, error) {

	u := "/retrieve"

	var p []string
	op := "showConfig"

	if path != "" {
		p = strings.Split(path, " ")
	}

	// if the options contain MultiValue, then set the op to returnValues
	if options != nil && options.MultiValue {
		op = "returnValues"
	}

	// Create a new request.
	request := Request{
		OPMode: OPMode(op),
		Path: p,
	}

	// Create the HTTP request.
	req, err := s.client.NewRequest(u, &request)
	if err != nil {
		return nil, nil, err
	}

	// Create the Response struct & send the request.
	v := new(ConfigResponse)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Exists checks if a configuration path exists in the VyOS API.
// Handles:
// - op: exists, path: []string{}
func (s *ConfigService) Exists(ctx context.Context, path string) (*ConfigResponse, *Response, error) {

	u := "/retrieve"

	var p []string

	if path != "" {
		p = strings.Split(path, " ")
	}

	// Create a new request.
	request := Request{
		OPMode: "exists",
		Path: p,
	}

	// Create the HTTP request.
	req, err := s.client.NewRequest(u, &request)
	if err != nil {
		return nil, nil, err
	}

	// Create the Response struct & send the request.
	v := new(ConfigResponse)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// endpoint: /configure
// method: POST

// Options
// op: set, path: []string{} - To set a configuration path
// NOTE: Special case is a list of Requests. []Request
// op: delete, path: []string{} - To delete a configuration path
// op: comment, path: []string{} - To comment a configuration path


// Set sets a configuration path in the VyOS API.
// NOTE: This does not send multiple requests to the API. It sends a single request with multiple paths.
func (s *ConfigService) Set(ctx context.Context, path ...string) (*ConfigResponse, *Response, error) {

	u := "/configure"

	var r []Request

	if path == nil {
		return nil, nil, ErrEmptyPath
	}

	for _, p := range path {

		if p == "" {
			return nil, nil, ErrEmptyPath
		}

		p := strings.Split(p, " ")
		if p == nil {
			return nil, nil, ErrEmptyPath
		}

		// Create a new request.
		request := Request{
			OPMode: OPMode("set"),
			Path: p,
		}

		r = append(r, request)
	}

	// Create the HTTP request.
	req, err := s.client.NewRequest(u, &r)
	if err != nil {
		return nil, nil, err
	}

	s.client.mu.Lock()
	defer s.client.mu.Unlock()

	// Create the Response struct & send the request.
	v := new(ConfigResponse)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil

}



// Delete deletes a configuration path in the VyOS API.
func (s *ConfigService) Delete(ctx context.Context, path string) (*ConfigResponse, *Response, error) {

	u := "/configure"

	if path == "" {
		return nil, nil, ErrEmptyPath
	}

	p := strings.Split(path, " ")
	if p == nil {
		return nil, nil, ErrEmptyPath
	}

	// Create a new request.
	request := Request{
		OPMode: OPMode("delete"),
		Path:  p,
	}

	// Create the HTTP request.
	req, err := s.client.NewRequest(u, &request)
	if err != nil {
		return nil, nil, err
	}

	s.client.mu.Lock()
	defer s.client.mu.Unlock()

	// Create the Response struct & send the request.
	v := new(ConfigResponse)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil

}

func (s *ConfigService) Comment(ctx context.Context, path string) (*ConfigResponse, *Response, error) {

	u := "/configure"

	if path == "" {
		return nil, nil, ErrEmptyPath
	}

	p := strings.Split(path, " ")
	if p == nil {
		return nil, nil, ErrEmptyPath
	}

	// Create a new request.
	request := Request{
		OPMode: OPMode("comment"),
		Path:  p,
	}

	// Create the HTTP request.
	req, err := s.client.NewRequest(u, &request)
	if err != nil {
		return nil, nil, err
	}

	s.client.mu.Lock()
	defer s.client.mu.Unlock()

	// Create the Response struct & send the request.
	v := new(ConfigResponse)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil

}

// Add adds a new image from a url
func (s *ConfigService) Save(ctx context.Context, file string) (*ConfigResponse, *Response, error) {

	u := "/config-file"

	// Create a new request.
	request := ConfigRequest{
		OPMode: "save",
		File: file,
	}

	// Create the HTTP request.
	req, err := s.client.NewRequest(u, &request)
	if err != nil {
		return nil, nil, err
	}

	s.client.mu.Lock()
	defer s.client.mu.Unlock()

	// Create the Response struct & send the request.
	v := new(ConfigResponse)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Add adds a new image from a url
func (s *ConfigService) Load(ctx context.Context, file string) (*ConfigResponse, *Response, error) {

	u := "/config-file"

	if file == "" {
		return nil, nil, ErrMustLoadFromFile
	}

	// Create a new request.
	request := ConfigRequest{
		OPMode: "load",
		File: file,
	}

	// Create the HTTP request.
	req, err := s.client.NewRequest(u, &request)
	if err != nil {
		return nil, nil, err
	}

	s.client.mu.Lock()
	defer s.client.mu.Unlock()

	// Create the Response struct & send the request.
	v := new(ConfigResponse)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}