package server

import (
	"apihut-server/model"
	"fmt"
	"strings"

	"github.com/gocolly/colly/proxy"

	"golang.org/x/text/encoding/simplifiedchinese"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func GetHot() ([]model.HotItem, error) {
	// 获取代理IP
	proxyIP, err := GetProxyIP()
	if err != nil {
		return nil, err
	}
	fmt.Println(proxyIP)

	return GetBaiduHot(&proxyIP)
}

func GetBaiduHot(proxyIP *string) (hotList []model.HotItem, err error) {

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
		colly.MaxDepth(1),
	)

	// 设置代理IP
	if p, err := proxy.RoundRobinProxySwitcher(
		*proxyIP,
	); err == nil {
		c.SetProxyFunc(p)
	}

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnHTML(".list-table tr:not(.item-tr)", func(e *colly.HTMLElement) {
		title, _ := simplifiedchinese.GBK.NewDecoder().Bytes([]byte(string(e.ChildText(".keyword .list-title"))))
		url := e.ChildAttr(".keyword .list-title", "href")
		num := e.ChildText(".last span")

		// 去除空格
		nok := strings.Replace(string(title), " ", "", -1)
		// 去除换行
		non := strings.Replace(nok, "\n", "", -1)

		hotList = append(hotList, model.HotItem{
			Title: non,
			Url:   url,
			Extra: num,
		})

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://top.baidu.com/buzz?b=1&fr=topindex")
	//c.Visit("http://baidu.apihut.net/")

	//fmt.Println(hotList[1:])

	return hotList[1:], err
}
