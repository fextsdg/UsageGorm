package dbUtil

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 用于新建Mysql数据库链接
func OpenConnect() *gorm.DB {
	url := "root:root@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名复数形式
		},
	})

	if err != nil {
		panic(fmt.Sprintf("连接数据库失败，错误为：%v\n", err))
	}
	return db
}
