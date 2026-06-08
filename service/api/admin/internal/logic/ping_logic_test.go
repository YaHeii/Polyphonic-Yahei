package logic

import (
	"context"
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/config"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/svc"
	"github.com/YaHeii/Polyphonic-Yahei/service/api/admin/internal/types"
)

func TestPingReturnsConfigInfo(t *testing.T) {
	logic := NewPingLogic(context.Background(), &svc.ServiceContext{
		Config: config.Config{},
	})
	logic.svcCtx.Config.Mode = "dev"
	logic.svcCtx.Config.Name = "admin-api"

	resp, err := logic.Ping(&types.PingReq{})
	if err != nil {
		t.Fatalf("Ping returned error: %v", err)
	}
	if resp.Env != "dev" || resp.Name != "admin-api" {
		t.Fatalf("unexpected ping response: %#v", resp)
	}
	if resp.Runtime == "" {
		t.Fatalf("expected runtime to be set")
	}
}
