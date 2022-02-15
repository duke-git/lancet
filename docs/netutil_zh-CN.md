# Netutil
netutil网络包支持获取ip地址，发送http请求。

<div STYLE="page-break-after: always;"></div>

## 源码:

[https://github.com/duke-git/lancet/blob/main/netutil/net.go](https://github.com/duke-git/lancet/blob/main/netutil/net.go)

[https://github.com/duke-git/lancet/blob/main/netutil/http.go](https://github.com/duke-git/lancet/blob/main/netutil/http.go)

<div STYLE="page-break-after: always;"></div>

## 用法:
```go
import (
    "github.com/duke-git/lancet/netutil"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录
- [ConvertMapToQueryString](#ConvertMapToQueryString)
- [GetInternalIp](#GetInternalIp)
- [GetIps](#GetIps)
- [GetMacAddrs](#GetMacAddrs)
- [GetPublicIpInfo](#GetPublicIpInfo)
- [IsPublicIP](#IsPublicIP)
- [HttpGet](#HttpGet)
- [HttpDelete](#HttpDelete)
- [HttpPost](#HttpPost)
- [HttpPut](#HttpPut)

- [HttpPatch](#HttpPatch)
- [ParseHttpResponse](#ParseHttpResponse)

<div STYLE="page-break-after: always;"></div>

## 文档


### <span id="ConvertMapToQueryString">ConvertMapToQueryString</span>
<p>将map转换成http查询字符串.</p>

<b>函数签名:</b>

```go
func ConvertMapToQueryString(param map[string]interface{}) string
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/netutil"
)

func main() {
	var m = map[string]interface{}{
		"c": 3,
		"a": 1,
		"b": 2,
	}
	qs := netutil.ConvertMapToQueryString(m)

	fmt.Println(qs) //a=1&b=2&c=3
}
```



### <span id="GetInternalIp">GetInternalIp</span>
<p>获取内部ip</p>

<b>函数签名:</b>

```go
func GetInternalIp() string
```
<b>例子:</b>

```go
package main

import (
    "fmt"
	"net"
    "github.com/duke-git/lancet/netutil"
)

func main() {
	internalIp := netutil.GetInternalIp()
	ip := net.ParseIP(internalIp)

	fmt.Println(ip) //192.168.1.9
}
```


### <span id="GetIps">GetIps</span>
<p>获取ipv4地址列表</p>

<b>函数签名:</b>

```go
func GetIps() []string
```
<b>例子:</b>

```go
package main

import (
    "fmt"
	"net"
    "github.com/duke-git/lancet/netutil"
)

func main() {
	ips := netutil.GetIps()
	fmt.Println(ips) //[192.168.1.9]
}
```



### <span id="GetMacAddrs">GetMacAddrs</span>
<p>获取mac地址列</p>

<b>函数签名:</b>

```go
func GetMacAddrs() []string {
```
<b>例子:</b>

```go
package main

import (
    "fmt"
	"net"
    "github.com/duke-git/lancet/netutil"
)

func main() {
	addrs := netutil.GetMacAddrs()
	fmt.Println(addrs)
}
```



### <span id="GetPublicIpInfo">GetPublicIpInfo</span>
<p>获取公网ip信息</p>

<b>函数签名:</b>

```go
func GetPublicIpInfo() (*PublicIpInfo, error)
type PublicIpInfo struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Ip          string  `json:"query"`
}
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/netutil"
)

func main() {
	publicIpInfo, err := netutil.GetPublicIpInfo()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(publicIpInfo)
}
```



### <span id="IsPublicIP">IsPublicIP</span>
<p>判断ip是否是公共ip</p>

<b>函数签名:</b>

```go
func IsPublicIP(IP net.IP) bool
```
<b>例子:</b>

```go
package main

import (
    "fmt"
	"net"
    "github.com/duke-git/lancet/netutil"
)

func main() {
	ip1 := net.ParseIP("192.168.0.1")
	ip2 := net.ParseIP("36.112.24.10")

	fmt.Println(netutil.IsPublicIP(ip1)) //false
	fmt.Println(netutil.IsPublicIP(ip2)) //true
}
```




### <span id="HttpGet">HttpGet</span>
<p>发送http get请求</p>

<b>函数签名:</b>

```go
// params[0] http请求header，类型必须是http.Header或者map[string]string
// params[1] http查询字符串，类型必须是url.Values或者map[string]interface{}
// params[2] post请求体，类型必须是[]byte
// params[3] http client，类型必须是http.Client
func HttpGet(url string, params ...interface{}) (*http.Response, error)
```
<b>例子:</b>

```go
package main

