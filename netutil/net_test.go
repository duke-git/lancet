package netutil

import (
	"net"
	"net/http"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestGetInternalIp(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestGetInternalIp")

	internalIp := GetInternalIp()
	ip := net.ParseIP(internalIp)
	assert.IsNotNil(ip)
}

func TestGetRequestPublicIp(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
		t.Fail()
	}

	expected := "http://www.lancet.com?a=1&b=%5B2%5D"
	assert.Equal(expected, encodedUrl)
}

func TestIsPingConnected(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsPingConnected")

	// in github action env, this will fail
	// result1 := IsPingConnected("www.baidu.com")
	// assert.Equal(true, result1)

	result2 := IsPingConnected("www.!@#&&&.com")
	assert.Equal(false, result2)
}

func TestTelnetConnected(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestTelnetConnected")

	result1 := IsTelnetConnected("bing.com", "80")
	assert.Equal(true, result1)

	result2 := IsTelnetConnected("www.baidu.com", "123")
	assert.Equal(false, result2)
}

func TestBuildUrl(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBuildUrl")

	tests := []struct {
		scheme  string
		host    string
		path    string
		query   map[string][]string
		want    string
		wantErr bool
	}{
		{
			scheme:  "http",
			host:    "www.test.com",
			path:    "/path/subpath",
			query:   map[string][]string{"a": {"1"}, "b": {"2"}},
			want:    "http://www.test.com/path/subpath?a=1&b=2",
			wantErr: false,
		},
		{
			scheme:  "http",
			host:    "www.test.com",
			path:    "/simple-path",
			query:   map[string][]string{"a": {"1"}, "b": {"2"}},
			want:    "http://www.test.com/simple-path?a=1&b=2",
			wantErr: false,
		},
		{
			scheme:  "http",
			host:    "www.test.com",
			path:    "",
			query:   map[string][]string{"a": {"foo", "bar"}, "b": {"baz"}},
			want:    "http://www.test.com/?a=foo&a=bar&b=baz",
			wantErr: false,
		},
		{
			scheme:  "https",
			host:    "www.test. com",
			path:    "/path",
			query:   nil,
			want:    "",
			wantErr: true,
		},
		{
			scheme:  "https",
			host:    "www.test.com",
			path:    "/path with spaces",
			query:   nil,
			want:    "https://www.test.com/path%20with%20spaces",
			wantErr: false,
		},
		{
			scheme:  "https",
			host:    "my.api.edu.cn",
			path:    "/api",
			query:   nil,
			want:    "https://my.api.edu.cn/api",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		got, err := BuildUrl(tt.scheme, tt.host, tt.path, tt.query)
		assert.Equal(tt.want, got)
		assert.Equal(tt.wantErr, err != nil)
	}

}

func TestAddQueryParams(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAddQueryParams")

	tests := []struct {
		url     string
		query   map[string][]string
		want    string
		wantErr bool
	}{
		{
			url:     "http://www.test.com",
			query:   map[string][]string{"a": {"1"}, "b": {"2"}},
			want:    "http://www.test.com?a=1&b=2",
			wantErr: false,
		},
		{
			url:     "http://www.test.com",
			query:   map[string][]string{"a": {"foo", "bar"}, "b": {"baz"}},
			want:    "http://www.test.com?a=foo&a=bar&b=baz",
			wantErr: false,
		},
		{
			url:     "http://www.test.com",
			query:   map[string][]string{},
			want:    "http://www.test.com",
			wantErr: false,
		},
		{
			url:     "http://www.test.com",
			query:   map[string][]string{"a": {"$%"}},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		got, err := AddQueryParams(tt.url, tt.query)
		assert.Equal(tt.want, got)
		assert.Equal(tt.wantErr, err != nil)
	}
}
