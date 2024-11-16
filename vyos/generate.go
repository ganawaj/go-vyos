package vyos

import (
	"context"
	"strings"
)

type GenerateService service

type GenerateResponse struct {
	*RawResponse
}

// Do sends a request to the VyOS API and returns the response.
func (s *GenerateService) Do(ctx context.Context, path string) (*GenerateResponse, *Response, error) {

	u := "/generate"
	p := strings.Split(path, " ")

	if p == nil {
		return nil, nil, ErrEmptyPath
	}

	// Create a new request.
	request := Request{
		OPMode: OPModeGenerate,
		Path:   p,
	}

	// Create the HTTP request.
	req, err := s.client.NewRequest(u, &request)
	if err != nil {
		return nil, nil, err
	}

	// Create the Response struct & send the request.
	v := new(GenerateResponse)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}