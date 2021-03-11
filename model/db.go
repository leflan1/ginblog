package model

import (
	"fmt"
	"ginblog/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	//数据库连接文件
	db, err := gorm.Open(utils.Db, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))
	if err != nil {
		fmt.Printf("连接数据库失败，请检查参数 err: %s\n", err)
	}

	//禁用默认表名的复数形式
	db.SingularTable(true)

	//模型迁入,将&User{}, &Article{}, &Category{}传入
	//AutoMigrate用于自动迁移多个model
	db.AutoMigrate(&User{}, &Article{}, &Category{})
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。----不要大于设置框架timeout时间
	db.DB().SetConnMaxLifetime(10 * time.Second)

	//关闭数据库‘
	//db.Close()

}
