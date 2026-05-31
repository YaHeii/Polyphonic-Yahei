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
	// todo: add your logic here and delete this line

	return
}