import (
    "fmt"
	"io/ioutil"
	"log"
    "github.com/duke-git/lancet/netutil"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	header := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := netutil.HttpGet(url, header)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
}
```



### <span id="HttpPost">HttpPost</span>
<p>发送http post请求</p>

<b>函数签名:</b>

```go
// params[0] http请求header，类型必须是http.Header或者map[string]string
// params[1] http查询字符串，类型必须是url.Values或者map[string]interface{}
// params[2] post请求体，类型必须是[]byte
// params[3] http client，类型必须是http.Client
func HttpPost(url string, params ...interface{}) (*http.Response, error)
```
<b>例子:</b>

```go
package main

import (
	"encoding/json"
    "fmt"
	"io/ioutil"
	"log"
    "github.com/duke-git/lancet/netutil"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos"
	header := map[string]string{
		"Content-Type": "application/json",
	}
	type Todo struct {
		UserId int    `json:"userId"`
		Title  string `json:"title"`
	}
	todo := Todo{1, "TestAddToDo"}
	bodyParams, _ := json.Marshal(todo)

	resp, err := netutil.HttpPost(url, header, nil, bodyParams)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
}
```



### <span id="HttpPut">HttpPut</span>
<p>发送http put请求</p>

<b>函数签名:</b>

```go
// params[0] http请求header，类型必须是http.Header或者map[string]string
// params[1] http查询字符串，类型必须是url.Values或者map[string]interface{}
// params[2] post请求体，类型必须是[]byte
// params[3] http client，类型必须是http.Client
func HttpPut(url string, params ...interface{}) (*http.Response, error)
```
<b>Example:</b>

```go
package main

import (
	"encoding/json"
    "fmt"
	"io/ioutil"
	"log"
    "github.com/duke-git/lancet/netutil"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	header := map[string]string{
		"Content-Type": "application/json",
	}
	type Todo struct {
		Id     int    `json:"id"`
		UserId int    `json:"userId"`
		Title  string `json:"title"`
	}
	todo := Todo{1, 1, "TestPutToDo"}
	bodyParams, _ := json.Marshal(todo)

	resp, err := netutil.HttpPut(url, header, nil, bodyParams)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
}
```



### <span id="HttpDelete">HttpDelete</span>
<p>发送http delete请求</p>

<b>函数签名:</b>

```go
// params[0] http请求header，类型必须是http.Header或者map[string]string
// params[1] http查询字符串，类型必须是url.Values或者map[string]interface{}
// params[2] post请求体，类型必须是[]byte
// params[3] http client，类型必须是http.Client
func HttpDelete(url string, params ...interface{}) (*http.Response, error)
```
<b>例子:</b>

```go
package main

import (
	"encoding/json"
    "fmt"
	"io/ioutil"
	"log"
    "github.com/duke-git/lancet/netutil"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	resp, err := netutil.HttpDelete(url)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
}
```



### <span id="HttpPatch">HttpPatch</span>
<p>发送http patch请求</p>

<b>函数签名:</b>

```go
// params[0] http请求header，类型必须是http.Header或者map[string]string
// params[1] http查询字符串，类型必须是url.Values或者map[string]interface{}
// params[2] post请求体，类型必须是[]byte
// params[3] http client，类型必须是http.Client
func HttpPatch(url string, params ...interface{}) (*http.Response, error)
```
<b>例子:</b>

```go
package main

import (
	"encoding/json"
    "fmt"
	"io/ioutil"
	"log"
    "github.com/duke-git/lancet/netutil"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	header := map[string]string{
		"Content-Type": "application/json",
	}
	type Todo struct {
		Id     int    `json:"id"`
		UserId int    `json:"userId"`
		Title  string `json:"title"`
	}
	todo := Todo{1, 1, "TestPatchToDo"}
	bodyParams, _ := json.Marshal(todo)

	resp, err := netutil.HttpPatch(url, header, nil, bodyParams)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
}
```



### <span id="ParseHttpResponse">ParseHttpResponse</span>
<p>将http请求响应解码成特定struct值</p>

<b>函数签名:</b>

```go
func ParseHttpResponse(resp *http.Response, obj interface{}) error
```
<b>例子:</b>

```go
package main

import (
	"encoding/json"
    "fmt"
	"io/ioutil"
	"log"
    "github.com/duke-git/lancet/netutil"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	header := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := netutil.HttpGet(url, header)
	if err != nil {
		log.Fatal(err)
	}

	type Todo struct {
		Id        int    `json:"id"`
		UserId    int    `json:"userId"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	toDoResp := &Todo{}
	err = netutil.ParseHttpResponse(resp, toDoResp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(toDoResp)
}
```

