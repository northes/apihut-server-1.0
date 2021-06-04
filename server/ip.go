package server

import (
	"apihut-server/model"
	"apihut-server/util"
	"errors"
	"fmt"
)

// GetIPInfo 获取IP信息
func GetIPInfo(i *model.IP) (ipInfo *model.GaoDeIP, err error) {
	ipInfo = new(model.GaoDeIP)

	address, version := util.ParseIP(i.IP)
	if address == nil || version == 0 {
		return nil, errors.New("IP格式错误")
	}

	err = util.HttpGetRequest(fmt.Sprintf(
		"https://restapi.amap.com/v5/ip?parameters&key=b9bd34580b7133934c40a831703cc3fb&ip=%s&type=%d",
		address.String(),
		version,
	), &ipInfo)
	if err != nil {
		return nil, err
	}

	if ipInfo.Status == "0" {
		return nil, util.ErrRequest
	}

	ipInfo.Status, ipInfo.Info, ipInfo.Infocode = "", "", ""

	return ipInfo, nil
}
