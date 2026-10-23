package netutil

import (
	"net/http/httptest"
	"testing"
)

func TestGetRequestPublicIpInvalidHeaders(t *testing.T) {
	for _, test := range []struct{ forwarded, real, want string }{
		{"unknown, 8.8.8.8", "", "8.8.8.8"},
		{"unknown", "invalid", "1.1.1.1"},
		{"10.0.0.1", "8.8.4.4", "8.8.4.4"},
	} {
		req := httptest.NewRequest("GET", "http://example.com", nil)
		req.RemoteAddr = "1.1.1.1:1234"
		req.Header.Set("X-Forwarded-For", test.forwarded)
		req.Header.Set("X-Real-Ip", test.real)
		if got := GetRequestPublicIp(req); got != test.want {
			t.Errorf("got %q, want %q", got, test.want)
		}
	}
}
