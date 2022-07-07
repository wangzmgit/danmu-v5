package main

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/viper"

	"kuukaa.fun/danmu-v5/common"
	"kuukaa.fun/danmu-v5/routes"
	"kuukaa.fun/danmu-v5/util"
)

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)

func main() {
	initConfig()
	util.CheckFolder() // 检测文件夹是否存在
	println("     _                               _____  _____")
	println("    | |                             |  ___||  _  |")
	println("  __| | __ _ _ __  _ __ ___  _   _  |___ \\ | |/' |")
	println(" / _` |/ _` | '_ \\| '_ ` _ \\| | | |     \\ \\|  /| |")
	println("| (_| | (_| | | | | | | | | | |_| | /\\__/ /\\ |_/ /")
	println("\\__,_|\\__,_|_| |_|_| |_| |_|\\__,_| \\____(_)\\___/ ")
	println("\tversion:" + common.Version)
	//初始化Redis
	common.InitCache()
	//初始化数据库
	common.InitDB()
	// 设置模式
	gin.SetMode(ReleaseMode)
	// 创建gin日志文件
	ginLog := initGinLog()
	//禁用控制台颜色
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(ginLog)
	r := gin.Default()
	r = routes.CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

//初始化配置文件
func initConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

//初始化gin日志
func initGinLog() *rotatelogs.RotateLogs {
	writer, _ := rotatelogs.New(
		"./file/logs/gin_%Y%m%d%H.log",
		//日志保存60天
		rotatelogs.WithMaxAge(time.Hour*24*60),
		//这里设置12小时产生一个日志文件
		rotatelogs.WithRotationTime(time.Hour*12),
	)

	return writer
}
