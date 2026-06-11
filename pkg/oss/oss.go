package oss

import (
	"fmt"
	"io"
	"path"
	"strings"
	"time"
)

type Config struct {
	LocalRootDir string `json:"local-root-dir" yaml:"local-root-dir"`
	LocalBaseURL string `json:"local-base-url" yaml:"local-base-url"`
}

const (
	DefaultLocalRootDir = "/home/yahei"
	DefaultLocalBaseURL = "/static"
)

func (c *Config) RootDir() string {
	if c == nil {
		return DefaultLocalRootDir
	}

	rootDir := strings.TrimSpace(c.LocalRootDir)
	if rootDir == "" {
		return DefaultLocalRootDir
	}

	return rootDir
}

func (c *Config) BaseURL() string {
	if c == nil {
		return DefaultLocalBaseURL
	}

	baseURL := strings.TrimSpace(c.LocalBaseURL)
	if baseURL == "" {
		return DefaultLocalBaseURL
	}

	baseURL = "/" + strings.Trim(baseURL, "/")
	if baseURL == "/" {
		return DefaultLocalBaseURL
	}

	return baseURL
}

// Uploader oss（Object Storage Service，对象存储服务）
type Uploader interface {
	// return BucketUrl+BasePath+prefix+filename
	UploadFile(f io.Reader, prefix string, filename string) (filepath string, err error)
	DeleteFile(filepath string) error
	ListFiles(prefix string, limit int) (files []*FileInfo, err error)
}

type FileInfo struct {
	IsDir    bool   `json:"is_dir"`    // 是否是目录
	FilePath string `json:"file_path"` // 文件路径
	FileName string `json:"file_name"` // 文件名
	FileType string `json:"file_type"` // 文件类型
	FileSize int64  `json:"file_size"` // 文件大小
	FileUrl  string `json:"file_url"`  // 文件访问地址
	UpTime   int64  `json:"up_time"`   // 上传时间
}

func NewFileNameWithDateTime(filename string) string {
	// 读取文件后缀
	ext := path.Ext(filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(filename, ext)
	// 拼接新文件名
	newFilename := fmt.Sprintf("%s-%s%s", name, time.Now().Format("20060102150405"), ext)

	return newFilename
}
