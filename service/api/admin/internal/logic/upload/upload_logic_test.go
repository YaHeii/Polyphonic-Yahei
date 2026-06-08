package upload

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/oss"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc"
)

type stubUploadUploader struct {
	listPrefix  string
	listLimit   int
	listResp    []*oss.FileInfo
	deletePaths []string
	uploadCalls []uploadCall
	uploadResp  string
}

type uploadCall struct {
	prefix   string
	filename string
	content  string
}

func (s *stubUploadUploader) UploadFile(f io.Reader, prefix, filename string) (string, error) {
	data, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	s.uploadCalls = append(s.uploadCalls, uploadCall{
		prefix:   prefix,
		filename: filename,
		content:  string(data),
	})
	return s.uploadResp, nil
}

func (s *stubUploadUploader) DeleteFile(filepath string) error {
	s.deletePaths = append(s.deletePaths, filepath)
	return nil
}

func (s *stubUploadUploader) ListFiles(prefix string, limit int) ([]*oss.FileInfo, error) {
	s.listPrefix = prefix
	s.listLimit = limit
	return s.listResp, nil
}

type stubUploadSyslogRPC struct {
	syslogrpc.SyslogRpc
	addReqs []*syslogrpc.AddFileLogReq
}

func (s *stubUploadSyslogRPC) AddFileLog(_ context.Context, in *syslogrpc.AddFileLogReq, _ ...grpc.CallOption) (*syslogrpc.AddFileLogResp, error) {
	s.addReqs = append(s.addReqs, in)
	return &syslogrpc.AddFileLogResp{
		FileLog: &syslogrpc.FileLog{
			FilePath:  in.FilePath,
			FileName:  in.FileName,
			FileType:  in.FileType,
			FileSize:  in.FileSize,
			FileUrl:   in.FileUrl,
			UpdatedAt: 123,
		},
	}, nil
}

func TestDeletesUploadFileDeletesEachPath(t *testing.T) {
	uploader := &stubUploadUploader{}
	logic := NewDeletesUploadFileLogic(context.Background(), &svc.ServiceContext{Uploader: uploader})

	resp, err := logic.DeletesUploadFile(&types.DeletesUploadFileReq{FilePaths: []string{"a.png", "b.png"}})
	if err != nil {
		t.Fatalf("DeletesUploadFile returned error: %v", err)
	}
	if len(uploader.deletePaths) != 2 || uploader.deletePaths[0] != "a.png" || uploader.deletePaths[1] != "b.png" {
		t.Fatalf("unexpected delete paths: %#v", uploader.deletePaths)
	}
	if resp.SuccessCount != 2 {
		t.Fatalf("unexpected response: %#v", resp)
	}
}

func TestListUploadFileMapsUploaderFiles(t *testing.T) {
	uploader := &stubUploadUploader{
		listResp: []*oss.FileInfo{
			{
				FilePath: "images",
				FileName: "a.png",
				FileType: ".png",
				FileSize: 10,
				FileUrl:  "https://cdn/a.png",
				UpTime:   111,
			},
			{
				FilePath: "images",
				FileName: "b.jpg",
				FileType: ".jpg",
				FileSize: 20,
				FileUrl:  "https://cdn/b.jpg",
				UpTime:   222,
			},
		},
	}
	logic := NewListUploadFileLogic(context.Background(), &svc.ServiceContext{Uploader: uploader})

	resp, err := logic.ListUploadFile(&types.ListUploadFileReq{FilePath: "images", Limit: 5})
	if err != nil {
		t.Fatalf("ListUploadFile returned error: %v", err)
	}
	if uploader.listPrefix != "images" || uploader.listLimit != 5 {
		t.Fatalf("unexpected list request: prefix=%s limit=%d", uploader.listPrefix, uploader.listLimit)
	}
	list, ok := resp.List.([]*types.FileInfoVO)
	if !ok || len(list) != 2 || list[0].FileName != "a.png" || list[1].UpdatedAt != 222 {
		t.Fatalf("unexpected list response: %#v", resp)
	}
}

