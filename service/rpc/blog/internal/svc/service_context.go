package svc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/collection"
	gzcache "github.com/zeromicro/go-zero/core/stores/cache"
	gzpostgres "github.com/zeromicro/go-zero/core/stores/postgres"
	gzredis "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/YaHeii/Polyphonic-Yahei/service/model"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/common/online"
	"github.com/YaHeii/Polyphonic-Yahei/service/rpc/blog/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	SqlConn    sqlx.SqlConn
	Redis      *redis.Client
	LocalCache *collection.Cache

	OnlineUserService *online.OnlineUserService

	// account models
	TUserModel      model.TUserModel
	TUserOauthModel model.TUserOauthModel
	TRoleModel      model.TRoleModel
	TApiModel       model.TApiModel
	TMenuModel      model.TMenuModel
	TUserRoleModel  model.TUserRoleModel
	TRoleApiModel   model.TRoleApiModel
	TRoleMenuModel  model.TRoleMenuModel

	// blog models
	TArticleModel  model.TArticleModel
	TCategoryModel model.TCategoryModel
	TTagModel      model.TTagModel

	// message models
	TCommentModel      model.TCommentModel
	TMessageModel      model.TMessageModel
	TSystemNoticeModel model.TSystemNoticeModel

	// website models
	TWebsiteConfigModel   model.TWebsiteConfigModel
	TAlbumModel           model.TAlbumModel
	TPhotoModel           model.TPhotoModel
	TFriendModel          model.TFriendModel
	TTalkModel            model.TTalkModel
	TPageModel            model.TPageModel
	TVisitDailyStatsModel model.TVisitDailyStatsModel
	TVisitorModel         model.TVisitorModel

	TVisitLogModel     model.TVisitLogModel
	TLoginLogModel     model.TLoginLogModel
	TOperationLogModel model.TOperationLogModel
	TFileLogModel      model.TFileLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := buildPostgresDSN(c.PgsqlConf)
	sqlConn := gzpostgres.New(dsn)

	rds, err := ConnectRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}

	cache, err := collection.NewCache(60 * time.Minute)
	if err != nil {
		panic(err)
	}

	modelCacheConf := buildModelCacheConf(c.RedisConf)

	return &ServiceContext{
		Config:     c,
		SqlConn:    sqlConn,
		Redis:      rds,
		LocalCache: cache,

		OnlineUserService: online.NewOnlineUserService(rds, 3600*24),
		// account models
		TUserModel:      model.NewTUserModel(sqlConn, modelCacheConf),
		TUserOauthModel: model.NewTUserOauthModel(sqlConn, modelCacheConf),
		TRoleModel:      model.NewTRoleModel(sqlConn, modelCacheConf),
		TApiModel:       model.NewTApiModel(sqlConn, modelCacheConf),
		TMenuModel:      model.NewTMenuModel(sqlConn, modelCacheConf),
		TUserRoleModel:  model.NewTUserRoleModel(sqlConn),
		TRoleApiModel:   model.NewTRoleApiModel(sqlConn),
		TRoleMenuModel:  model.NewTRoleMenuModel(sqlConn),
		// blog models
		TArticleModel:  model.NewTArticleModel(sqlConn, modelCacheConf),
		TCategoryModel: model.NewTCategoryModel(sqlConn, modelCacheConf),
		TTagModel:      model.NewTTagModel(sqlConn, modelCacheConf),
		// message models
		TCommentModel:      model.NewTCommentModel(sqlConn, modelCacheConf),
		TMessageModel:      model.NewTMessageModel(sqlConn, modelCacheConf),
		TSystemNoticeModel: model.NewTSystemNoticeModel(sqlConn, modelCacheConf),
		// website models
		TWebsiteConfigModel:   model.NewTWebsiteConfigModel(sqlConn, modelCacheConf),
		TAlbumModel:           model.NewTAlbumModel(sqlConn, modelCacheConf),
		TPhotoModel:           model.NewTPhotoModel(sqlConn, modelCacheConf),
		TFriendModel:          model.NewTFriendModel(sqlConn, modelCacheConf),
		TTalkModel:            model.NewTTalkModel(sqlConn, modelCacheConf),
		TPageModel:            model.NewTPageModel(sqlConn, modelCacheConf),
		TVisitDailyStatsModel: model.NewTVisitDailyStatsModel(sqlConn),
		TVisitorModel:         model.NewTVisitorModel(sqlConn),
		TVisitLogModel:        model.NewTVisitLogModel(sqlConn),
		TLoginLogModel:        model.NewTLoginLogModel(sqlConn),
		TOperationLogModel:    model.NewTOperationLogModel(sqlConn),
		TFileLogModel:         model.NewTFileLogModel(sqlConn),
	}
}

func buildPostgresDSN(c config.PgsqlConf) string {
	parts := []string{
		fmt.Sprintf("host=%s", c.Host),
		fmt.Sprintf("port=%s", c.Port),
		fmt.Sprintf("user=%s", c.Username),
		fmt.Sprintf("password=%s", c.Password),
		fmt.Sprintf("dbname=%s", c.Dbname),
	}

	if extra := strings.TrimSpace(c.Config); extra != "" {
		parts = append(parts, extra)
	}

	return strings.Join(parts, " ")
}

func buildModelCacheConf(c config.RedisConf) gzcache.CacheConf {
	return gzcache.CacheConf{
		{
			RedisConf: gzredis.RedisConf{
				Host: fmt.Sprintf("%s:%s", c.Host, c.Port),
				Type: "node",
				Pass: c.Password,
			},
			Weight: 100,
		},
	}
}

func ConnectRedis(c config.RedisConf) (*redis.Client, error) {
	address := c.Host + ":" + c.Port
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Username: "",
		Password: c.Password,
		DB:       c.DB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("redis 连接失败: %v", err)
	}

	return client, nil
}
