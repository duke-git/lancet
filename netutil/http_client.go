package netutil

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/duke-git/lancet/v2/slice"
)

// HttpRequest struct is a composed http request
type HttpRequest struct {
	RawURL      string
	Method      string
	Headers     http.Header
	QueryParams url.Values
	FormData    url.Values
	Body        []byte
}

// HttpClientConfig contains some configurations for http client
type HttpClientConfig struct {
	SSLEnabled       bool
	TLSConfig        *tls.Config
	Compressed       bool
	HandshakeTimeout time.Duration
	ResponseTimeout  time.Duration
	Verbose          bool
}

// defaultHttpClientConfig defalut client config
var defaultHttpClientConfig = &HttpClientConfig{
	Compressed:       false,
	HandshakeTimeout: 20 * time.Second,
	ResponseTimeout:  40 * time.Second,
}

// HttpClient is used for sending http request
type HttpClient struct {
	*http.Client
	TLS     *tls.Config
	Request *http.Request
	Config  HttpClientConfig
}

// NewHttpClient make a HttpClient instance
func NewHttpClient() *HttpClient {
	client := &HttpClient{
		Client: &http.Client{
			Transport: &http.Transport{
				TLSHandshakeTimeout:   defaultHttpClientConfig.HandshakeTimeout,
				ResponseHeaderTimeout: defaultHttpClientConfig.ResponseTimeout,
				DisableCompression:    !defaultHttpClientConfig.Compressed,
			},
		},
		Config: *defaultHttpClientConfig,
	}

	return client
}

// NewHttpClientWithConfig make a HttpClient instance with pass config
func NewHttpClientWithConfig(config *HttpClientConfig) *HttpClient {
	if config == nil {
		config = defaultHttpClientConfig
	}

	client := &HttpClient{
		Client: &http.Client{
			Transport: &http.Transport{
				TLSHandshakeTimeout:   config.HandshakeTimeout,
				ResponseHeaderTimeout: config.ResponseTimeout,
				DisableCompression:    !config.Compressed,
			},
		},
		Config: *config,
	}

	if config.SSLEnabled {
		client.TLS = config.TLSConfig
	}

	return client
}

// SendRequest send http request
func (client *HttpClient) SendRequest(request *HttpRequest) (*http.Response, error) {
	err := validateRequest(request)
	if err != nil {
		return nil, err
	}

	rawUrl := request.RawURL

	req, err := http.NewRequest(request.Method, rawUrl, bytes.NewBuffer(request.Body))
	if err != nil {
		return nil, err
	}

	client.setTLS(rawUrl)
	client.setHeader(req, request.Headers)

	err = client.setQueryParam(req, rawUrl, request.QueryParams)
	if err != nil {
		return nil, err
	}

	if request.FormData != nil {
		client.setFormData(req, request.FormData)
	}

	client.Request = req

	resp, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DecodeResponse decode response into target object
func (client *HttpClient) DecodeResponse(resp *http.Response, target any) error {
	if resp == nil {
		return errors.New("invalid target param")
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

// setTLS set http client transport TLSClientConfig
func (client *HttpClient) setTLS(rawUrl string) {
	if strings.HasPrefix(rawUrl, "https") {
		if transport, ok := client.Client.Transport.(*http.Transport); ok {
			transport.TLSClientConfig = client.TLS
		}
	}
}

// setHeader set http rquest header
func (client *HttpClient) setHeader(req *http.Request, headers http.Header) {
	if headers == nil {
		headers = make(http.Header)
	}

	if _, ok := headers["Accept"]; !ok {
		headers["Accept"] = []string{"*/*"}
	}
	if _, ok := headers["Accept-Encoding"]; !ok && client.Config.Compressed {
		headers["Accept-Encoding"] = []string{"deflate, gzip"}
	}

	req.Header = headers
}

// setQueryParam set http request query string param
func (client *HttpClient) setQueryParam(req *http.Request, reqUrl string, queryParam url.Values) error {
	if queryParam != nil {
		if !strings.Contains(reqUrl, "?") {
			reqUrl = reqUrl + "?" + queryParam.Encode()
		} else {
			reqUrl = reqUrl + "&" + queryParam.Encode()
		}
		u, err := url.Parse(reqUrl)
		if err != nil {
			return err
		}
		req.URL = u
	}
	return nil
}

func (client *HttpClient) setFormData(req *http.Request, values url.Values) {
	formData := []byte(values.Encode())
	req.Body = ioutil.NopCloser(bytes.NewReader(formData))
	req.ContentLength = int64(len(formData))
}

// validateRequest check if a request has url, and valid method.
func validateRequest(req *HttpRequest) error {
	if req.RawURL == "" {
		return errors.New("invalid request url")
	}

	// common HTTP methods
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH",
		"HEAD", "CONNECT", "OPTIONS", "TRACE"}

	if !slice.Contain(methods, strings.ToUpper(req.Method)) {
		return errors.New("invalid request method")
	}

	return nil
}

// StructToUrlValues convert struct to url valuse,
// only convert the field which is exported and has `json` tag
func StructToUrlValues(targetStruct any) url.Values {
	rv := reflect.ValueOf(targetStruct)
	rt := reflect.TypeOf(targetStruct)

	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	if rt.Kind() != reflect.Struct {
		panic(fmt.Errorf("data type %T not support, shuld be struct or pointer to struct", targetStruct))
	}

	result := url.Values{}

	fieldNum := rt.NumField()
	pattern := `^[A-Z]`
	regex := regexp.MustCompile(pattern)
	for i := 0; i < fieldNum; i++ {
		name := rt.Field(i).Name
		tag := rt.Field(i).Tag.Get("json")
		if regex.MatchString(name) && tag != "" {
			result.Add(tag, fmt.Sprintf("%v", rv.Field(i).Interface()))
		}
	}

	return result
}
