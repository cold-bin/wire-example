package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
	"time"
	"wire-example/conf"
)

var (
	gdb  *gorm.DB
	once = &sync.Once{}
)

// 单例模式
func DBClient() *gorm.DB {
	if gdb != nil {
		return gdb
	}

	once.Do(func() {
		source := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.MysqlUsername,
			conf.MysqlPassword,
			conf.MysqlHost,
			conf.MysqlPort,
			conf.MysqlDbname,
		)
		db_, err := gorm.Open(mysql.Open(source), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
		if err != nil {
			panic(err)
		}
		db, err := db_.DB()
		if err != nil {
			panic(err)
		}
		db.SetMaxOpenConns(conf.MysqlMaxOpenConns)
		db.SetMaxIdleConns(conf.MysqlMaxIdleConns)
		db.SetConnMaxLifetime(conf.MysqlConnMaxLifeMinutes * time.Minute)
		err = db_.Migrator().AutoMigrate()
		if err != nil {
			panic(err)
		}
		gdb = db_
	})
	return gdb
}
