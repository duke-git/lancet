// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.

// Package netutil implements some basic functions to send http request and get ip info.
// Note:
// HttpGet, HttpPost, HttpDelete, HttpPut, HttpPatch, function param `url` is required.
// HttpGet, HttpPost, HttpDelete, HttpPut, HttpPatch, function param `params` is variable, the order is:
// params[0] is header which type should be http.Header or map[string]string,
// params[1] is query string param which type should be url.Values or map[string]string, when content-type header is
// multipart/form-data or application/x-www-form-urlencoded
// params[2] is post body which type should be []byte.
// params[3] is http client which type should be http.Client.
package netutil

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/slice"
)

// HttpGet send get http request.
func HttpGet(url string, params ...any) (*http.Response, error) {
	return doHttpRequest(http.MethodGet, url, params...)
}

// HttpPost send post http request.
func HttpPost(url string, params ...any) (*http.Response, error) {
	return doHttpRequest(http.MethodPost, url, params...)
}

// HttpPut send put http request.
func HttpPut(url string, params ...any) (*http.Response, error) {
	return doHttpRequest(http.MethodPut, url, params...)
}

// HttpDelete send delete http request.
func HttpDelete(url string, params ...any) (*http.Response, error) {
	return doHttpRequest(http.MethodDelete, url, params...)
}

// HttpPatch send patch http request.
func HttpPatch(url string, params ...any) (*http.Response, error) {
	return doHttpRequest(http.MethodPatch, url, params...)
}

// ParseHttpResponse decode http response to specified interface.
func ParseHttpResponse(resp *http.Response, obj any) error {
	if resp == nil {
		return errors.New("InvalidResp")
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(obj)
}

// ConvertMapToQueryString convert map to sorted url query string.
// Play: https://go.dev/play/p/jnNt_qoSnRi
func ConvertMapToQueryString(param map[string]any) string {
	if param == nil {
		return ""
	}
	var keys []string
	for key := range param {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var build strings.Builder
	for i, v := range keys {
		build.WriteString(v)
		build.WriteString("=")
		build.WriteString(fmt.Sprintf("%v", param[v]))
		if i != len(keys)-1 {
			build.WriteString("&")
		}
	}
	return build.String()
}

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

// SendRequest send http request.
// Play: https://go.dev/play/p/jUSgynekH7G
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

// DecodeResponse decode response into target object.
// Play: https://go.dev/play/p/jUSgynekH7G
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

// setHeader set http request header
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
	req.Body = io.NopCloser(bytes.NewReader(formData))
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
// only convert the field which is exported and has `json` tag.
// Play: https://go.dev/play/p/pFqMkM40w9z
func StructToUrlValues(targetStruct any) (url.Values, error) {
	result := url.Values{}
	s, err := convertor.StructToMap(targetStruct)
	if err != nil {
		return nil, err
	}
	for k, v := range s {
		result.Add(k, fmt.Sprintf("%v", v))
	}

	return result, nil
}
