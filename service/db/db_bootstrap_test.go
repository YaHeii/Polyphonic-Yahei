package db_test

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func repoRoot(t *testing.T) string {
	t.Helper()

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}

	return filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
}

func readFile(t *testing.T, root string, parts ...string) string {
	t.Helper()

	path := filepath.Join(append([]string{root}, parts...)...)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile(%s) failed: %v", path, err)
	}

	return string(data)
}

func TestServiceDBLayoutExists(t *testing.T) {
	root := repoRoot(t)

	expected := []string{
		filepath.Join(root, "service", "db", "README.md"),
		filepath.Join(root, "service", "db", "migrations", "000001_blog_init.up.sql"),
		filepath.Join(root, "service", "db", "migrations", "000001_blog_init.down.sql"),
		filepath.Join(root, "service", "db", "seeds", "bootstrap", "001_auth_bootstrap.sql"),
		filepath.Join(root, "service", "db", "seeds", "bootstrap", "002_permission_bootstrap.sql"),
		filepath.Join(root, "service", "db", "seeds", "bootstrap", "003_site_bootstrap.sql"),
	}

	for _, path := range expected {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("expected path to exist: %s: %v", path, err)
		}
	}
}

func TestBaselineMigrationContainsOnlySchema(t *testing.T) {
	root := repoRoot(t)
	upSQL := readFile(t, root, "service", "db", "migrations", "000001_blog_init.up.sql")

	required := []string{
		"CREATE OR REPLACE FUNCTION set_updated_at()",
		"CREATE TABLE t_user (",
		"CREATE INDEX idx_article_tags ON t_article USING gin (tags);",
		"CREATE TRIGGER trg_t_user_set_updated_at BEFORE UPDATE ON t_user",
	}
	for _, item := range required {
		if !strings.Contains(upSQL, item) {
			t.Fatalf("expected baseline migration to contain %q", item)
		}
	}

	forbidden := []string{
		"INSERT INTO t_user",
		"INSERT INTO t_article",
		"BEGIN;",
		"COMMIT;",
	}
	for _, item := range forbidden {
		if strings.Contains(upSQL, item) {
			t.Fatalf("expected baseline migration to exclude %q", item)
		}
	}
}

func TestBootstrapSeedOnlyContainsRequiredSystemData(t *testing.T) {
	root := repoRoot(t)
	files := []string{
		readFile(t, root, "service", "db", "seeds", "bootstrap", "001_auth_bootstrap.sql"),
		readFile(t, root, "service", "db", "seeds", "bootstrap", "002_permission_bootstrap.sql"),
		readFile(t, root, "service", "db", "seeds", "bootstrap", "003_site_bootstrap.sql"),
	}
	seedSQL := strings.Join(files, "\n")

	required := []string{
		"INSERT INTO t_user",
		"INSERT INTO t_role",
		"INSERT INTO t_user_role",
		"INSERT INTO t_menu",
		"INSERT INTO t_api",
		"INSERT INTO t_role_menu",
		"INSERT INTO t_role_api",
		"INSERT INTO t_website_config",
	}
	for _, item := range required {
		if !strings.Contains(seedSQL, item) {
			t.Fatalf("expected bootstrap seed to contain %q", item)
		}
	}

	forbidden := []string{
		"INSERT INTO t_category",
		"INSERT INTO t_tag",
		"INSERT INTO t_article",
		"INSERT INTO t_message",
		"INSERT INTO t_system_notice",
		"INSERT INTO t_friend",
		"INSERT INTO t_page",
		"INSERT INTO t_album",
		"INSERT INTO t_photo",
	}
	for _, item := range forbidden {
		if strings.Contains(seedSQL, item) {
			t.Fatalf("expected bootstrap seed to exclude %q", item)
		}
	}
}

func TestComposeAndReadmeUseExplicitDatabaseBootstrap(t *testing.T) {
	root := repoRoot(t)
	compose := readFile(t, root, "docker-compose.yaml")
	makefile := readFile(t, root, "makefile")
	readme := readFile(t, root, "README.md")
	dbReadme := readFile(t, root, "service", "db", "README.md")

	forbiddenCompose := []string{
		"./blog-init.sql:/docker-entrypoint-initdb.d/01-blog-init.sql:ro",
		"./seed.sql:/docker-entrypoint-initdb.d/02-seed.sql:ro",
	}
	for _, item := range forbiddenCompose {
		if strings.Contains(compose, item) {
			t.Fatalf("expected docker-compose.yaml to exclude %q", item)
		}
	}

	requiredReadme := []string{
		"make migrate-up",
		"make seed-bootstrap",
	}
	for _, item := range requiredReadme {
		if !strings.Contains(readme, item) {
			t.Fatalf("expected README.md to contain %q", item)
		}
	}

	requiredMakeTargets := []string{
		"migrate-up:",
		"migrate-down:",
		"migrate-version:",
		"migrate-force:",
		"seed-bootstrap:",
		"service/db/migrations",
		"service/db/seeds/bootstrap",
	}
	for _, item := range requiredMakeTargets {
		if !strings.Contains(makefile, item) {
			t.Fatalf("expected makefile to contain %q", item)
		}
	}

	forbiddenDocs := []string{
		"根目录 `blog-init.sql` / `seed.sql` 仅保留为过渡期遗留文件",
		"根目录 `blog-init.sql` / `seed.sql` 仅保留为过渡期遗留文件，不再由 compose 自动执行",
	}
	for _, item := range forbiddenDocs {
		if strings.Contains(readme, item) {
			t.Fatalf("expected README.md to exclude %q", item)
		}
		if strings.Contains(dbReadme, item) {
			t.Fatalf("expected service/db/README.md to exclude %q", item)
		}
	}
}

func TestLegacyRootSQLFilesRemoved(t *testing.T) {
	root := repoRoot(t)

	legacyFiles := []string{
		filepath.Join(root, "blog-init.sql"),
		filepath.Join(root, "seed.sql"),
	}
	for _, path := range legacyFiles {
		if _, err := os.Stat(path); err == nil {
			t.Fatalf("expected legacy file to be removed: %s", path)
		} else if !os.IsNotExist(err) {
			t.Fatalf("unexpected stat error for %s: %v", path, err)
		}
	}
}
