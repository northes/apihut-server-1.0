package mysql

import (
	"apihut-server/constant"
	"apihut-server/model"
)

func GetGreet(time constant.TimeCode) (greetList []model.Greet, err error) {
	greetList = make([]model.Greet, 0)
	err = engine.Where("time=?", time).
		Find(&greetList)
	if err != nil {
		return nil, err
	}
	return
}
