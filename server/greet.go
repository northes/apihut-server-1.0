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
	d := getDayCode()
	if g.Day == 0 {
		g.Day = d
	}
	// 查数据库
	greetList, err := mysql.GetGreet(g)
	if err != nil {
		if err != mysql.ErrNotExist || g.NoRange {
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
	case hour < 18:
		t = constant.TimeAfternoon
	case hour < 21:
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
		timeTextArr = []string{"晚安"}
	}
	// 返回随机内容
	return timeTextArr[util.GetRange(len(timeTextArr))]
}

// 获取星期代码
func getDayCode() constant.DayCode {
	day := time.Now().Day() + 1
	switch day {
	case 1:
		return constant.DayMonday
	case 2:
		return constant.DayTuesday
	case 3:
		return constant.DayWednesday
	case 4:
		return constant.DayThursday
	case 5:
		return constant.DayFriday
	case 6:
		return constant.DaySaturday
	case 7:
		return constant.DaySunday
	default:
		return constant.DayDefault
	}
}