func TestUploadFileReadsMultipartRequest(t *testing.T) {
	uploader := &stubUploadUploader{uploadResp: "https://cdn/uploaded.png"}
	syslogRPC := &stubUploadSyslogRPC{}
	req := newMultipartRequest(t, map[string]string{"file_path": "images"}, []multipartFile{
		{field: "file", filename: "avatar.png", content: "image-content"},
	})

	var parsed types.UploadFileReq
	if err := httpx.Parse(req, &parsed); err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}

	logic := NewUploadFileLogic(req.Context(), &svc.ServiceContext{
		Uploader:  uploader,
		SyslogRpc: syslogRPC,
	})
	resp, err := logic.UploadFile(&parsed, req)
	if err != nil {
		t.Fatalf("UploadFile returned error: %v", err)
	}

	if len(uploader.uploadCalls) != 1 {
		t.Fatalf("unexpected upload calls: %#v", uploader.uploadCalls)
	}
	call := uploader.uploadCalls[0]
	if call.prefix != "images" || call.content != "image-content" || !strings.HasSuffix(call.filename, ".png") || !strings.Contains(call.filename, "avatar-") {
		t.Fatalf("unexpected upload call: %#v", call)
	}
	if len(syslogRPC.addReqs) != 1 || syslogRPC.addReqs[0].FileName != "avatar.png" || syslogRPC.addReqs[0].FileType != ".png" || syslogRPC.addReqs[0].FileSize != int64(len("image-content")) {
		t.Fatalf("unexpected file log request: %#v", syslogRPC.addReqs)
	}
	if resp.FilePath != "images" || resp.FileUrl != "https://cdn/uploaded.png" || resp.FileName != "avatar.png" {
		t.Fatalf("unexpected upload response: %#v", resp)
	}
}

func TestMultiUploadFileReadsMultipartRequest(t *testing.T) {
	uploader := &stubUploadUploader{uploadResp: "https://cdn/uploaded"}
	syslogRPC := &stubUploadSyslogRPC{}
	req := newMultipartRequest(t, map[string]string{"file_path": "docs"}, []multipartFile{
		{field: "files", filename: "a.txt", content: "A"},
		{field: "files", filename: "b.txt", content: "BB"},
	})

	var parsed types.MultiUploadFileReq
	if err := httpx.Parse(req, &parsed); err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}

	logic := NewMultiUploadFileLogic(req.Context(), &svc.ServiceContext{
		Uploader:  uploader,
		SyslogRpc: syslogRPC,
	})
	resp, err := logic.MultiUploadFile(&parsed, req)
	if err != nil {
		t.Fatalf("MultiUploadFile returned error: %v", err)
	}

	if len(uploader.uploadCalls) != 2 || uploader.uploadCalls[0].prefix != "docs" || uploader.uploadCalls[1].content != "BB" {
		t.Fatalf("unexpected upload calls: %#v", uploader.uploadCalls)
	}
	if len(syslogRPC.addReqs) != 2 || syslogRPC.addReqs[0].FileName != "a.txt" || syslogRPC.addReqs[1].FileName != "b.txt" {
		t.Fatalf("unexpected file log requests: %#v", syslogRPC.addReqs)
	}
	if len(resp) != 2 || resp[0].FileName != "a.txt" || resp[1].FileName != "b.txt" {
		t.Fatalf("unexpected response: %#v", resp)
	}
}

type multipartFile struct {
	field    string
	filename string
	content  string
}

func newMultipartRequest(t *testing.T, fields map[string]string, files []multipartFile) *http.Request {
	t.Helper()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	for key, value := range fields {
		if err := writer.WriteField(key, value); err != nil {
			t.Fatalf("WriteField returned error: %v", err)
		}
	}
	for _, file := range files {
		part, err := writer.CreateFormFile(file.field, file.filename)
		if err != nil {
			t.Fatalf("CreateFormFile returned error: %v", err)
		}
		if _, err := io.Copy(part, strings.NewReader(file.content)); err != nil {
			t.Fatalf("Copy returned error: %v", err)
		}
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("Close returned error: %v", err)
	}

	req := httptest.NewRequest("POST", "/upload", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req
}
