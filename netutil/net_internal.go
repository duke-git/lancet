package netutil

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func doHttpRequest(method, reqUrl string, params ...interface{}) (*http.Response, error) {
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

func setHeaderAndQueryParam(req *http.Request, reqUrl string, header, queryParam interface{}) error {
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

func setHeaderAndQueryAndBody(req *http.Request, reqUrl string, header, queryParam, body interface{}) error {
	err := setHeader(req, header)
	if err != nil {
		return err
	}
	err = setQueryParam(req, reqUrl, queryParam)
	if err != nil {
		return err
	}
	if req.Header.Get("Content-Type") == "multipart/form-data" || req.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		formData := queryParam.(url.Values)
		err = setBodyByte(req, formData.Encode())
	} else {
		err = setBodyByte(req, body)
	}
	if err != nil {
		return err
	}
	return nil
}

func setHeader(req *http.Request, header interface{}) error {
	if header != nil {
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

func setQueryParam(req *http.Request, reqUrl string, queryParam interface{}) error {
	var values url.Values
	if queryParam != nil {
		switch v := queryParam.(type) {
		case map[string]interface{}:
			values = url.Values{}
			for k := range v {
				values.Set(k, fmt.Sprintf("%v", v[k]))
			}
		case url.Values:
			values = v
		default:
			return errors.New("query params type should be url.Values or map[string]interface{}")
		}
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

func setBodyByte(req *http.Request, body interface{}) error {
	if body != nil {
		switch b := body.(type) {
		case []byte:
			req.Body = ioutil.NopCloser(bytes.NewReader(b))
			req.ContentLength = int64(len(b))
		default:
			return errors.New("body type should be []byte")
		}
	}
	return nil
}

func getClient(client interface{}) (*http.Client, error) {
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
