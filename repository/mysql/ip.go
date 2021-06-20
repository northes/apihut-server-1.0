package mysql

import "apihut-server/model"

func CreatIP(i *model.IP) (err error) {
	affected, err := engine.Insert(i)
	if err != nil {
		return err
	}
	if affected < 1 {
		return ErrCreat
	}
	return nil
}

func GetIP(ip string) (ipInfo *model.IP, err error) {
	ipInfo = new(model.IP)
	has, err := engine.Where("ip=?", ip).Get(ipInfo)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, ErrNotExist
	}
	return
}
