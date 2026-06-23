package tokenx

type JwtTokenManager struct {
	access *AccessTokenManager
	refresh *RefreshTokenManager
}

func NewJwtTokenManager(store TokenStore, secretKey, issuer string, accessExpire, refreshExpire int64) *JwtTokenManager {
	return &JwtTokenManager{
		access:  NewAccessTokenManager(secretKey, issuer, accessExpire),
		refresh: NewRefreshTokenManager(store, refreshExpire),
	}
}

func (m *JwtTokenManager) GenerateToken(uid string) (*Token, error) {
	accessToken, expiresIn, err := m.access.Generate(uid)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshExpiresAt, err := m.refresh.Issue(uid)
	if err != nil {
		return nil, err
	}

	return &Token{
		TokenType:         TokenTypeBearer,
		AccessToken:       accessToken,
		ExpiresIn:         expiresIn,
		RefreshToken:      refreshToken,
		RefreshExpiresIn:  m.refresh.expire,
		RefreshExpiresAt:  refreshExpiresAt,
	}, nil
}

func (m *JwtTokenManager) RefreshToken(refreshToken string) (*RefreshTokenSession, error) {
	uid, err := m.refresh.Verify(refreshToken)
	if err != nil {
		return nil, err
	}
	tk, err := m.GenerateToken(uid)
	if err != nil {
		return nil, err
	}
	return &RefreshTokenSession{
		UserID: uid,
		Token:  tk,
	}, nil
}

func (m *JwtTokenManager) RevokeRefreshToken(uid string) error {
	return m.refresh.Revoke(uid)
}
