package netutil

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestDownloadFileStatus(t *testing.T) {
	for _, status := range []int{http.StatusOK, http.StatusNotFound, http.StatusInternalServerError} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status)
				w.Write([]byte("response"))
			}))
			defer server.Close()
			path := filepath.Join(t.TempDir(), "download.txt")
			if err := os.WriteFile(path, []byte("original"), 0600); err != nil {
				t.Fatal(err)
			}
			err := DownloadFile(path, server.URL)
			if (err != nil) != (status != http.StatusOK) {
				t.Errorf("unexpected result: %v", err)
			}
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			want := "original"
			if status == http.StatusOK {
				want = "response"
			}
			if string(data) != want {
				t.Errorf("got %q, want %q", data, want)
			}
		})
	}
}
