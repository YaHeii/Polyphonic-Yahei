// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package file_log

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/client/syslogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesFileLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除文件日志
func NewDeletesFileLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesFileLogLogic {
	return &DeletesFileLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletesFileLogLogic) DeletesFileLog(req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := &syslogrpc.DeletesFileLogReq{
		Ids: req.Ids,
	}

	out, err := l.svcCtx.SyslogRpc.DeletesFileLog(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{SuccessCount: out.SuccessCount}, nil
}
