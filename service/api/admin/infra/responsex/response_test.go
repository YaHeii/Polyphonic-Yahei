package responsex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
)

func TestResponseUsesPostgresErrorMapping(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	pgErr := &pgconn.PgError{
		Code:    "23505",
		Message: "duplicate key value violates unique constraint",
	}

	Response(req, rec, nil, fmt.Errorf("wrapped: %w", pgErr))

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status: %d", rec.Code)
	}

	var body Body
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if body.Code != http.StatusInternalServerError {
		t.Fatalf("unexpected body code: %d", body.Code)
	}
	if body.Msg != pgErr.Error() {
		t.Fatalf("unexpected body message: %q", body.Msg)
	}
}
