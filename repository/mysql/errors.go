package mysql

import "errors"

var (
	ErrCreat    = errors.New("创建失败")
	ErrNotExist = errors.New("不存在")
)
