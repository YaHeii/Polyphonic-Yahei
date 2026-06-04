package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TUserModel = (*customTUserModel)(nil)

type (
	// TUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTUserModel.
	TUserModel interface {
		tUserModel
		FindOneByEmail(ctx context.Context, email string) (*TUser, error)
		FindOneByPhone(ctx context.Context, phone string) (*TUser, error)
		UpdateNicknameInfo(ctx context.Context, userID, nickname, info string) error
		UpdateAvatarByUserID(ctx context.Context, userID, avatar string) error
		UpdatePasswordByUserID(ctx context.Context, userID, password string) error
		UpdateEmailByUserID(ctx context.Context, userID, email string) error
		UpdatePhoneByUserID(ctx context.Context, userID, phone string) error
	}

	customTUserModel struct {
		*defaultTUserModel
	}
)

var cachePublicTUserEmailPrefix = "cache:public:tUser:email:"
var cachePublicTUserPhonePrefix = "cache:public:tUser:phone:"

// NewTUserModel returns a model for the database table.
func NewTUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TUserModel {
	return &customTUserModel{
		defaultTUserModel: newTUserModel(conn, c, opts...),
	}
}

func (m *customTUserModel) FindOneByEmail(ctx context.Context, email string) (*TUser, error) {
	publicTUserEmailKey := fmt.Sprintf("%s%v", cachePublicTUserEmailPrefix, email)
	var resp TUser
	err := m.QueryRowIndexCtx(ctx, &resp, publicTUserEmailKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (any, error) {
		query := fmt.Sprintf("select %s from %s where email = $1 limit 1", tUserRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, email); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customTUserModel) FindOneByPhone(ctx context.Context, phone string) (*TUser, error) {
	publicTUserPhoneKey := fmt.Sprintf("%s%v", cachePublicTUserPhonePrefix, phone)
	var resp TUser
	err := m.QueryRowIndexCtx(ctx, &resp, publicTUserPhoneKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (any, error) {
		query := fmt.Sprintf("select %s from %s where phone = $1 limit 1", tUserRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, phone); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customTUserModel) UpdateNicknameInfo(ctx context.Context, userID, nickname, info string) error {
	user, err := m.FindOneByUserId(ctx, userID)
	if err != nil {
		return err
	}

	user.Nickname = nickname
	user.Info = info
	return m.updateUserWithExtraKeys(ctx, user, nil)
}

func (m *customTUserModel) UpdateAvatarByUserID(ctx context.Context, userID, avatar string) error {
	user, err := m.FindOneByUserId(ctx, userID)
	if err != nil {
		return err
	}

	user.Avatar = avatar
	return m.updateUserWithExtraKeys(ctx, user, nil)
}

func (m *customTUserModel) UpdatePasswordByUserID(ctx context.Context, userID, password string) error {
	user, err := m.FindOneByUserId(ctx, userID)
	if err != nil {
		return err
	}

	user.Password = password
	return m.updateUserWithExtraKeys(ctx, user, nil)
}

func (m *customTUserModel) UpdateEmailByUserID(ctx context.Context, userID, email string) error {
	user, err := m.FindOneByUserId(ctx, userID)
	if err != nil {
		return err
	}

	oldEmail := user.Email
	user.Email = email

	keys := []string{
		fmt.Sprintf("%s%v", cachePublicTUserEmailPrefix, oldEmail),
		fmt.Sprintf("%s%v", cachePublicTUserEmailPrefix, email),
	}
	return m.updateUserWithExtraKeys(ctx, user, keys)
}

func (m *customTUserModel) UpdatePhoneByUserID(ctx context.Context, userID, phone string) error {
	user, err := m.FindOneByUserId(ctx, userID)
	if err != nil {
		return err
	}

	oldPhone := user.Phone
	user.Phone = phone

	keys := []string{
		fmt.Sprintf("%s%v", cachePublicTUserPhonePrefix, oldPhone),
		fmt.Sprintf("%s%v", cachePublicTUserPhonePrefix, phone),
	}
	return m.updateUserWithExtraKeys(ctx, user, keys)
}

func (m *customTUserModel) updateUserWithExtraKeys(ctx context.Context, user *TUser, extraKeys []string) error {
	publicTUserIdKey := fmt.Sprintf("%s%v", cachePublicTUserIdPrefix, user.Id)
	publicTUserUserIdKey := fmt.Sprintf("%s%v", cachePublicTUserUserIdPrefix, user.UserId)
	publicTUserUsernameKey := fmt.Sprintf("%s%v", cachePublicTUserUsernamePrefix, user.Username)

	keys := []string{publicTUserIdKey, publicTUserUserIdKey, publicTUserUsernameKey}
	for _, key := range extraKeys {
		if key != "" {
			keys = append(keys, key)
		}
	}

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, tUserRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, user.Id, user.UserId, user.Username, user.Password, user.Nickname, user.Avatar, user.Email, user.Phone, user.Info, user.Status, user.RegisterType, user.IpAddress, user.IpSource)
	}, keys...)
	return err
}
