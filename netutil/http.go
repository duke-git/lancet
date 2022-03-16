// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.

// Package netutil implements some basic functions to send http request and get ip info.
// Note:
// HttpGet, HttpPost, HttpDelete, HttpPut, HttpPatch, function param `url` is required.
// HttpGet, HttpPost, HttpDelete, HttpPut, HttpPatch, function param `params` is variable, the order is:
// params[0] is header which type should be http.Header or map[string]string,
// params[1] is query param which type should be url.Values or map[string]any,
// params[2] is post body which type should be []byte.
// params[3] is http client which type should be http.Client.
package netutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

//HttpGet send get http request
func HttpGet(url string, params ...any) (*http.Response, error) {
	return doHttpRequest(http.MethodGet, url, params...)
}

//HttpPost send post http request
func HttpPost(url string, params ...any) (*http.Response, error) {
	return doHttpRequest(http.MethodPost, url, params...)
}

//HttpPut send put http request
func HttpPut(url string, params ...any) (*http.Response, error) {
	return doHttpRequest(http.MethodPut, url, params...)
}

//HttpDelete send delete http request
func HttpDelete(url string, params ...any) (*http.Response, error) {
	return doHttpRequest(http.MethodDelete, url, params...)
}

// HttpPatch send patch http request
func HttpPatch(url string, params ...any) (*http.Response, error) {
	return doHttpRequest(http.MethodPatch, url, params...)
}

// ParseHttpResponse decode http response to specified interface
func ParseHttpResponse(resp *http.Response, obj any) error {
	if resp == nil {
		return errors.New("InvalidResp")
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(obj)
}

// ConvertMapToQueryString convert map to sorted url query string
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
