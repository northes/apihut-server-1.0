package model

type ProxyIP struct {
	Code    int64   `json:"code"`
	Success bool    `json:"success"`
	Msg     string  `json:"msg"`
	Data    []Datum `json:"data"`
}

type Datum struct {
	IP         string `json:"ip"`
	Port       int64  `json:"port"`
	ISP        string `json:"isp"`
	Address    string `json:"address"`
	ExpireTime string `json:"expire_time"`
}
