package model

import "time"

type IP struct {
	ID        int       `json:"-" xorm:"pk autoincr notnull comment(id)"`
	Status    string    `json:"status,omitempty" xorm:"-"`
	Info      string    `json:"info,omitempty" xorm:"-"`
	InfoCode  string    `json:"infocode,omitempty" xorm:"-"`
	IP        string    `json:"ip" xorm:"index pk notnull comment(IP)"`
	Country   string    `json:"country"  xorm:"comment(国家)"`
	Province  string    `json:"province" xorm:"comment(省份)"`
	City      string    `json:"city" xorm:"comment(城市)"`
	District  string    `json:"district" xorm:"comment(地区)"`
	ISP       string    `json:"isp" xorm:"comment(运营商)"`
	Location  string    `json:"location" xorm:"comment(经纬度)"`
	CreatedAt time.Time `json:"-" xorm:"created comment(创建时间)"`
	UpdatedAt time.Time `json:"-" xorm:"updated comment(更新时间)"`
}
