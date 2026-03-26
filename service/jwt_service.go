package service

import (
	"context"
	"gin_rbac_api/global"
	"gin_rbac_api/utils"
)

type JwtService struct{}

//@param: userName string
//@return: redisJWT string, err error

func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.SYS_REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := utils.ParseDuration(global.SYS_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	err = global.SYS_REDIS.Set(context.Background(), userName, jwt, dr).Err()
	return err
}

func (jwtService *JwtService) SetRedisAuth(value string, token string) (err error) {
	dr, err := utils.ParseDuration(global.SYS_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	err = global.SYS_REDIS.Set(context.Background(), "_U_M_"+token, value, dr).Err()
	return err
}
func (jwtService *JwtService) GetRedisAuth(token string) (value string, err error) {
	value, err = global.SYS_REDIS.Get(context.Background(), "_U_M_"+token).Result()
	return value, err
}

func (jwtService *JwtService) DelRedisAuth(token string) (err error) {
	err = global.SYS_REDIS.Del(context.Background(), "_U_M_"+token).Err()
	return err
}

func (jwtService *JwtService) SetRedisAdminAuth(value string, token string) (err error) {
	dr, err := utils.ParseDuration(global.SYS_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	err = global.SYS_REDIS.Set(context.Background(), "_A_W_"+token, value, dr).Err()
	return err
}

func (jwtService *JwtService) GetRedisAdminAuth(token string) (value string, err error) {
	value, err = global.SYS_REDIS.Get(context.Background(), "_A_W_"+token).Result()
	return value, err
}

func (jwtService *JwtService) DelRedisAdminAuth(token string) (err error) {
	err = global.SYS_REDIS.Del(context.Background(), "_A_W_"+token).Err()
	return err
}

func (jwtService *JwtService) SetRedisSmsAuth(code string, phone string) (err error) {
	dr, err := utils.ParseDuration(global.SYS_CONFIG.Captcha.Expired)
	if err != nil {
		return err
	}
	err = global.SYS_REDIS.Set(context.Background(), "_SMS_"+phone, code, dr).Err()
	return err
}

func (jwtService *JwtService) GetRedisSmsAuth(phone string) (value string, err error) {
	value, err = global.SYS_REDIS.Get(context.Background(), "_SMS_"+phone).Result()
	return value, err
}

func (jwtService *JwtService) DelRedisSmsAuth(phone string) (err error) {
	err = global.SYS_REDIS.Del(context.Background(), "_SMS_"+phone).Err()
	return err
}

func (jwtService *JwtService) SetRedisCityConfig(val string) (err error) {
	if err != nil {
		return err
	}
	err = global.SYS_REDIS.Set(context.Background(), "_city_config_", val, 0).Err()
	return err
}

func (jwtService *JwtService) GetRedisCityConfig() (value string, err error) {
	value, err = global.SYS_REDIS.Get(context.Background(), "_city_config_").Result()
	return value, err
}

func (jwtService *JwtService) SetRedisIndustryConfig(val string) (err error) {
	if err != nil {
		return err
	}
	err = global.SYS_REDIS.Set(context.Background(), "_industry_config_", val, 0).Err()
	return err
}

func (jwtService *JwtService) GetRedisIndustryConfig() (value string, err error) {
	value, err = global.SYS_REDIS.Get(context.Background(), "_industry_config_").Result()
	return value, err
}

func (jwtService *JwtService) SetRedisJobConfig(val string) (err error) {
	if err != nil {
		return err
	}
	err = global.SYS_REDIS.Set(context.Background(), "_job_config_", val, 0).Err()
	return err
}

func (jwtService *JwtService) GetRedisJobConfig() (value string, err error) {
	value, err = global.SYS_REDIS.Get(context.Background(), "_job_config_").Result()
	return value, err
}

func (jwtService *JwtService) GetSensitiveWord() (value string, err error) {
	value, err = global.SYS_REDIS.Get(context.Background(), "_sensitive_word_").Result()
	return value, err
}

func (jwtService *JwtService) SetSensitiveWord(val string) (err error) {
	if err != nil {
		return err
	}
	err = global.SYS_REDIS.Set(context.Background(), "_sensitive_word_", val, 0).Err()
	return err
}
