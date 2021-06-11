package mysql

import (
	"apihut-server/model"
	"errors"
)

func CreateRank(hot *model.Rank) (err error) {
	affected, err := engine.Insert(hot)
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("创建失败")
	}

	return nil
}

func GetRank(siteName string) (hot *model.Rank, err error) {
	hot = new(model.Rank)
	has, err := engine.Where("site_name=?", siteName).Desc("created_time").Get(hot)
	if !has {
		return nil, errors.New("不存在")
	}
	return hot, nil
}
