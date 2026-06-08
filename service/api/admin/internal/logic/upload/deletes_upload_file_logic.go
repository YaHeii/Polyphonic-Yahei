// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package upload

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesUploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除文件列表
func NewDeletesUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesUploadFileLogic {
	return &DeletesUploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletesUploadFileLogic) DeletesUploadFile(req *types.DeletesUploadFileReq) (resp *types.BatchResp, err error) {
	for _, filePath := range req.FilePaths {
		if err := l.svcCtx.Uploader.DeleteFile(filePath); err != nil {
			return nil, err
		}
	}

	return &types.BatchResp{
		SuccessCount: int64(len(req.FilePaths)),
	}, nil
}
