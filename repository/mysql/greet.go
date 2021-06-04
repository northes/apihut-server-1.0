package mysql

import (
	"apihut-server/model"
)

func GetGreet(g *model.Greet) (greetList []model.Greet, err error) {
	greetList = make([]model.Greet, 0)
	err = engine.Where("time=?", g.Time).
		Or("day=?", g.Day).
		Find(&greetList)
	if err != nil {
		return nil, err
	}
	if len(greetList) == 0 {
		return nil, ErrNotExist
	}
	return
}
