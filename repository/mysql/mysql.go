package mysql

import (
	"apihut-server/config"
	"apihut-server/model"
	"fmt"
	"time"

	"xorm.io/xorm/names"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func Init(cfg *config.MySQLConfig) (err error) {
	//var err error
	engine, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	))
	if err != nil {
		fmt.Println("connect failed,err:", err.Error())
	}
	//	如果需要设置最大打开连接数
	engine.SetMaxOpenConns(config.Conf.MySQLConfig.MaxOpenConns)
	//	如果需要设置连接池的空闲数大小，可以使用
	engine.SetMaxIdleConns(config.Conf.MySQLConfig.MaxIdleConns)
	//	如果需要设置连接的最大生存时间
	engine.SetConnMaxLifetime(time.Minute * 5)
	//	设置名称映射规则
	engine.SetMapper(names.GonicMapper{})
	// 设置数据库时区
	location, err := time.LoadLocation("Asia/Shanghai")
	engine.TZLocation = location
	// 连接数据库
	err = engine.Ping()
	if err != nil {
		fmt.Println("mysql failed,err:", err.Error())
		return err
	}
	// 同步数据库
	err = engine.Sync2(
		new(model.Greet),
	)
	if err != nil {
		fmt.Println("同步数据表结构错误", err.Error())
		return err
	}
	return nil
}

func Close() {
	err := engine.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
}
