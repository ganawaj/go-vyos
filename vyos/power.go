package vyos

import (
	"context"
)

// Response represents a response from the VyOS API.
type PowerService service

// ShowResponse represents a response from the VyOS API.
type PowerResponse struct {
	*RawResponse
}

// PowerOff sends a request to the VyOS API to power off the system.
func (s *PowerService) PowerOff(ctx context.Context) (*PowerResponse, *Response, error) {

	u := "/poweroff"

	// Create a new request.
	request := Request{
		OPMode: "poweroff",
		Path:   []string{"now"},
	}

	// Create the HTTP request.
	req, err := s.client.NewRequest(u, &request)
	if err != nil {
		return nil, nil, err
	}

	// Create the Response struct & send the request.
	v := new(PowerResponse)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Reboot sends a request to the VyOS API to reboot the system.
func (s *PowerService) Reboot(ctx context.Context) (*PowerResponse, *Response, error) {

	u := "/reboot"

	// Create a new request.
	request := Request{
		OPMode: "reboot",
		Path:   []string{"now"},
	}

	// Create the HTTP request.
	req, err := s.client.NewRequest(u, &request)
	if err != nil {
		return nil, nil, err
	}

	// Create the Response struct & send the request.
	v := new(PowerResponse)
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}


// PowerOff is a helper function to power off the VyOS instance from the client struct.
func (s *Client) PowerOff(ctx context.Context) (*PowerResponse, *Response, error) {
	return s.Power.PowerOff(ctx)
}

// Reboot is a helper function to reboot the VyOS instance from the client struct.
func (s *Client) Reboot(ctx context.Context) (*PowerResponse, *Response, error) {
	return s.Power.Reboot(ctx)
}