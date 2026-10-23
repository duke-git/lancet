package netutil

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type uploadResponseBody struct {
	io.Reader
	closed bool
}

func (b *uploadResponseBody) Close() error { b.closed = true; return nil }

type uploadTransport struct{ response *http.Response }

func (t uploadTransport) RoundTrip(*http.Request) (*http.Response, error) { return t.response, nil }

func TestUploadFileResponse(t *testing.T) {
	path := filepath.Join(t.TempDir(), "upload.txt")
	if err := os.WriteFile(path, []byte("content"), 0600); err != nil {
		t.Fatal(err)
	}
	original := http.DefaultClient
	defer func() { http.DefaultClient = original }()
	for _, status := range []int{http.StatusCreated, http.StatusInternalServerError} {
		body := &uploadResponseBody{Reader: strings.NewReader("response")}
		http.DefaultClient = &http.Client{Transport: uploadTransport{&http.Response{StatusCode: status, Body: body}}}
		ok, err := UploadFile(path, "https://example.com/upload")
		if ok != (status == http.StatusCreated) || (err != nil) != (status != http.StatusCreated) {
			t.Errorf("status %d: got %v, %v", status, ok, err)
		}
		if !body.closed {
			t.Errorf("status %d: response body was not closed", status)
		}
	}
}
