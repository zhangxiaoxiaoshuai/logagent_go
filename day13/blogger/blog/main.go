package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"

	"code.oldboy.com/studygolang/day12/blogger/controller"

	"code.oldboy.com/studygolang/day12/blogger/db"
	"github.com/gin-gonic/gin"
	"github.com/DeanThompson/ginpprof"
)

// blogger 个人博客程序

var log = logrus.New()

func main() {
	// 0.初始化日志
	f, err := os.OpenFile("./log/gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return
	}
	log.Out = f
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Out
	// Only log the warning severity or above.
	log.Level = logrus.DebugLevel

	// 1.初始化数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_web?parseTime=true"
	err = db.Init(dsn)
	if err != nil {
		panic(err)
	}

	// 生成默认的路由处理引擎
	r := gin.Default()
	// 加载静态文件
	r.Static("/static", "./static")
	// 加载模板文件
	r.LoadHTMLGlob("views/*")

	// 路由
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "views/404.html", nil)
	})

	// 首页
	r.GET("/", controller.IndexHandler)
	ginpprof.Wrap(r)
	// 启动
	err = r.Run()
	if err != nil {
		panic(err)
	}
}
