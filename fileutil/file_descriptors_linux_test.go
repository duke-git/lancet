package fileutil

import (
	"archive/zip"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime/debug"
	"strings"
	"syscall"
	"testing"
)

// Run each limit check in its own process so the process-wide resource limit
// and disabled garbage collector cannot affect other package tests.
func withLowDescriptorLimit(t *testing.T, check func()) {
	t.Helper()
	if os.Getenv("LANCET_DESCRIPTOR_TEST") != t.Name() {
		executable, err := os.Executable()
		if err != nil {
			t.Fatal(err)
		}
		cmd := exec.Command(executable, "-test.run=^"+regexp.QuoteMeta(t.Name())+"$", "-test.timeout=15s")
		cmd.Env = append(os.Environ(), "LANCET_DESCRIPTOR_TEST="+t.Name())
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("descriptor-limit subprocess failed: %v\n%s", err, out)
		}
		return
	}
	var limit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
		t.Fatal(err)
	}
	originalLimit := limit
	defer syscall.Setrlimit(syscall.RLIMIT_NOFILE, &originalLimit)
	if limit.Cur > 64 {
		limit.Cur = 64
		if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
			t.Fatal(err)
		}
	}
	// A file finalizer must not hide a missing Close while this test runs.
	debug.SetGCPercent(-1)
	check()
}

func TestMiMeTypeBoundsOwnedDescriptors(t *testing.T) {
	for _, kind := range []string{"content", "empty", "directory"} {
		t.Run(kind, func(t *testing.T) {
			withLowDescriptorLimit(t, func() {
				dir := t.TempDir()
				path := dir
				want := ""
				if kind != "directory" {
					path = filepath.Join(dir, "sample.txt")
					var content []byte
					if kind == "content" {
						content = []byte(strings.Repeat("sample file content", 64))
						want = "text/plain; charset=utf-8"
					}
					if err := os.WriteFile(path, content, 0600); err != nil {
						t.Fatal(err)
					}
				}
				for i := 0; i < 256; i++ {
					if got := MiMeType(path); got != want {
						t.Fatalf("call %d: MiMeType = %q, want %q", i, got, want)
					}
				}
				probe, err := os.Open(path)
				if err != nil {
					t.Fatalf("could not open a file after MIME detection: %v", err)
				}
				probe.Close()
			})
		})
	}
}

func TestUnZipBoundsDescriptorsOnErrors(t *testing.T) {
	for _, kind := range []string{"checksum", "output-open"} {
		t.Run(kind, func(t *testing.T) {
			withLowDescriptorLimit(t, func() {
				dir := t.TempDir()
				archive := filepath.Join(dir, "error.zip")
				dest := filepath.Join(dir, "out")
				if kind == "checksum" {
					makeChecksumErrorArchive(t, archive)
				} else {
					makeManyEntryArchive(t, archive, 1)
					// A directory at the output file path forces OpenFile to fail.
					if err := os.MkdirAll(filepath.Join(dest, "nested", "file-000.txt"), 0700); err != nil {
						t.Fatal(err)
					}
				}
				for i := 0; i < 128; i++ {
					err := UnZip(archive, dest)
					if kind == "checksum" {
						if !errors.Is(err, zip.ErrChecksum) {
							t.Fatalf("attempt %d: got %v, want checksum error", i, err)
						}
					} else {
						var pathErr *os.PathError
						if !errors.As(err, &pathErr) || pathErr.Op != "open" || pathErr.Path != filepath.Join(dest, "nested", "file-000.txt") {
							t.Fatalf("attempt %d: got %v, want output-open error", i, err)
						}
					}
				}
				probe, err := os.Open(archive)
				if err != nil {
					t.Fatalf("descriptors leaked after extraction errors: %v", err)
				}
				probe.Close()
			})
		})
	}
}

func TestUnZipBoundsEntryDescriptors(t *testing.T) {
	withLowDescriptorLimit(t, func() {
		dir := t.TempDir()
		archive := filepath.Join(dir, "many.zip")
		dest := filepath.Join(dir, "out")
		const entries = 128
		makeManyEntryArchive(t, archive, entries)
		if err := UnZip(archive, dest); err != nil {
			t.Fatalf("UnZip of %d entries with 64-descriptor limit: %v", entries, err)
		}
		checkManyExtractedEntries(t, dest, entries)
	})
}

func TestZipAppendEntryBoundsDescriptorsOnErrors(t *testing.T) {
	for _, kind := range []string{"missing-archive", "invalid-archive", "missing-source", "checksum"} {
		t.Run(kind, func(t *testing.T) {
			withLowDescriptorLimit(t, func() {
				dir := t.TempDir()
				t.Setenv("TMPDIR", dir)
				archive := filepath.Join(dir, "input.zip")
				source := filepath.Join(dir, "missing-source")
				switch kind {
				case "invalid-archive":
					if err := os.WriteFile(archive, []byte("not a zip archive"), 0600); err != nil {
						t.Fatal(err)
					}
				case "missing-source":
					makeManyEntryArchive(t, archive, 1)
				case "checksum":
					makeChecksumErrorArchive(t, archive)
				}
				original, _ := os.ReadFile(archive)
				for i := 0; i < 128; i++ {
					err := ZipAppendEntry(source, archive)
					switch kind {
					case "missing-archive", "missing-source":
						if !errors.Is(err, os.ErrNotExist) {
							t.Fatalf("attempt %d: got %v, want not-exist error", i, err)
						}
					case "invalid-archive":
						if !errors.Is(err, zip.ErrFormat) {
							t.Fatalf("attempt %d: got %v, want zip format error", i, err)
						}
					case "checksum":
						if !errors.Is(err, zip.ErrChecksum) {
							t.Fatalf("attempt %d: got %v, want checksum error", i, err)
						}
					}
				}
				if kind != "missing-archive" {
					after, err := os.ReadFile(archive)
					if err != nil {
						t.Fatal(err)
					}
					if string(after) != string(original) {
						t.Fatal("failed append changed original archive")
					}
				}
				entries, err := os.ReadDir(dir)
				if err != nil {
					t.Fatal(err)
				}
				for _, entry := range entries {
					if strings.HasPrefix(entry.Name(), "temp.zip") {
						t.Fatalf("temporary archive remains: %s", entry.Name())
					}
				}
			})
		})
	}
}
