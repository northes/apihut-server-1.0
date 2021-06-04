package model

type IP struct {
	IP string `json:"ip" form:"ip"`
}

type GaoDeIP struct {
	Status   string `json:"status,omitempty"`
	Info     string `json:"info,omitempty"`
	Infocode string `json:"infocode,omitempty"`
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	ISP      string `json:"isp"`
	Location string `json:"location"`
	IP       string `json:"ip"`
}
