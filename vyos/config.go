package vyos

import (
	"context"
	"errors"
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

var (
	ErrMustLoadFromFile = errors.New("file must not be empty or nil")
)

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