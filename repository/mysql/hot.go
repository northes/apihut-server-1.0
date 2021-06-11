package mysql

import (
	"apihut-server/model"
	"errors"
)

func CreateHot(hot *model.Hot) (err error) {
	affected, err := engine.Insert(hot)
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("创建失败")
	}

	return nil
}

func GetHot(siteName string) (hot *model.Hot, err error) {
	hot = new(model.Hot)
	has, err := engine.Where("site_name=?", siteName).Desc("created_time").Get(hot)
	if !has {
		return nil, errors.New("不存在")
	}
	return hot, nil
}
