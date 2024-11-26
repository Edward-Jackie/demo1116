package db

import (
	cfg "demo1116/config"
	"demo1116/logger"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var Db *gorm.DB

func initMysql() {
	c := cfg.Cfg.Mysql
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DataBase,
	)

	var err error
	// 连接
	Db, err = gorm.Open("mysql", dsn)
	if err != nil {
		logger.GLog.Error("数据库连接发生错误 err =", zap.Error(err))
	} else {
		logger.GLog.Info("数据库连接成功!")
	}

	// 打开GORM日志模式
	Db.LogMode(true)
	// 指定日志打印器
	Db.SetLogger(gorm.Logger{logger.GormLogger{}})
}
