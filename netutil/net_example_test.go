package netutil

import (
	"fmt"
	"net"
	"net/http"
)

func ExampleGetInternalIp() {
	internalIp := GetInternalIp()

	result := IsInternalIP(net.ParseIP(internalIp))
	fmt.Println(result)

	// Output:
	// true
}

func ExampleGetPublicIpInfo() {
	ipInfo, err := GetPublicIpInfo()
	if err != nil {
		return
	}

	result := IsPublicIP(net.ParseIP(ipInfo.Ip))
	fmt.Println(result)

	// Output:
	// true
}

func ExampleGetRequestPublicIp() {
	ip := "36.112.24.10"

	request := http.Request{
		Method: "GET",
		Header: http.Header{
			"X-Forwarded-For": {ip},
		},
	}
	publicIp := GetRequestPublicIp(&request)

	fmt.Println(publicIp)

	// Output:
	// 36.112.24.10
}

func ExampleIsInternalIP() {
	ip1 := IsInternalIP(net.ParseIP("127.0.0.1"))
	ip2 := IsInternalIP(net.ParseIP("192.168.0.1"))
	ip3 := IsInternalIP(net.ParseIP("36.112.24.10"))

	fmt.Println(ip1)
	fmt.Println(ip2)
	fmt.Println(ip3)

	// Output:
	// true
	// true
	// false
}

func ExampleIsPublicIP() {
	ip1 := IsPublicIP(net.ParseIP("127.0.0.1"))
	ip2 := IsPublicIP(net.ParseIP("192.168.0.1"))
	ip3 := IsPublicIP(net.ParseIP("36.112.24.10"))

	fmt.Println(ip1)
	fmt.Println(ip2)
	fmt.Println(ip3)

	// Output:
	// false
	// false
	// true
}

func ExampleEncodeUrl() {
	urlAddr := "http://www.lancet.com?a=1&b=[2]"

	encodedUrl, err := EncodeUrl(urlAddr)
	if err != nil {
		return
	}

	fmt.Println(encodedUrl)

	// Output:
	// http://www.lancet.com?a=1&b=%5B2%5D
}
