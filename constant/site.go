package constant

type SiteName string

const (
	Baidu    SiteName = "baidu"    // 百度
	Weibo    SiteName = "weibo"    // 微博
	ThePaper SiteName = "thepaper" // 澎湃新闻
	Zhihu    SiteName = "zhihu"    // 知乎
	Bilibili SiteName = "bilibili" // B站
	ITHome   SiteName = "ithome"   // it之家
)

var rankUrlMap = map[SiteName]string{
	Baidu:    "http://top.baidu.com/buzz?b=1&fr=topindex",
	Weibo:    "https://s.weibo.com/top/summary?cate=realtimehot",
	ThePaper: "https://www.thepaper.cn/",
	Zhihu:    "https://www.zhihu.com/billboard",
	Bilibili: "https://www.bilibili.com/v/popular/rank/all",
	ITHome:   "https://www.ithome.com/",
}

const (
	SinaHrefUrl = "https://s.weibo.com"
)

func (s SiteName) RankUrl() string {
	return rankUrlMap[s]
}

func (s SiteName) Text() string {
	return string(s)
}

func (s SiteName) IsSiteName() bool {
	siteNameList := [6]string{"baidu", "weibo", "thepaper", "zhihu", "bilibili", "ithome"}
	for _, name := range siteNameList {
		if name == string(s) {
			return true
		}
	}
	return false
}
