package tokenx

import (
	"errors"
	"testing"
)

type stubSignTokenStore struct {
	data        map[string]string
	deleteCalls []string
}

func (s *stubSignTokenStore) Set(key string, value string, expireSeconds int) error {
	if s.data == nil {
		s.data = make(map[string]string)
	}
	s.data[key] = value
	return nil
}

func (s *stubSignTokenStore) Get(key string) (string, error) {
	if s.data == nil {
		return "", nil
	}
	return s.data[key], nil
}

func (s *stubSignTokenStore) Delete(key string) error {
	s.deleteCalls = append(s.deleteCalls, key)
	if s.data != nil {
		delete(s.data, key)
	}
	return nil
}

func (s *stubSignTokenStore) Exists(key string) (bool, error) {
	if s.data == nil {
		return false, nil
	}
	_, ok := s.data[key]
	return ok, nil
}

func (s *stubSignTokenStore) SetExpire(string, int) error {
	return nil
}

func TestSignTokenGenerateAndValidateSingleToken(t *testing.T) {
	store := &stubSignTokenStore{}
	manager := NewSignTokenManager(store, "secret", "admin", 300, 600)

	token, err := manager.GenerateToken("u-1")
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}
	if token.TokenType != TokenTypeSign {
		t.Fatalf("unexpected token type: %s", token.TokenType)
	}
	if token.AccessToken == "" {
		t.Fatal("expected access token")
	}
	if token.RefreshToken != "" {
		t.Fatalf("expected empty refresh token, got %q", token.RefreshToken)
	}
	if token.RefreshExpiresIn != 0 || token.RefreshExpiresAt != 0 {
		t.Fatalf("expected no refresh expiry, got %+v", token)
	}
	if store.data[SignKey("u-1")] != token.AccessToken {
		t.Fatalf("unexpected store data: %#v", store.data)
	}
	if err := manager.ValidateToken("u-1", token.AccessToken); err != nil {
		t.Fatalf("ValidateToken returned error: %v", err)
	}
}

func TestSignTokenRefreshUsesCurrentToken(t *testing.T) {
	store := &stubSignTokenStore{}
	manager := NewSignTokenManager(store, "secret", "admin", 300, 600)

	seed, err := manager.GenerateToken("u-1")
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}

	refreshed, err := manager.RefreshToken("u-1", seed.AccessToken)
	if err != nil {
		t.Fatalf("RefreshToken returned error: %v", err)
	}
	if refreshed.AccessToken == "" {
		t.Fatal("expected refreshed access token")
	}
	if refreshed.AccessToken == seed.AccessToken {
		t.Fatalf("expected rotated token, got same token %q", refreshed.AccessToken)
	}
	if refreshed.RefreshToken != "" {
		t.Fatalf("expected empty refresh token, got %q", refreshed.RefreshToken)
	}
	if err := manager.ValidateToken("u-1", seed.AccessToken); !errors.Is(err, ErrTokenInvalid) {
		t.Fatalf("expected old token invalid, got %v", err)
	}
	if err := manager.ValidateToken("u-1", refreshed.AccessToken); err != nil {
		t.Fatalf("ValidateToken returned error: %v", err)
	}
}

func TestSignTokenRejectsExpiredToken(t *testing.T) {
	store := &stubSignTokenStore{}
	manager := NewSignTokenManager(store, "secret", "admin", -1, 600)

	token, err := manager.GenerateToken("u-1")
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}
	if err := manager.ValidateToken("u-1", token.AccessToken); !errors.Is(err, ErrTokenExpired) {
		t.Fatalf("expected ErrTokenExpired, got %v", err)
	}
}
