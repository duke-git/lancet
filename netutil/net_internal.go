package netutil

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func doHttpRequest(method, reqUrl string, params ...any) (*http.Response, error) {
	if len(reqUrl) == 0 {
		return nil, errors.New("url should be specified")
	}

	req := &http.Request{
		Method:     method,
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}

	client := &http.Client{}
	err := setUrl(req, reqUrl)
	if err != nil {
		return nil, err
	}

	switch len(params) {
	case 1:
		err = setHeader(req, params[0])
		if err != nil {
			return nil, err
		}
	case 2:
		err := setHeaderAndQueryParam(req, reqUrl, params[0], params[1])
		if err != nil {
			return nil, err
		}
	case 3:
		err := setHeaderAndQueryAndBody(req, reqUrl, params[0], params[1], params[2])
		if err != nil {
			return nil, err
		}
	case 4:
		err := setHeaderAndQueryAndBody(req, reqUrl, params[0], params[1], params[2])
		if err != nil {
			return nil, err
		}
		client, err = getClient(params[3])
		if err != nil {
			return nil, err
		}
	}

	resp, e := client.Do(req)
	return resp, e
}

func setHeaderAndQueryParam(req *http.Request, reqUrl string, header, queryParam any) error {
	err := setHeader(req, header)
	if err != nil {
		return err
	}
	err = setQueryParam(req, reqUrl, queryParam)
	if err != nil {
		return err
	}
	return nil
}

func setHeaderAndQueryAndBody(req *http.Request, reqUrl string, header, queryParam, body any) error {
	if err := setHeader(req, header); err != nil {
		return err
	} else if err = setQueryParam(req, reqUrl, queryParam); err != nil {
		return err
	} else if err = setBodyByte(req, body); err != nil {
		return err
	}
	return nil
}

func setHeader(req *http.Request, header any) error {
	if header == nil {
		return nil
	}

	switch v := header.(type) {
	case map[string]string:
		for k := range v {
			req.Header.Add(k, v[k])
		}
	case http.Header:
		for k, vv := range v {
			for _, vvv := range vv {
				req.Header.Add(k, vvv)
			}
		}
	default:
		return errors.New("header params type should be http.Header or map[string]string")
	}

	if host := req.Header.Get("Host"); host != "" {
		req.Host = host
	}

	return nil
}

func setUrl(req *http.Request, reqUrl string) error {
	u, err := url.Parse(reqUrl)
	if err != nil {
		return err
	}
	req.URL = u
	return nil
}

func setQueryParam(req *http.Request, reqUrl string, queryParam any) error {
	if queryParam == nil {
		return nil
	}

	var values url.Values
	switch v := queryParam.(type) {
	case map[string]string:
		values = url.Values{}
		for k := range v {
			values.Set(k, v[k])
		}
	case url.Values:
		values = v
	default:
		return errors.New("query string params type should be url.Values or map[string]string")
	}

	// set url
	if values != nil {
		if !strings.Contains(reqUrl, "?") {
			reqUrl = reqUrl + "?" + values.Encode()
		} else {
			reqUrl = reqUrl + "&" + values.Encode()
		}
	}
	u, err := url.Parse(reqUrl)
	if err != nil {
		return err
	}
	req.URL = u

	return nil
}

func setBodyByte(req *http.Request, body any) error {
	if body == nil {
		return nil
	}
	var bodyReader *bytes.Reader
	switch b := body.(type) {
	case io.Reader:
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, b); err != nil {
			return err
		}
		req.Body = ioutil.NopCloser(buf)
		req.ContentLength = int64(buf.Len())
	case []byte:
		bodyReader = bytes.NewReader(b)
		req.Body = ioutil.NopCloser(bodyReader)
		req.ContentLength = int64(bodyReader.Len())
	case map[string]interface{}:
		values := url.Values{}
		for k := range b {
			values.Set(k, fmt.Sprintf("%v", b[k]))
		}
		bodyReader = bytes.NewReader([]byte(values.Encode()))
		req.Body = ioutil.NopCloser(bodyReader)
		req.ContentLength = int64(bodyReader.Len())
	case url.Values:
		bodyReader = bytes.NewReader([]byte(b.Encode()))
		req.Body = ioutil.NopCloser(bodyReader)
		req.ContentLength = int64(bodyReader.Len())
	default:
		return fmt.Errorf("invalid body type: %T", b)
	}
	return nil
}

func getClient(client any) (*http.Client, error) {
	c := http.Client{}
	if client != nil {
		switch v := client.(type) {
		case http.Client:
			c = v
		default:
			return nil, errors.New("client type should be http.Client")
		}
	}

	return &c, nil
}
