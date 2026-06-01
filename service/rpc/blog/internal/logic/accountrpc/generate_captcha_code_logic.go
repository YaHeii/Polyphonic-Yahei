package accountrpclogic

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/pb/accountrpc"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateCaptchaCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateCaptchaCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateCaptchaCodeLogic {
	return &GenerateCaptchaCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 生成验证码
func (l *GenerateCaptchaCodeLogic) GenerateCaptchaCode(in *accountrpc.GenerateCaptchaCodeReq) (*accountrpc.GenerateCaptchaCodeResp, error) {
	// todo: add your logic here and delete this line

	return &accountrpc.GenerateCaptchaCodeResp{}, nil
}
