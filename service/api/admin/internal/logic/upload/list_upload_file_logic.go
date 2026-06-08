// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package upload

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取文件列表
func NewListUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUploadFileLogic {
	return &ListUploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUploadFileLogic) ListUploadFile(req *types.ListUploadFileReq) (resp *types.PageResp, err error) {
	files, err := l.svcCtx.Uploader.ListFiles(req.FilePath, int(req.Limit))
	if err != nil {
		return nil, err
	}

	list := make([]*types.FileInfoVO, 0, len(files))
	for _, file := range files {
		list = append(list, &types.FileInfoVO{
			FilePath:  file.FilePath,
			FileName:  file.FileName,
			FileType:  file.FileType,
			FileSize:  file.FileSize,
			FileUrl:   file.FileUrl,
			UpdatedAt: file.UpTime,
		})
	}

	return &types.PageResp{
		Page:     1,
		PageSize: req.Limit,
		Total:    int64(len(list)),
		List:     list,
	}, nil
}
