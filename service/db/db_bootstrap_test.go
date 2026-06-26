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
		filepath.Join(root, "service", "db", "migrations", "000002_drop_t_page.up.sql"),
		filepath.Join(root, "service", "db", "migrations", "000002_drop_t_page.down.sql"),
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
		"role_id",
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

func TestRBACSchemaUsesSingleRolePerUser(t *testing.T) {
	root := repoRoot(t)
	upSQL := readFile(t, root, "service", "db", "migrations", "000001_blog_init.up.sql")
	authSeed := readFile(t, root, "service", "db", "seeds", "bootstrap", "001_auth_bootstrap.sql")
	permissionSeed := readFile(t, root, "service", "db", "seeds", "bootstrap", "002_permission_bootstrap.sql")
	makefile := readFile(t, root, "makefile")
	roleTable := upSQL[strings.Index(upSQL, "CREATE TABLE t_role ("):strings.Index(upSQL, "CREATE TABLE t_role_api (")]

	requiredUp := []string{
		"CREATE TABLE t_user (",
		"role_id integer NOT NULL DEFAULT 0",
		"CREATE TABLE t_role (",
		"role_key varchar(64) NOT NULL DEFAULT ''",
		"CONSTRAINT uk_role_key UNIQUE (role_key)",
		"CHECK (role_key IN ('root', 'super_admin', 'visitor'))",
	}
	for _, item := range requiredUp {
		if !strings.Contains(upSQL, item) {
			t.Fatalf("expected baseline migration to contain %q", item)
		}
	}

	forbiddenUp := []string{
		"CREATE TABLE t_user_role (",
	}
	for _, item := range forbiddenUp {
		if strings.Contains(upSQL, item) {
			t.Fatalf("expected baseline migration to exclude %q", item)
		}
	}

	forbiddenRoleFields := []string{
		"parent_id integer NOT NULL DEFAULT 0",
		"role_label varchar(64) NOT NULL DEFAULT ''",
		"is_default boolean NOT NULL DEFAULT false",
	}
	for _, item := range forbiddenRoleFields {
		if strings.Contains(roleTable, item) {
			t.Fatalf("expected t_role schema to exclude %q", item)
		}
	}

	requiredSeed := []string{
		"(1, 'root', 'System Owner', 0)",
		"(2, 'super_admin', 'super admin', 0)",
		"(3, 'visitor', 'default registered visitor', 0)",
		"'root'",
		"'visitor'",
		"'super_admin'",
		"role_id = EXCLUDED.role_id",
	}
	for _, item := range requiredSeed {
		if !strings.Contains(authSeed, item) {
			t.Fatalf("expected auth bootstrap seed to contain %q", item)
		}
	}

	forbiddenSeed := []string{
		"INSERT INTO t_user_role",
		"pg_get_serial_sequence('t_user_role', 'id')",
	}
	for _, item := range forbiddenSeed {
		if strings.Contains(authSeed, item) {
			t.Fatalf("expected auth bootstrap seed to exclude %q", item)
		}
	}

	requiredPermissionSeed := []string{
		"(1, 1, 1)",
		"(101, 2, 1)",
		"(1, 1, 1)",
		"(101, 2, 1)",
	}
	for _, item := range requiredPermissionSeed {
		if !strings.Contains(permissionSeed, item) {
			t.Fatalf("expected permission bootstrap seed to contain %q", item)
		}
	}

	if strings.Contains(makefile, "t_user_role") {
		t.Fatalf("expected makefile relation tables to exclude %q", "t_user_role")
	}
}

func TestDropPageMigrationExistsAndDropsTable(t *testing.T) {
	root := repoRoot(t)
	upSQL := readFile(t, root, "service", "db", "migrations", "000002_drop_t_page.up.sql")
	downSQL := readFile(t, root, "service", "db", "migrations", "000002_drop_t_page.down.sql")

	requiredUp := []string{
		"DROP TRIGGER IF EXISTS trg_t_page_set_updated_at ON t_page;",
		"DROP TABLE IF EXISTS t_page;",
		"DROP TRIGGER IF EXISTS trg_t_tag_set_updated_at ON t_tag;",
		"DROP TABLE IF EXISTS t_tag;",
	}
	for _, item := range requiredUp {
		if !strings.Contains(upSQL, item) {
			t.Fatalf("expected drop migration up to contain %q", item)
		}
	}

	requiredDown := []string{
		"CREATE TABLE t_page (",
		"CREATE TRIGGER trg_t_page_set_updated_at BEFORE UPDATE ON t_page",
		"CREATE TABLE t_tag (",
		"CREATE TRIGGER trg_t_tag_set_updated_at BEFORE UPDATE ON t_tag",
	}
	for _, item := range requiredDown {
		if !strings.Contains(downSQL, item) {
			t.Fatalf("expected drop migration down to contain %q", item)
		}
	}
}

func TestTagManagementSurfaceRemoved(t *testing.T) {
	root := repoRoot(t)
	makefile := readFile(t, root, "makefile")
	seed := readFile(t, root, "service", "db", "seeds", "bootstrap", "002_permission_bootstrap.sql")
	adminAPI := readFile(t, root, "service", "api", "admin", "proto", "admin.api")
	tagAPI := filepath.Join(root, "service", "api", "admin", "proto", "admin", "tag.api")

	forbiddenMakefile := []string{
		"t_tag",
	}
	for _, item := range forbiddenMakefile {
		if strings.Contains(makefile, item) {
			t.Fatalf("expected makefile to exclude %q", item)
		}
	}

	forbiddenSeed := []string{
		"'tag', 'Tag'",
		"/admin-api/v1/tag/find_tag_list",
	}
	for _, item := range forbiddenSeed {
		if strings.Contains(seed, item) {
			t.Fatalf("expected tag permission seed to exclude %q", item)
		}
	}

	if strings.Contains(adminAPI, `import "admin/tag.api"`) {
		t.Fatal(`expected admin.api to exclude import "admin/tag.api"`)
	}

	if _, err := os.Stat(tagAPI); err == nil {
		t.Fatalf("expected tag api file to be removed: %s", tagAPI)
	} else if !os.IsNotExist(err) {
		t.Fatalf("unexpected stat error for %s: %v", tagAPI, err)
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

func TestSingleUserRoleCompatMigrationRemoved(t *testing.T) {
	root := repoRoot(t)

	legacyFiles := []string{
		filepath.Join(root, "service", "db", "migrations", "000003_single_user_role.up.sql"),
		filepath.Join(root, "service", "db", "migrations", "000003_single_user_role.down.sql"),
	}
	for _, path := range legacyFiles {
		if _, err := os.Stat(path); err == nil {
			t.Fatalf("expected compat migration to be removed: %s", path)
		} else if !os.IsNotExist(err) {
			t.Fatalf("unexpected stat error for %s: %v", path, err)
		}
	}
}
