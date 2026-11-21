package fileutil

import (
	"archive/zip"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMiMeTypeKeepsCallerFileOpen(t *testing.T) {
	path := filepath.Join(t.TempDir(), "sample.txt")
	if err := os.WriteFile(path, []byte(strings.Repeat("a", 1024)), 0600); err != nil {
		t.Fatal(err)
	}
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if got := MiMeType(f); got != "text/plain; charset=utf-8" {
		t.Fatalf("MiMeType(file) = %q", got)
	}
	if _, err := f.Stat(); err != nil {
		t.Fatalf("caller-owned file was closed: %v", err)
	}
	if offset, err := f.Seek(0, 1); err != nil || offset != 512 {
		t.Fatalf("caller-owned file position = %d, %v; want 512, nil", offset, err)
	}
}

func makeManyEntryArchive(t *testing.T, path string, entries int) {
	t.Helper()
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	w := zip.NewWriter(f)
	for i := 0; i < entries; i++ {
		entry, err := w.Create(fmt.Sprintf("nested/file-%03d.txt", i))
		if err != nil {
			f.Close()
			t.Fatal(err)
		}
		if _, err := fmt.Fprintf(entry, "entry %d", i); err != nil {
			f.Close()
			t.Fatal(err)
		}
	}
	if err := w.Close(); err != nil {
		f.Close()
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
}

func makeChecksumErrorArchive(t *testing.T, path string) {
	t.Helper()
	var data bytes.Buffer
	w := zip.NewWriter(&data)
	entry, err := w.CreateHeader(&zip.FileHeader{Name: "file.txt", Method: zip.Store})
	if err != nil {
		t.Fatal(err)
	}
	payload := []byte("checksum test payload")
	if _, err := entry.Write(payload); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	encoded := data.Bytes()
	index := bytes.Index(encoded, payload)
	if index < 0 {
		t.Fatal("stored payload not found in archive")
	}
	encoded[index] ^= 1
	if err := os.WriteFile(path, encoded, 0600); err != nil {
		t.Fatal(err)
	}
}

func checkManyExtractedEntries(t *testing.T, dest string, entries int) {
	t.Helper()
	for i := 0; i < entries; i++ {
		path := filepath.Join(dest, "nested", fmt.Sprintf("file-%03d.txt", i))
		got, err := os.ReadFile(path)
		if err != nil || string(got) != fmt.Sprintf("entry %d", i) {
			t.Fatalf("entry %d = %q, %v", i, got, err)
		}
	}
}
