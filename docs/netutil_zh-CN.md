# Netutil

netutil 网络包支持获取 ip 地址，发送 http 请求。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/netutil/net.go](https://github.com/duke-git/lancet/blob/main/netutil/net.go)

-   [https://github.com/duke-git/lancet/blob/main/netutil/http.go](https://github.com/duke-git/lancet/blob/main/netutil/http.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/netutil"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [ConvertMapToQueryString](#ConvertMapToQueryString)
-   [EncodeUrl](#EncodeUrl)
-   [GetInternalIp](#GetInternalIp)
-   [GetIps](#GetIps)
-   [GetMacAddrs](#GetMacAddrs)
-   [GetPublicIpInfo](#GetPublicIpInfo)
-   [GetRequestPublicIp](#GetRequestPublicIp)
-   [IsPublicIP](#IsPublicIP)
-   [IsInternalIP](#IsInternalIP)
-   [HttpRequest](#HttpRequest)
-   [HttpClient](#HttpClient)
-   [SendRequest](#SendRequest)
-   [DecodeResponse](#DecodeResponse)
-   [StructToUrlValues](#StructToUrlValues)
-   [HttpGet<sup>Deprecated</sup>](#HttpGet)
-   [HttpDelete<sup>Deprecated</sup>](#HttpDelete)
-   [HttpPost<sup>Deprecated</sup>](#HttpPost)
-   [HttpPut<sup>Deprecated</sup>](#HttpPut)
-   [HttpPatch<sup>Deprecated</sup>](#HttpPatch)
-   [ParseHttpResponse](#ParseHttpResponse)
-   [DownloadFile](#DownloadFile)
-   [UploadFile](#UploadFile)
-   [IsPingConnected](#IsPingConnected)
-   [IsTelnetConnected](#IsTelnetConnected)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="ConvertMapToQueryString">ConvertMapToQueryString</span>

<p>将map转换成http查询字符串.</p>

<b>函数签名:</b>

```go
func ConvertMapToQueryString(param map[string]any) string
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    var m = map[string]any{
        "c": 3,
        "a": 1,
        "b": 2,
    }
    qs := netutil.ConvertMapToQueryString(m)

    fmt.Println(qs)

    // Output:
    // a=1&b=2&c=3
}
```

### <span id="EncodeUrl">EncodeUrl</span>

<p>编码url query string的值</p>

<b>函数签名:</b>

```go
func EncodeUrl(urlStr string) (string, error)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    urlAddr := "http://www.lancet.com?a=1&b=[2]"
    encodedUrl, err := netutil.EncodeUrl(urlAddr)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(encodedUrl)

    // Output:
    // http://www.lancet.com?a=1&b=%5B2%5D
}
```

### <span id="GetInternalIp">GetInternalIp</span>

<p>获取内部ip</p>

<b>函数签名:</b>

```go
func GetInternalIp() string
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "net"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    internalIp := netutil.GetInternalIp()
    ip := net.ParseIP(internalIp)

    fmt.Println(ip)

    // Output:
    // 192.168.1.9
}
```

### <span id="GetIps">GetIps</span>

<p>获取ipv4地址列表</p>

<b>函数签名:</b>

```go
func GetIps() []string
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "net"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    ips := netutil.GetIps()
    fmt.Println(ips)

    // Output:
    // [192.168.1.9]
}
```

### <span id="GetMacAddrs">GetMacAddrs</span>

<p>获取mac地址列</p>

<b>函数签名:</b>

```go
func GetMacAddrs() []string {
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "net"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    macAddrs := netutil.GetMacAddrs()
    fmt.Println(macAddrs)

    // Output:
    // [18:31:bf:09:d1:56 76:ee:2a:e6:2e:0f 74:ee:2a:e6:2e:0f 74:ee:2a:e6:2e:0f]
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

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    publicIpInfo, err := netutil.GetPublicIpInfo()
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(publicIpInfo)
}
```

### <span id="GetRequestPublicIp">GetRequestPublicIp</span>

<p>获取http请求ip</p>

<b>函数签名:</b>

```go
func GetRequestPublicIp(req *http.Request) string
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    ip := "36.112.24.10"

    request := http.Request{
        Method: "GET",
        Header: http.Header{
            "X-Forwarded-For": {ip},
        },
    }
    publicIp := netutil.GetRequestPublicIp(&request)

    fmt.Println(publicIp)

    // Output:
    // 36.112.24.10
}
```

### <span id="IsPublicIP">IsPublicIP</span>

<p>判断ip是否是公共ip</p>

<b>函数签名:</b>

```go
func IsPublicIP(IP net.IP) bool
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "net"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    ip1 := netutil.IsPublicIP(net.ParseIP("127.0.0.1"))
    ip2 := netutil.IsPublicIP(net.ParseIP("192.168.0.1"))
    ip3 := netutil.IsPublicIP(net.ParseIP("36.112.24.10"))

    fmt.Println(ip1)
    fmt.Println(ip2)
    fmt.Println(ip3)

    // Output:
    // false
    // false
    // true
}
```

### <span id="IsInternalIP">IsInternalIP</span>

<p>判断ip是否是局域网ip.</p>

<b>函数签名:</b>

```go
func IsInternalIP(IP net.IP) bool
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "net"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    ip1 := netutil.IsInternalIP(net.ParseIP("127.0.0.1"))
    ip2 := netutil.IsInternalIP(net.ParseIP("192.168.0.1"))
    ip3 := netutil.IsInternalIP(net.ParseIP("36.112.24.10"))

    fmt.Println(ip1)
    fmt.Println(ip2)
    fmt.Println(ip3)

    // Output:
    // true
    // true
    // false
}
```

### <span id="HttpRequest">HttpRequest</span>

<p>HttpRequest用于抽象HTTP请求实体的结构</p>

<b>函数签名:</b>

```go
type HttpRequest struct {
    RawURL      string
    Method      string
    Headers     http.Header
    QueryParams url.Values
    FormData    url.Values
    Body        []byte
}
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "net"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    header := http.Header{}
    header.Add("Content-Type", "multipart/form-data")

    postData := url.Values{}
    postData.Add("userId", "1")
    postData.Add("title", "testItem")

    request := &netutil.HttpRequest{
        RawURL:   "https://jsonplaceholder.typicode.com/todos",
        Method:   "POST",
        Headers:  header,
        FormData: postData,
    }
}
```

### <span id="HttpClient">HttpClient</span>

<p>HttpClient是用于发送HTTP请求的结构体。它可以用一些配置参数或无配置实例化.</p>

<b>函数签名:</b>

```go
type HttpClient struct {
    *http.Client
    TLS     *tls.Config
    Request *http.Request
    Config  HttpClientConfig
}

type HttpClientConfig struct {
    SSLEnabled       bool
    TLSConfig        *tls.Config
    Compressed       bool
    HandshakeTimeout time.Duration
    ResponseTimeout  time.Duration
    Verbose          bool
}

func NewHttpClient() *HttpClient

func NewHttpClientWithConfig(config *HttpClientConfig) *HttpClient

```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "net"
    "time"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    httpClientCfg := netutil.HttpClientConfig{
        SSLEnabled: true,
        HandshakeTimeout:10 * time.Second
    }
    httpClient := netutil.NewHttpClientWithConfig(&httpClientCfg)
}
```

### <span id="SendRequest">SendRequest</span>

<p>HttpClient发送http请求</p>

<b>函数签名:</b>

```go
func (client *HttpClient) SendRequest(request *HttpRequest) (*http.Response, error)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "net"
    "time"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    request := &netutil.HttpRequest{
        RawURL: "https://jsonplaceholder.typicode.com/todos/1",
        Method: "GET",
    }

    httpClient := netutil.NewHttpClient()
    resp, err := httpClient.SendRequest(request)
    if err != nil || resp.StatusCode != 200 {
        return
    }

    type Todo struct {
        UserId    int    `json:"userId"`
        Id        int    `json:"id"`
        Title     string `json:"title"`
        Completed bool   `json:"completed"`
    }

    var todo Todo
    err = httpClient.DecodeResponse(resp, &todo)
    if err != nil {
        return
    }

    fmt.Println(todo.Id)

    // Output:
    // 1
}
```

### <span id="DecodeResponse">DecodeResponse</span>

<p>解析http响应体到目标结构体</p>

<b>函数签名:</b>

```go
func (client *HttpClient) DecodeResponse(resp *http.Response, target any) error
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "net"
    "time"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    request := &netutil.HttpRequest{
        RawURL: "https://jsonplaceholder.typicode.com/todos/1",
        Method: "GET",
    }

    httpClient := netutil.NewHttpClient()
    resp, err := httpClient.SendRequest(request)
    if err != nil || resp.StatusCode != 200 {
        return
    }

    type Todo struct {
        UserId    int    `json:"userId"`
        Id        int    `json:"id"`
        Title     string `json:"title"`
        Completed bool   `json:"completed"`
    }

    var todo Todo
    err = httpClient.DecodeResponse(resp, &todo)
    if err != nil {
        return
    }

    fmt.Println(todo.Id)

    // Output:
    // 1
}
```

### <span id="StructToUrlValues">StructToUrlValues</span>

<p>将结构体转为url values, 仅转化结构体导出字段并且包含`json` tag</p>

<b>函数签名:</b>

```go
func StructToUrlValues(targetStruct any) url.Values
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    type TodoQuery struct {
        Id     int    `json:"id"`
        UserId int    `json:"userId"`
        Name   string `json:"name,omitempty"`
        Status string
    }
    item := TodoQuery{
        Id:     1,
        UserId: 123,
        Name:   "test",
        Status: "completed",
    }
    queryValues := netutil.StructToUrlValues(item)

    fmt.Println(todoValues.Get("id"))
    fmt.Println(todoValues.Get("userId"))
    fmt.Println(todoValues.Get("name"))
    fmt.Println(todoValues.Get("status"))

    // Output:
    // 1
    // 123
    // test
    //
}
```

### <span id="HttpGet">HttpGet</span>

<p>发送http get请求。(已废弃:使用SendRequest)</p>

<b>函数签名:</b>

```go
// params[0] http请求header，类型必须是http.Header或者map[string]string
// params[1] http查询字符串，类型必须是url.Values或者map[string]string
// params[2] post请求体，类型必须是[]byte
// params[3] http client，类型必须是http.Client
func HttpGet(url string, params ...any) (*http.Response, error)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "github.com/duke-git/lancet/v2/netutil"
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

<p>发送http post请求。(已废弃:使用SendRequest)</p>

<b>函数签名:</b>

```go
// params[0] http请求header，类型必须是http.Header或者map[string]string
// params[1] http查询字符串，类型必须是url.Values或者map[string]string
// params[2] post请求体，类型必须是[]byte
// params[3] http client，类型必须是http.Client
func HttpPost(url string, params ...any) (*http.Response, error)
```

<b>示例:</b>

```go
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    url := "https://jsonplaceholder.typicode.com/todos"
    header := map[string]string{
        "Content-Type": "application/x-www-form-urlencoded",
    }
    
    postData := url.Values{}
	postData.Add("userId", "1")
	postData.Add("title", "TestToDo")

    resp, err := netutil.HttpPost(apiUrl, header, nil, postData)
	if err != nil {
		log.Fatal(err)
	}

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(body)
}
```

### <span id="HttpPut">HttpPut</span>

<p>发送http put请求。(已废弃:使用SendRequest)</p>

<b>函数签名:</b>

```go
// params[0] http请求header，类型必须是http.Header或者map[string]string
// params[1] http查询字符串，类型必须是url.Values或者map[string]string
// params[2] post请求体，类型必须是[]byte
// params[3] http client，类型必须是http.Client
func HttpPut(url string, params ...any) (*http.Response, error)
```

<b>示例:</b>

```go
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "github.com/duke-git/lancet/v2/netutil"
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

<p>发送http delete请求。(已废弃:使用SendRequest)</p>

<b>函数签名:</b>

```go
// params[0] http请求header，类型必须是http.Header或者map[string]string
// params[1] http查询字符串，类型必须是url.Values或者map[string]string
// params[2] post请求体，类型必须是[]byte
// params[3] http client，类型必须是http.Client
func HttpDelete(url string, params ...any) (*http.Response, error)
```

<b>示例:</b>

```go
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "github.com/duke-git/lancet/v2/netutil"
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

<p>发送http patch请求。(已废弃:使用SendRequest)</p>

<b>函数签名:</b>

```go
// params[0] http请求header，类型必须是http.Header或者map[string]string
// params[1] http查询字符串，类型必须是url.Values或者map[string]string
// params[2] post请求体，类型必须是[]byte
// params[3] http client，类型必须是http.Client
func HttpPatch(url string, params ...any) (*http.Response, error)
```

<b>示例:</b>

```go
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "github.com/duke-git/lancet/v2/netutil"
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
func ParseHttpResponse(resp *http.Response, obj any) error
```

<b>示例:</b>

```go
package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "github.com/duke-git/lancet/v2/netutil"
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

### <span id="DownloadFile">DownloadFile</span>

<p>从指定的server地址下载文件。</p>

<b>函数签名:</b>

```go
func DownloadFile(filepath string, url string) error
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    err := netutil.DownloadFile("./lancet_logo.jpg", "https://picx.zhimg.com/v2-fc82a4199749de9cfb71e32e54f489d3_720w.jpg?source=172ae18b")

    fmt.Println(err)
}
```

### <span id="UploadFile">UploadFile</span>

<p>将文件上传指定的server地址。</p>

<b>函数签名:</b>

```go
func UploadFile(filepath string, server string) (bool, error)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    ok, err := netutil.UploadFile("./a.jpg", "http://www.xxx.com/bucket/test")

    fmt.Println(ok)
    fmt.Println(err)
}
```

### <span id="IsPingConnected">IsPingConnected</span>

<p>检查能否ping通主机。</p>

<b>函数签名:</b>

```go
func IsPingConnected(host string) bool
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    result1 := netutil.IsPingConnected("www.baidu.com")
    result2 := netutil.IsPingConnected("www.!@#&&&.com")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="IsTelnetConnected">IsTelnetConnected</span>

<p>检查能否telnet到主机。</p>

<b>函数签名:</b>

```go
func IsTelnetConnected(host string, port string) bool
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/netutil"
)

func main() {
    result1 := netutil.IsTelnetConnected("www.baidu.com", "80")
    result2 := netutil.IsTelnetConnected("www.baidu.com", "123")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```
