package netutil

import (
	"net"
	"testing"
)

func TestIsTelnetConnectedIPv6(t *testing.T) {
	listener, err := net.Listen("tcp6", "[::1]:0")
	if err != nil {
		t.Skipf("IPv6 loopback unavailable: %v", err)
	}
	defer listener.Close()
	host, port, err := net.SplitHostPort(listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	if !IsTelnetConnected(host, port) {
		t.Error("could not connect to the IPv6 loopback listener")
	}
}
