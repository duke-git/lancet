package netutil

import (
	"net"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestGetInternalIp(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetInternalIp")

	internalIp := GetInternalIp()
	ip := net.ParseIP(internalIp)
	assert.IsNotNil(ip)
}

func TestGetPublicIpInfo(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetPublicIpInfo")

	publicIpInfo, err := GetPublicIpInfo()
	assert.IsNil(err)

	t.Logf("public ip info is: %+v \n", *publicIpInfo)
}

func TestIsPublicIP(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsPublicIP")

	ips := []net.IP{
		net.ParseIP("127.0.0.1"),
		net.ParseIP("192.168.0.1"),
		net.ParseIP("10.91.210.131"),
		net.ParseIP("172.20.16.1"),
		net.ParseIP("36.112.24.10"),
	}

	expected := []bool{false, false, false, false, true}

	for i := 0; i < len(ips); i++ {
		actual := IsPublicIP(ips[i])
		assert.Equal(expected[i], actual)
	}
}

func TestGetIps(t *testing.T) {
	ips := GetIps()
	t.Log(ips)
}

func TestGetMacAddrs(t *testing.T) {
	macAddrs := GetMacAddrs()
	t.Log(macAddrs)
}
