package vyos

import (
	"context"
	"strings"
)

type ConfigureService service

type ConfigureResponse struct {
	*RawResponse
}


// Do sends a request to the VyOS API and returns the response.
func (s *ConfigureService) Set(ctx context.Context, path string) (*ConfigureResponse, *Response, error) {

	u := "/configure"

	// Create a new request.
	request := Request{
		OPMode: OPModeSet,
		Path:  strings.Split(path, " "),
	}

	// Create the HTTP request.
	req, err := s.client.NewRequest(u, &request)
	if err != nil {
		return nil, nil, err
	}

	s.client.mu.Lock()
	defer s.client.mu.Unlock()

	// Create the Response struct & send the request.
	v := new(ConfigureResponse)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil

}

func (s *ConfigureService) Delete(ctx context.Context, path string) (*ConfigureResponse, *Response, error) {

	u := "/configure"

	// Create a new request.
	request := Request{
		OPMode: "delete",
		Path: strings.Split(path, " "),
	}

	// Create the HTTP request.
	req, err := s.client.NewRequest(u, &request)
	if err != nil {
		return nil, nil, err
	}

	s.client.mu.Lock()
	defer s.client.mu.Unlock()

	// Create the Response struct & send the request.
	v := new(ConfigureResponse)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil

}

func (s *ConfigureService) Comment(ctx context.Context, path string) (*ConfigureResponse, *Response, error) {

	u := "/configure"

	// Create a new request.
	request := Request{
		OPMode: "comment",
		Path: strings.Split(path, " "),
	}

	// Create the HTTP request.
	req, err := s.client.NewRequest(u, &request)
	if err != nil {
		return nil, nil, err
	}

	s.client.mu.Lock()
	defer s.client.mu.Unlock()

	// Create the Response struct & send the request.
	v := new(ConfigureResponse)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil

}