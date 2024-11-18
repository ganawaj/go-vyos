package vyos

import (
	"net/http"
	"testing"
)

// TestNewClient tests the NewClient function.
func TestNewClient(t *testing.T) {

	t.Parallel()
	c := NewClient(nil)

	if got, want := c.UserAgent, defaultUserAgent; got != want {
		t.Errorf("NewClient UserAgent is %v, want %v", got, want)
	}

	c2 := NewClient(nil)
	if c.client == c2.client {
		t.Error("NewClient returned same http.Clients, but they should differ")
	}

}

// TestClientCopy tests the copy method of the Client struct.
func TestClientCopy(t *testing.T) {

	t.Parallel()
	c := NewClient(nil)
	c2 := c.copy()

	if c.client == c2.client {
		t.Error("Client copy returned same http.Clients, but they should differ")
	}

}

// TestClientWithURL tests the WithURL method of the Client struct.
func TestClientWithURL(t *testing.T) {

	t.Parallel()
	c := NewClient(nil).WithURL("https://test.com")

	if got, want := c.BaseURL, "https://test.com"; got != want {
		t.Errorf("Client WithURL BaseURL is %v, want %v", got, want)
	}
}

func TestClientWithToken(t *testing.T) {

	t.Parallel()
	c := NewClient(nil).WithToken("test")

	if got, want := c.Token, "test"; got != want {
		t.Errorf("Client WithToken Token is %v, want %v", got, want)
	}
}

func TestClientInsecure(t *testing.T) {

	t.Parallel()
	c := NewClient(nil).Insecure()

	if c.client.Transport == nil {
		t.Error("Client Insecure Transport is nil, but it should be set")
	}

	if c.client.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify != true {
		t.Error("Client Insecure InsecureSkipVerify is false, but it should be true")
	}

}
