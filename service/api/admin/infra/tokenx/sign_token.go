package tokenx

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/YaHeii/Polyphonic-Yahei/pkg/utils/cryptox"
)

// SignTokenManager 基于签名的 Token 管理器实现，支持单设备登录。
// Sign token 只有一份签名串，不区分 access/refresh。
type SignTokenManager struct {
	store            TokenStore
	secretKey        string
	issuer           string
	accessExpireTime int64
}

// NewSignTokenManager 创建签名 Token 管理器
func NewSignTokenManager(store TokenStore, secretKey, issuer string, accessExpire, refreshExpire int64) *SignTokenManager {
	return &SignTokenManager{
		store:            store,
		secretKey:        secretKey,
		issuer:           issuer,
		accessExpireTime: accessExpire,
	}
}

// GenerateToken 生成签名 Token
func (m *SignTokenManager) GenerateToken(uid string) (*Token, error) {
	if uid == "" {
		return nil, fmt.Errorf("uid is empty")
	}
	now := time.Now().Unix()
	expireAt := now + m.accessExpireTime
	accessToken := m.sign(uid, expireAt)

	if err := m.store.Set(SignKey(uid), accessToken, int(m.accessExpireTime)); err != nil {
		return nil, err
	}

	return &Token{
		TokenType:   TokenTypeSign,
		AccessToken: accessToken,
		ExpiresIn:   m.accessExpireTime,
	}, nil
}

// ValidateToken 验证 Token 有效性
func (m *SignTokenManager) ValidateToken(uid, accessToken string) error {
	if accessToken == "" {
		return ErrTokenEmpty
	}
	expireAt, err := m.parse(uid, accessToken)
	if err != nil {
		return err
	}
	if time.Now().Unix() > expireAt {
		return ErrTokenExpired
	}

	storedToken, err := m.store.Get(SignKey(uid))
	if err != nil {
		return err
	}
	if storedToken == "" {
		return ErrTokenExpired
	}
	if storedToken != accessToken {
		return ErrTokenInvalid
	}
	return nil
}

// RefreshToken 使用当前签名 token 轮换新的签名 token。
func (m *SignTokenManager) RefreshToken(uid, accessToken string) (*Token, error) {
	if err := m.ValidateToken(uid, accessToken); err != nil {
		return nil, err
	}

	return m.GenerateToken(uid)
}

// RevokeToken 撤销 Token
func (m *SignTokenManager) RevokeToken(uid string, isRefresh bool) error {
	return m.store.Delete(SignKey(uid))
}

// sign 生成签名 token: base64(uid:expireAt:issuedAt:issuer:signature).
func (m *SignTokenManager) sign(uid string, expireAt int64) string {
	issuedAt := time.Now().UnixNano()
	payload := fmt.Sprintf("%s:%d:%d:%s", uid, expireAt, issuedAt, m.issuer)
	signature := cryptox.Md5v(payload, m.secretKey)
	raw := fmt.Sprintf("%s:%s", payload, signature)
	return base64.RawURLEncoding.EncodeToString([]byte(raw))
}

func (m *SignTokenManager) parse(uid, accessToken string) (int64, error) {
	raw, err := base64.RawURLEncoding.DecodeString(accessToken)
	if err != nil {
		return 0, ErrTokenInvalid
	}
	parts := strings.Split(string(raw), ":")
	if len(parts) != 5 {
		return 0, ErrTokenInvalid
	}
	if parts[0] != uid || parts[3] != m.issuer {
		return 0, ErrTokenInvalid
	}
	expireAt, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, ErrTokenInvalid
	}
	if _, err = strconv.ParseInt(parts[2], 10, 64); err != nil {
		return 0, ErrTokenInvalid
	}
	payload := strings.Join(parts[:4], ":")
	if cryptox.Md5v(payload, m.secretKey) != parts[4] {
		return 0, ErrTokenInvalid
	}
	return expireAt, nil
}
