package server

import (
	"apihut-server/constant"
	"apihut-server/model"
	"apihut-server/repository/mysql"
	"apihut-server/util"
	"time"
)

// Greet 招呼
func Greet(g *model.Greet) (*model.GreetRespond, error) {
	// 参数校验
	t := getTimeCode()
	if g.Time == 0 {
		g.Time = t
	}
	// 查数据库
	greetList, err := mysql.GetGreet(g)
	if err != nil {
		if err != mysql.ErrNotExist {
			return nil, err
		}
		greetList, err = mysql.GetGreet(&model.Greet{Time: 0})
	}
	// 随机抽取一条数据
	if len(greetList) > 0 {
		*g = greetList[util.GetRange(len(greetList))]
	}
	return &model.GreetRespond{Words: getTimeText(t), Sentence: g.Sentence, Author: g.Author}, err
}

// 获取时间代码
func getTimeCode() constant.TimeCode {
	hour := time.Now().Hour()
	var t constant.TimeCode
	switch {
	case hour < 4:
		t = constant.TimeLateNight
	case hour < 6:
		t = constant.TimeEarlyMorning
	case hour < 11:
		t = constant.TimeMorning
	case hour < 13:
		t = constant.TimeNoon
	case hour < 16:
		t = constant.TimeAfternoon
	case hour < 20:
		t = constant.TimeEvening
	case hour < 24:
		t = constant.TimeNight
	default:
		t = constant.TimeDefault
	}
	return t
}

// 获取时间代码对应的文字
func getTimeText(code constant.TimeCode) string {
	var timeTextArr []string
	timeTextArr = []string{"Hello"}
	switch code {
	case constant.TimeEarlyMorning:
		timeTextArr = []string{"凌晨啦"}
	case constant.TimeMorning:
		timeTextArr = []string{"早上好", "早安", "早"}
	case constant.TimeNoon:
		timeTextArr = []string{"中午好"}
	case constant.TimeAfternoon:
		timeTextArr = []string{"下午好"}
	case constant.TimeEvening:
		timeTextArr = []string{"傍晚好"}
	case constant.TimeNight:
		timeTextArr = []string{"晚上好", "晚安"}
	case constant.TimeLateNight:
		timeTextArr = []string{"夜深了", "晚安", "好梦"}
	}
	// 返回随机内容
	return timeTextArr[util.GetRange(len(timeTextArr))]
}
