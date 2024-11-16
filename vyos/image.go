package vyos

import (
	"context"
)

// Response represents a response from the VyOS API.
type ImageService service

// ShowResponse represents a response from the VyOS API.
type ImageResponse struct {
	*RawResponse
}

type ImageAddRequest struct {
	OPMode OPMode `json:"op,omitempty"`
	URL    string `json:"url,omitempty"`
}

type ImageDeleteRequest struct {
	OPMode OPMode `json:"op,omitempty"`
	Name   string `json:"name,omitempty"`
}

// Add adds a new image from a url
func (s *ImageService) Add(ctx context.Context, url string) (*ImageResponse, *Response, error) {

	u := "/image"

	// Create a new request.
	request := ImageAddRequest{
		OPMode: "add",
		URL:   url,
	}

	// Create the HTTP request.
	req, err := s.client.NewRequest(u, &request)
	if err != nil {
		return nil, nil, err
	}

	s.client.mu.Lock()
	defer s.client.mu.Unlock()

	// Create the Response struct & send the request.
	v := new(ImageResponse)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Delete deletes an existing image
func (s *ImageService) Delete(ctx context.Context, name string) (*ImageResponse, *Response, error) {

	u := "/image"

	// Create a new request.
	request := ImageDeleteRequest{
		OPMode: "add",
		Name:   name,
	}

	// Create the HTTP request.
	req, err := s.client.NewRequest(u, &request)
	if err != nil {
		return nil, nil, err
	}

	s.client.mu.Lock()
	defer s.client.mu.Unlock()

	// Create the Response struct & send the request.
	v := new(ImageResponse)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
