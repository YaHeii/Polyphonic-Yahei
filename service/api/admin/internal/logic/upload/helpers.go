package upload

import (
	"context"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/oss"
	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/cryptox"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"
)

func buildFileInfoVO(log *syslogrpc.FileLog) *types.FileInfoVO {
	if log == nil {
		return nil
	}

	return &types.FileInfoVO{
		FilePath:  log.GetFilePath(),
		FileName:  log.GetFileName(),
		FileType:  log.GetFileType(),
		FileSize:  log.GetFileSize(),
		FileUrl:   log.GetFileUrl(),
		UpdatedAt: log.GetUpdatedAt(),
	}
}

func uploadMultipartFile(ctx context.Context, uploader oss.Uploader, syslogRPC syslogrpc.SyslogRpc, prefix string, header *multipart.FileHeader) (*types.FileInfoVO, error) {
	file, err := header.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	url, err := uploader.UploadFile(file, prefix, oss.NewFileNameWithDateTime(header.Filename))
	if err != nil {
		return nil, err
	}

	rpcCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	out, err := syslogRPC.AddFileLog(rpcCtx, &syslogrpc.AddFileLogReq{
		FilePath: prefix,
		FileName: header.Filename,
		FileType: filepath.Ext(header.Filename),
		FileSize: header.Size,
		FileMd5:  cryptox.Md5v(header.Filename, ""),
		FileUrl:  url,
	})
	if err != nil {
		return nil, err
	}

	return buildFileInfoVO(out.GetFileLog()), nil
}
