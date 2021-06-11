package server

import (
	"apihut-server/config"
	"apihut-server/constant"
	"apihut-server/model"
	"apihut-server/repository/mysql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/proxy"

	"golang.org/x/text/encoding/simplifiedchinese"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

var (
	myProxyIP              string
	ErrDataUpdateMore1Hour = errors.New("热榜数据更新间隔超过1小时")
)

func GetHot(site string) (hotList []model.HotItem, err error) {
	// 获取代理IP
	myProxyIP, err = GetProxyIP()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	hot, err := getFromLocalCache(site)
	if err != nil {
		// 本地获取失败则在线抓取
		switch site {
		case constant.SiteNameBaidu:
			return getBaiduHot()
		case constant.SiteNameSina:
			return getSinaHot()
		case constant.SiteNameThePaper:
			return getThePaperHot()
		case constant.SiteNameZhihu:
			return getZhihuHot()
		case constant.SiteNameBilibili:
			return getBiliBiliHot()
		case constant.SiteNameBilibiliShort:
			return getBiliBiliHot()
		default:
			return getBaiduHot()
		}
	}
	// 返回本地缓存数据
	return hot.HotList, nil
}

// 获取百度热榜
func getBaiduHot() (hotList []model.HotItem, err error) {

	c := getColly()

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
			Trend:   trend,
		})
	})

	c.Visit(constant.BaiduHotUrl)
	//c.Visit("http://baidu.apihut.net/")

	return hotList[1:], err
}

// 获取微博热搜
func getSinaHot() (hotList []model.HotItem, err error) {
	c := getColly()

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

// 获取澎湃新闻热闻
func getThePaperHot() (hotList []model.HotItem, err error) {

	c := getColly()

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

// 获取知乎热榜
func getZhihuHot() (hotList []model.HotItem, err error) {
	var zhihu model.ZhihuHot
	c := getColly()
	c.OnHTML("#js-initialData", func(e *colly.HTMLElement) {
		//fmt.Println(e.Text)
		err := json.Unmarshal([]byte(e.Text), &zhihu)
		if err != nil {
			fmt.Println(err.Error())
		}
	})

	c.Visit(constant.ZhihuHotUrl)

	zhihuList := zhihu.InitialState.Topstory.HotList
	for i := 0; i < len(zhihuList); i++ {
		item := zhihuList[i].Target
		hotList = append(hotList, model.HotItem{
			Title:   item.TitleArea.Text,
			Url:     item.Link.URL,
			Popular: item.MetricsArea.Text,
			Extra:   item.ExcerptArea.Text,
		})
	}

	return hotList, nil
}

// 获取Bilibili热榜
func getBiliBiliHot() (hotList []model.HotItem, err error) {
	c := getColly()
	c.OnHTML(".rank-list li", func(e *colly.HTMLElement) {
		title := e.ChildText(".info .title")
		href := e.ChildAttr(".info .title", "href")
		pts := e.ChildText(".info .pts div")
		author := e.ChildText(".info .detail .up-name")
		play := e.ChildText(".info .detail > span:first-child")

		hotList = append(hotList, model.HotItem{
			Title:   title,
			Url:     "https:" + href,
			Popular: pts,
			Author:  author,
			Extra:   play,
		})
	})

	c.Visit(constant.BilibiliHotUrl)

	return hotList, nil
}

// 返回Colly收集器
func getColly() *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
		colly.MaxDepth(1),
	)

	//设置代理IP
	if p, err := proxy.RoundRobinProxySwitcher(
		myProxyIP,
	); err == nil && config.Conf.Mode == gin.ReleaseMode {
		c.SetProxyFunc(p)
		fmt.Println("Use ProxyIP: " + myProxyIP)
	}

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	return c
}

func getFromLocalCache(siteName string) (hot *model.Hot, err error) {
	hot, err = mysql.GetHot(siteName)
	if err != nil {
		return nil, err
	}
	// 判断热榜数据是否超过一个小时
	if time.Now().Sub(hot.CreatedTime).Hours() >= 1 {
		// 刷新全部数据
		return hot, ErrDataUpdateMore1Hour
	}
	return hot, nil
}
