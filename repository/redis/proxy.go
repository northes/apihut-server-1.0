package redis

import (
	"time"
)

// SetProxyIP 设置代理IP
func SetProxyIP(ip string) (err error) {
	err = rdb.Set(KeyPrefix+KeyProxyIP, ip, 5*time.Hour).Err()
	return
}

// GetProxyIP 获取代理IP
func GetProxyIP() (ip string, err error) {
	ip, err = rdb.Get(KeyPrefix + KeyProxyIP).Result()
	if err != nil && len(ip) == 0 {
		return "", ErrValueNotExit
	}
	return
}
