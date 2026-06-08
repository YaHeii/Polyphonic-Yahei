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

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传文件
func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFileLogic) UploadFile(req *types.UploadFileReq, r *http.Request) (resp *types.FileInfoVO, err error) {
	_, header, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}

	return uploadMultipartFile(l.ctx, l.svcCtx.Uploader, l.svcCtx.SyslogRpc, req.FilePath, header)
}
