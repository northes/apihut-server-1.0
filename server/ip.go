package server

import (
	"apihut-server/model"
	"apihut-server/repository/mysql"
	"apihut-server/util"
	"errors"
	"fmt"
	"net"
)

var (
	ErrIPFormat = errors.New("IP格式错误")
)

// GetIPInfo 获取IP信息
func GetIPInfo(ip string) (ipInfo *model.IP, err error) {
	ipInfo = new(model.IP)

	address, version := util.ParseIP(ip)
	if address == nil || version == 0 {
		return nil, ErrIPFormat
	}

	ipInfo, err = mysql.GetIP(ip)
	if err != nil && err == mysql.ErrNotExist {
		fmt.Println("从网络获取")
		ipInfo, err = getIPInfoFromGaode(address, version)
		_ = mysql.CreatIP(ipInfo)
	}

	return ipInfo, nil
}

func getIPInfoFromGaode(address net.IP, version int) (ipInfo *model.IP, err error) {
	ipInfo = new(model.IP)

	err = util.HttpGetRequest(fmt.Sprintf(
		"https://restapi.amap.com/v5/ip?parameters&key=b9bd34580b7133934c40a831703cc3fb&ip=%s&type=%d",
		address.String(),
		version,
	), &ipInfo)
	if err != nil {
		return nil, err
	}

	if ipInfo.Status == "0" {
		return nil, util.ErrAPIRequest
	}

	ipInfo.Status, ipInfo.Info, ipInfo.Infocode = "", "", ""

	return ipInfo, nil
}
