package server

import (
	"apihut-server/model"
	"apihut-server/repository/redis"
	"apihut-server/util"
	"fmt"
	"strconv"
)

// GetProxyIP 获取代理IP
func GetProxyIP() (proxyIP string, err error) {
	proxyIP, err = redis.GetProxyIP()
	if err != nil && err == redis.ErrValueNotExit {
		proxyIP, err = GetProxyIPFromKuaiNiao()
		if len(proxyIP) != 0 {
			err = redis.SetProxyIP(proxyIP)
		}
		if err != nil {
			return "", err
		}
	}
	return
}

// GetProxyIPFromKuaiNiao 从快鸟代理IP网获取IP
// http://www.kuainiaoip.com/
func GetProxyIPFromKuaiNiao() (proxyIP string, err error) {
	var p model.ProxyIP
	err = util.HttpGetRequest(
		"http://api.kuainiaoip.com/index.php?fetch_type=2021062321322782921&pool_id=&qty=1&time=101&province=%E6%89%80%E6%9C%89&city=%E6%89%80%E6%9C%89&protocol=1&format=json&dt=1",
		&p,
	)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if p.Code != 0 || !p.Success {
		return "", util.ErrAPIRequest
	}

	proxyIP = fmt.Sprintf("http://%s:%s", p.Data[0].IP, strconv.FormatInt(p.Data[0].Port, 10))

	return proxyIP, err
}
