package svc

import (
	"testing"

	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/config"
)

func TestBuildPostgresDSN(t *testing.T) {
	t.Parallel()

	dsn := buildPostgresDSN(config.PgsqlConf{
		Host:     "127.0.0.1",
		Port:     "5432",
		Username: "root",
		Password: "secret",
		Dbname:   "blog-init",
		Config:   "sslmode=disable TimeZone=Asia/Shanghai",
	})

	want := "host=127.0.0.1 port=5432 user=root password=secret dbname=blog-init sslmode=disable TimeZone=Asia/Shanghai"
	if dsn != want {
		t.Fatalf("unexpected dsn: got %q want %q", dsn, want)
	}
}
