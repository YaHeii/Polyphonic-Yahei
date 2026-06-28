// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package website

import (
	"context"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/systemx"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSystemStateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取服务器信息
func NewGetSystemStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSystemStateLogic {
	return &GetSystemStateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSystemStateLogic) GetSystemState(req *types.EmptyReq) (resp *types.Server, err error) {
	server := &types.Server{}

	osInfo := systemx.InitOS()
	server.Os = types.ServerOs{
		Goos:         osInfo.GOOS,
		NumCpu:       int64(osInfo.NumCPU),
		Compiler:     osInfo.Compiler,
		GoVersion:    osInfo.GoVersion,
		NumGoroutine: int64(osInfo.NumGoroutine),
	}

	cpuInfo, err := systemx.InitCPU()
	if err != nil {
		return server, err
	}
	server.Cpu = types.ServerCpu{
		Cpus:  cpuInfo.Cpus,
		Cores: int64(cpuInfo.Cores),
	}

	ramInfo, err := systemx.InitRAM()
	if err != nil {
		return server, err
	}
	server.Ram = types.ServerRam{
		UsedMb:      int64(ramInfo.UsedMB),
		TotalMb:     int64(ramInfo.TotalMB),
		UsedPercent: int64(ramInfo.UsedPercent),
	}

	diskInfo, err := systemx.InitDisk()
	if err != nil {
		return server, err
	}
	server.Disk = types.ServerDisk{
		UsedMb:      int64(diskInfo.UsedMB),
		UsedGb:      int64(diskInfo.UsedGB),
		TotalMb:     int64(diskInfo.TotalMB),
		TotalGb:     int64(diskInfo.TotalGB),
		UsedPercent: int64(diskInfo.UsedPercent),
	}

	return server, nil
}
