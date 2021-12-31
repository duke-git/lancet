package netutil

import (
	"fmt"
	"net"
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestGetInternalIp(t *testing.T) {
	internalIp := GetInternalIp()
	ip := net.ParseIP(internalIp)
	if ip == nil {
		internal.LogFailedTestInfo(t, "GetInternalIp", "GetInternalIp", "", ip)
		t.FailNow()
	}
}

func TestGetPublicIpInfo(t *testing.T) {
	publicIpInfo, err := GetPublicIpInfo()
	if err != nil {
		t.FailNow()
	}
	fmt.Printf("public ip info is: %+v \n", *publicIpInfo)
}

func TestIsPublicIP(t *testing.T) {
	ips := []net.IP{
		net.ParseIP("127.0.0.1"),
		net.ParseIP("192.168.0.1"),
		net.ParseIP("10.91.210.131"),
		net.ParseIP("172.20.16.1"),
		net.ParseIP("36.112.24.10"),
	}

	expected := []bool{false, false, false, false, true}

	for i := 0; i < len(ips); i++ {
		res := IsPublicIP(ips[i])

		if res != expected[i] {
			internal.LogFailedTestInfo(t, "IsPublicIP", ips[i], expected[i], res)
			t.FailNow()
		}
	}
}
