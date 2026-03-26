package utils

import (
	"context"
	"gin_rbac_api/global"
	"time"

	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

func NewCaptchaRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
	}
}

func NewSmsCodeRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "_SMS_",
	}
}

/*
func NewAdminRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Minute * 120,
		PreKey:     "_A_W_", //admin web
	}
}

func NewUserMoibleRedisStore() *RedisStore {
	return &RedisStore{
		PreKey: "_U_M_", //user mobile
	}
}
func NewUserWebRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Minute * 120,
		PreKey:     "_U_W_", //user Web
	}
} */

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

func (rs *RedisStore) Set(id string, value string) error {
	err := global.SYS_REDIS.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	global.SYS_LOG.Info(id + "---" + value)
	if err != nil {
		global.SYS_LOG.Error("RedisStoreSetError!", zap.Error(err))
	}
	return err
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := global.SYS_REDIS.Get(rs.Context, key).Result()
	if err != nil {
		global.SYS_LOG.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := global.SYS_REDIS.Del(rs.Context, key).Err()
		if err != nil {
			global.SYS_LOG.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Del(key string) error {
	err := global.SYS_REDIS.Del(rs.Context, key).Err()
	return err
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}

func (rs *RedisStore) VerifyPhoneCode(phone, captcha string) bool {
	key := rs.PreKey + phone
	v := rs.Get(key, false)
	if v == captcha {
		rs.Del(key)
	}
	return v == captcha
}
