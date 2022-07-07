package common

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kuukaa.fun/danmu-v5/model"
	"kuukaa.fun/danmu-v5/util"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		username,
		password,
		host,
		port,
		database)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		util.LogError("failed to connect database ,err:" + err.Error())
		return nil
	}
	// 连接池
	sqlDB, err := db.DB()
	if err != nil {
		util.LogError("connect db server failed, err:" + err.Error())
		return nil
	}
	sqlDB.SetMaxIdleConns(10)           // 设置连接池中空闲连接的最大数量
	sqlDB.SetMaxOpenConns(100)          //设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) //设置了连接可复用的最大时间
	//数据库迁移
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Partition{})
	db.AutoMigrate(&model.Video{})
	db.AutoMigrate(&model.Resource{})
	db.AutoMigrate(&model.Liked{})       //点赞表
	db.AutoMigrate(&model.Collect{})     //收藏表
	db.AutoMigrate(&model.Collection{})  //收藏夹表
	db.AutoMigrate(&model.Follow{})      //关注表
	db.AutoMigrate(&model.Comment{})     //评论表
	db.AutoMigrate(&model.Announce{})    //公告表
	db.AutoMigrate(&model.UserMessage{}) //私信表
	db.AutoMigrate(&model.Danmaku{})     //弹幕表
	db.AutoMigrate(&model.Carousel{})    //轮播图表

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
