package model

type IP struct {
	Status   string `json:"status,omitempty" xorm:"-"`
	Info     string `json:"info,omitempty" xorm:"-"`
	Infocode string `json:"infocode,omitempty" xorm:"-"`
	IP       string `json:"ip" xorm:"index notnull comment(IP)"`
	Country  string `json:"country"  xorm:"comment(国家)"`
	Province string `json:"province" xorm:"comment(省份)"`
	City     string `json:"city" xorm:"comment(城市)"`
	District string `json:"district" xorm:"comment(地区)"`
	ISP      string `json:"isp" xorm:"comment(运营商)"`
	Location string `json:"location" xorm:"comment(经纬度)"`
}
