package model

type HotRespond struct {
	ReportTime string    `json:"report_time"`
	List       []HotItem `json:"lists"`
}

type HotItem struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Extra string `json:"extra"`
}
