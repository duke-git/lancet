package netutil

import (
	"net"
	"net/http"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestGetInternalIp(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetInternalIp")

	internalIp := GetInternalIp()
	ip := net.ParseIP(internalIp)
	assert.IsNotNil(ip)
}

func TestGetRequestPublicIp(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetPublicIpInfo")

	ip := "36.112.24.10"

	request := http.Request{
		Method: "GET",
		Header: http.Header{
			"X-Forwarded-For": {ip},
		},
	}
	publicIp := GetRequestPublicIp(&request)
	assert.Equal(publicIp, ip)

	request = http.Request{
		Method: "GET",
		Header: http.Header{
			"X-Real-Ip": {ip},
		},
	}
	publicIp = GetRequestPublicIp(&request)
	assert.Equal(publicIp, ip)
}

// func TestGetPublicIpInfo(t *testing.T) {
// 	assert := internal.NewAssert(t, "TestGetPublicIpInfo")

// 	publicIpInfo, err := GetPublicIpInfo()
// 	assert.IsNil(err)

// 	t.Logf("public ip info is: %+v \n", *publicIpInfo)
// }

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

func TestIsInternalIP(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsInternalIP")

	ips := []net.IP{
		net.ParseIP("127.0.0.1"),
		net.ParseIP("192.168.0.1"),
		net.ParseIP("10.91.210.131"),
		net.ParseIP("172.20.16.1"),
		net.ParseIP("36.112.24.10"),
	}

	expected := []bool{true, true, true, true, false}

	for i := 0; i < len(ips); i++ {
		actual := IsInternalIP(ips[i])
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

func TestEncodeUrl(t *testing.T) {
	assert := internal.NewAssert(t, "TestEncodeUrl")

	urlAddr := "http://www.lancet.com?a=1&b=[2]"
	encodedUrl, err := EncodeUrl(urlAddr)
	if err != nil {
		t.Log(err)
	}

	expected := "http://www.lancet.com?a=1&b=%5B2%5D"
	assert.Equal(expected, encodedUrl)
}

// func TestDownloadFile(t *testing.T) {
// 	assert := internal.NewAssert(t, "TestDownloadFile")

// 	err := DownloadFile("./lancet_logo.jpg", "https://picx.zhimg.com/v2-fc82a4199749de9cfb71e32e54f489d3_720w.jpg?source=172ae18b")
// 	assert.IsNil(err)
// }

func TestIsPingConnected(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsPingConnected")

	// in github action env, this will fail
	// result1 := IsPingConnected("bing.com")
	// assert.Equal(true, result1)

	result2 := IsPingConnected("www.!@#&&&.com")
	assert.Equal(false, result2)
}

func TestTelnetConnected(t *testing.T) {
	assert := internal.NewAssert(t, "TestTelnetConnected")

	result1 := IsTelnetConnected("bing.com", "80")
	assert.Equal(true, result1)

	result2 := IsTelnetConnected("www.baidu.com", "123")
	assert.Equal(false, result2)
}
