// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package upload

import (
	"context"
	"net/http"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MultiUploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传文件列表
func NewMultiUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MultiUploadFileLogic {
	return &MultiUploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MultiUploadFileLogic) MultiUploadFile(req *types.MultiUploadFileReq, r *http.Request) (resp []*types.FileInfoVO, err error) {
	if r.MultipartForm == nil {
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			return nil, err
		}
	}

	files := r.MultipartForm.File["files"]
	resp = make([]*types.FileInfoVO, 0, len(files))
	for _, header := range files {
		info, err := uploadMultipartFile(l.ctx, l.svcCtx.Uploader, l.svcCtx.SyslogRpc, req.FilePath, header)
		if err != nil {
			return nil, err
		}
		resp = append(resp, info)
	}

	return resp, nil
}
