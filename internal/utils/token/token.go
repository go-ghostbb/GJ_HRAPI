package token

import (
	"fmt"
	gbredis "ghostbb.io/gb/database/gb_redis"
	gbctx "ghostbb.io/gb/os/gb_ctx"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"time"
)

const (
	DefaultExpiration time.Duration = 1 * time.Minute // default time, 1 min
	Prefix            string        = "jwt"
	whiteKey          string        = Prefix + "-white:%s"
	blackKey          string        = Prefix + "-black:%s"
	refreshKey        string        = Prefix + "-refresh:%s"
)

// SetWhite 設定白名單token
func SetWhite(r *gbredis.Redis, token string) error {
	return r.SetEX(gbctx.New(), fmt.Sprintf(whiteKey, token), true, 60)
}

// IsWhite 判斷是否為白名單
func IsWhite(r *gbredis.Redis, token string) (bool, error) {
	exists, err := r.Exists(gbctx.New(), fmt.Sprintf(whiteKey, token))
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

// SetBlack 設置黑名單token
func SetBlack(r *gbredis.Redis, token string) error {
	return r.SetEX(gbctx.New(), fmt.Sprintf(blackKey, token), true, 600)
}

// IsBlack 判斷是否為黑名單
func IsBlack(r *gbredis.Redis, token string) (bool, error) {
	exists, err := r.Exists(gbctx.New(), fmt.Sprintf(blackKey, token))
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

// SetRefreshToken 設置refresh token到redis緩存
func SetRefreshToken(r *gbredis.Redis, key string, token string, ep time.Duration) error {
	return r.SetEX(gbctx.New(), fmt.Sprintf(refreshKey, key), token, gbconv.Int64(ep.Seconds()))
}

// DelRefreshToken 刪除使用者refresh token
func DelRefreshToken(r *gbredis.Redis, key string) error {
	_, err := r.Del(gbctx.New(), fmt.Sprintf(refreshKey, key))
	return err
}

// GetRefreshToken 從redis獲取refresh token
func GetRefreshToken(r *gbredis.Redis, key string) (string, error) {
	result, err := r.Get(gbctx.New(), fmt.Sprintf(refreshKey, key))
	if err != nil {
		return "", err
	}
	return result.String(), nil
}
