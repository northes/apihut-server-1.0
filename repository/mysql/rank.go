package mysql

import (
	"apihut-server/constant"
	"apihut-server/model"
)

func CreateRank(hot *model.Rank) (err error) {
	affected, err := engine.Insert(hot)
	if err != nil {
		return err
	}
	if affected < 1 {
		return ErrCreat
	}

	return nil
}

func GetRank(siteName constant.SiteName) (hot *model.Rank, err error) {
	hot = new(model.Rank)
	has, err := engine.Where("site_name=?", siteName).Desc("created_time").Get(hot)
	if !has {
		return nil, ErrNotExist
	}
	return hot, nil
}
