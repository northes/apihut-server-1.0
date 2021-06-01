package server

import (
	"apihut-server/constant"
	"apihut-server/model"
	"apihut-server/repository/mysql"
	"apihut-server/util"
	"time"
)

func Greet() (*model.GreetRespond, error) {
	t := getTimeCode()
	greetList, err := mysql.GetGreet(t)
	if err != nil {
		return nil, err
	}

	greet := greetList[util.GetRange(len(greetList))]
	return &model.GreetRespond{Words: getTimeText(t), Sentence: greet.Sentence, Author: greet.Author}, err
}

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

func getTimeText(code constant.TimeCode) string {
	var timeTextArr []string
	timeTextArr = []string{"Hello"}
	switch code {
	case constant.TimeEarlyMorning:
		timeTextArr = []string{"凌晨啦"}
	case constant.TimeMorning:
		timeTextArr = []string{"早上好", "早安", "早~"}
	case constant.TimeNoon:
		timeTextArr = []string{"中午好"}
	case constant.TimeAfternoon:
		timeTextArr = []string{"下午好"}
	case constant.TimeEvening:
		timeTextArr = []string{"傍晚好"}
	case constant.TimeNight:
		timeTextArr = []string{"晚上好"}
	case constant.TimeLateNight:
		timeTextArr = []string{"夜深了", "晚安", "好梦"}
	}
	// 返回随机内容
	return timeTextArr[util.GetRange(len(timeTextArr))]
}
