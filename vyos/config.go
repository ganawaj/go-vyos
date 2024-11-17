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

type RetriveConfigRequest struct {
	OPMode OPMode `json:"op,omitempty"`
	Path   Path `json:"path"`
}


var (
	ErrMustLoadFromFile = errors.New("file must not be empty or nil")
)

// Get retrives the configuration from the VyOS API. If path is empty, it will return the entire configuration.
// Note: the `Data` field of `ConfigResponse` struct will not encode the data to any struct.
func (s *ConfigService) Get(ctx context.Context, path string) (*ConfigResponse, *Response, error) {

	u := "/retrieve"

	var p []string

	if path != "" {
		p = strings.Split(path, " ")
	}

	// Create a new request.
	request := RetriveConfigRequest{
		OPMode: "showConfig",
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