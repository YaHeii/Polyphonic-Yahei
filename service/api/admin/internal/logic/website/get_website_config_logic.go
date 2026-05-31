// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package website

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWebsiteConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取网站配置
func NewGetWebsiteConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWebsiteConfigLogic {
	return &GetWebsiteConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWebsiteConfigLogic) GetWebsiteConfig(req *types.EmptyReq) (resp *types.WebsiteConfigVO, err error) {
	// todo: add your logic here and delete this line

	return
}
