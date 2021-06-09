package model

type HotRespond struct {
	ReportTime string    `json:"report_time"`
	List       []HotItem `json:"lists"`
}

type HotItem struct {
	// 标题
	Title string `json:"title"`
	// 链接
	Url string `json:"url"`
	// 热度
	Popular string `json:"popular"`
	// 趋势
	Trend string `json:"trend,omitempty"`
	// 作者
	Author string `json:"author,omitempty"`
	// 附加
	Extra string `json:"extra"`
}
