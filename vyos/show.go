package vyos

import (
	"context"
	"strings"
)

// Response represents a response from the VyOS API.
type ShowService service

// ShowResponse represents a response from the VyOS API.
type ShowResponse struct {
	*RawResponse
}

// Do sends a request to the VyOS API and returns the response.
func (s *ShowService) Do(ctx context.Context, path string) (*ShowResponse, *Response, error) {

	u := "/show"
	p := strings.Split(path, " ")

	if p == nil {
		return nil, nil, ErrEmptyPath
	}

	// Create a new request.
	request := Request{
		OPMode: OPModeShow,
		Path:   p,
	}

	// Create the HTTP request.
	req, err := s.client.NewRequest(u, &request)
	if err != nil {
		return nil, nil, err
	}

	// Create the Response struct & send the request.
	v := new(ShowResponse)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
