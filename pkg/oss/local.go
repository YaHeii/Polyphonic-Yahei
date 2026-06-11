package oss

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Local struct {
	rootDir string
	baseURL string
}

func (s *Local) UploadFile(f io.Reader, prefix string, filename string) (fileURL string, err error) {
	key := cleanRelativePath(path.Join(prefix, filename))
	target := filepath.Join(s.rootDir, filepath.FromSlash(key))

	if err := os.MkdirAll(filepath.Dir(target), os.ModePerm); err != nil {
		return "", fmt.Errorf("Local.UploadFile MkdirAll() Failed, err: %v", err)
	}

	out, err := os.Create(target)
	if err != nil {
		return "", fmt.Errorf("Local.UploadFile Create() Failed, err: %v", err)
	}
	defer out.Close()

	_, copyErr := io.Copy(out, f)
	if copyErr != nil {
		return "", fmt.Errorf("Local.UploadFile Copy() Failed, err: %v", copyErr)
	}

	return s.publicURL(key), nil
}

func (s *Local) DeleteFile(filepath string) error {
	relativePath := cleanRelativePath(filepath)
	p := filepathJoin(s.rootDir, relativePath)

	if err := os.Remove(p); err != nil {
		return fmt.Errorf("Local.DeleteFile Remove() Failed, err: %v", err)
	}
	return nil
}

func (s *Local) ListFiles(prefix string, limit int) (files []*FileInfo, err error) {
	root := filepathJoin(s.rootDir, cleanRelativePath(prefix))
	walkErr := filepath.Walk(root, func(currentPath string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if info.IsDir() {
			return nil
		}

		relativePath, err := filepath.Rel(s.rootDir, currentPath)
		if err != nil {
			return err
		}
		relativePath = filepath.ToSlash(relativePath)

		f := &FileInfo{
			IsDir:    false,
			FilePath: relativePath,
			FileName: info.Name(),
			FileType: path.Ext(info.Name()),
			FileSize: info.Size(),
			FileUrl:  s.publicURL(relativePath),
			UpTime:   info.ModTime().UnixMilli(),
		}
		files = append(files, f)

		if limit > 0 && len(files) >= limit {
			return errLimitReached
		}

		return nil
	})

	if errors.Is(walkErr, os.ErrNotExist) {
		return []*FileInfo{}, nil
	}
	if walkErr != nil && !errors.Is(walkErr, errLimitReached) {
		return nil, walkErr
	}
	return files, nil
}

func NewLocal(dir, baseURL string) *Local {
	return &Local{
		rootDir: (&Config{LocalRootDir: dir}).RootDir(),
		baseURL: (&Config{LocalBaseURL: baseURL}).BaseURL(),
	}
}

var errLimitReached = errors.New("limit reached")

func (s *Local) publicURL(relativePath string) string {
	return strings.TrimRight(s.baseURL, "/") + "/" + cleanRelativePath(relativePath)
}

func filepathJoin(rootDir, relativePath string) string {
	return filepath.Join(rootDir, filepath.FromSlash(relativePath))
}

func cleanRelativePath(value string) string {
	cleaned := path.Clean("/" + strings.TrimSpace(value))
	return strings.TrimPrefix(cleaned, "/")
}
