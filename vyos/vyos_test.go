package vyos

import (
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
