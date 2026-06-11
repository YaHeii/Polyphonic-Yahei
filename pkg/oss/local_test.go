package oss

import (
	"os"
	"strings"
	"testing"
)

func TestLocalUploadListAndDelete(t *testing.T) {
	rootDir := t.TempDir()
	uploader := NewLocal(rootDir, "/files")

	fileURL, err := uploader.UploadFile(strings.NewReader("hello"), "article/image", "a.txt")
	if err != nil {
		t.Fatalf("UploadFile returned error: %v", err)
	}

	if fileURL != "/files/article/image/a.txt" {
		t.Fatalf("unexpected file url: %q", fileURL)
	}

	if _, err := os.Stat(rootDir + "/article/image/a.txt"); err != nil {
		t.Fatalf("uploaded file not found on disk: %v", err)
	}

	files, err := uploader.ListFiles("article", 10)
	if err != nil {
		t.Fatalf("ListFiles returned error: %v", err)
	}

	if len(files) != 1 {
		t.Fatalf("expected 1 file, got %d", len(files))
	}

	if files[0].FilePath != "article/image/a.txt" {
		t.Fatalf("unexpected file path: %q", files[0].FilePath)
	}

	if files[0].FileUrl != "/files/article/image/a.txt" {
		t.Fatalf("unexpected listed file url: %q", files[0].FileUrl)
	}

	if err := uploader.DeleteFile("article/image/a.txt"); err != nil {
		t.Fatalf("DeleteFile returned error: %v", err)
	}

	if _, err := os.Stat(rootDir + "/article/image/a.txt"); !os.IsNotExist(err) {
		t.Fatalf("expected file to be deleted, stat err=%v", err)
	}
}

func TestNewLocalUsesDefaults(t *testing.T) {
	uploader := NewLocal("", "")

	if uploader.rootDir != DefaultLocalRootDir {
		t.Fatalf("unexpected default root dir: %q", uploader.rootDir)
	}
	if uploader.baseURL != DefaultLocalBaseURL {
		t.Fatalf("unexpected default base url: %q", uploader.baseURL)
	}
}

func TestConfigUsesDefaultsAndNormalizesValues(t *testing.T) {
	var empty *Config
	if empty.RootDir() != DefaultLocalRootDir {
		t.Fatalf("unexpected empty root dir: %q", empty.RootDir())
	}
	if empty.BaseURL() != DefaultLocalBaseURL {
		t.Fatalf("unexpected empty base url: %q", empty.BaseURL())
	}

	cfg := &Config{
		LocalRootDir: "  /data/upload  ",
		LocalBaseURL: " static/ ",
	}
	if cfg.RootDir() != "/data/upload" {
		t.Fatalf("unexpected config root dir: %q", cfg.RootDir())
	}
	if cfg.BaseURL() != "/static" {
		t.Fatalf("unexpected config base url: %q", cfg.BaseURL())
	}
}
