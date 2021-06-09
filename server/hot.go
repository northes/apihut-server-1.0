package server

import (
	"apihut-server/constant"
	"apihut-server/model"
	"fmt"
	"strings"

	"github.com/gocolly/colly/proxy"

	"golang.org/x/text/encoding/simplifiedchinese"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func GetHot(site string) ([]model.HotItem, error) {
	// 获取代理IP
	proxyIP, err := GetProxyIP()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(proxyIP)

	switch site {
	case constant.SiteNameBaidu:
		return GetBaiduHot(&proxyIP)
	case constant.SiteNameSina:
		return GetSinaHot(&proxyIP)
	case constant.SiteNameThePaper:
		return GetThePaperHot(&proxyIP)
	default:
		return GetBaiduHot(&proxyIP)
	}

	//return GetSinaHot(&proxyIP)
}

func GetBaiduHot(proxyIP *string) (hotList []model.HotItem, err error) {

	c := GetColly(proxyIP)

	c.OnHTML(".list-table tr:not(.item-tr)", func(e *colly.HTMLElement) {
		title, _ := simplifiedchinese.GBK.NewDecoder().Bytes([]byte(string(e.ChildText(".keyword .list-title"))))
		url := e.ChildAttr(".keyword .list-title", "href")
		popular := e.ChildText(".last span")
		trendName := e.ChildAttr(".last span", "class")

		var trend string
		if trendName == "icon-rise" {
			trend = "rise"
		} else {
			trend = "fall"
		}

		// 去除空格
		nok := strings.Replace(string(title), " ", "", -1)
		// 去除换行
		non := strings.Replace(nok, "\n", "", -1)

		hotList = append(hotList, model.HotItem{
			Title:   non,
			Url:     url,
			Popular: popular,
			Extra:   trend,
		})

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(constant.BaiduHotUrl)
	//c.Visit("http://baidu.apihut.net/")

	//fmt.Println(hotList[1:])

	return hotList[1:], err
}

func GetSinaHot(proxyIP *string) (hotList []model.HotItem, err error) {
	c := GetColly(proxyIP)

	c.OnHTML("tbody tr ", func(e *colly.HTMLElement) {
		title := e.ChildText(".td-02 a")
		href := e.ChildAttr(".td-02 a", "href")
		hrefTo := e.ChildAttr(".td-02 a", "href_to")
		popular := e.ChildText(".td-02 span")
		tag := e.ChildText(".td-03 i")

		var url string
		if len(hrefTo) != 0 {
			url = constant.SinaHrefUrl + hrefTo
		} else {
			url = constant.SinaHrefUrl + href
		}

		hotList = append(hotList, model.HotItem{
			Title:   title,
			Url:     url,
			Popular: popular,
			Extra:   tag,
		})
	})

	c.Visit(constant.SinaHotUrl)

	return hotList, err
}

func GetThePaperHot(proxyIP *string) (hotList []model.HotItem, err error) {

	c := GetColly(proxyIP)

	c.OnHTML("#listhot0 li:not(.list_more)", func(e *colly.HTMLElement) {
		title := e.ChildText("a")
		href := e.ChildAttr("a", "href")
		//fmt.Printf("Title:%s \n URL: https://www.thepaper.cn/%s \n", title, href)

		hotList = append(hotList, model.HotItem{
			Title: title,
			Url:   "https://www.thepaper.cn/" + href,
		})
	})

	c.Visit(constant.ThePaperHotUrl)
	return hotList, nil
}

// GetColly 返回Colly收集器
func GetColly(proxyIP *string) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
		colly.MaxDepth(1),
	)

	//设置代理IP
	if p, err := proxy.RoundRobinProxySwitcher(
		*proxyIP,
	); err == nil {
		c.SetProxyFunc(p)
	}

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	return c
}
