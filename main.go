/**
* @Time : 2019/11/15 9:05 下午
* @Author : GaoYuanMing
* @Package : 超星学习通
* @FileName : main.go
 */
package main

import (
	"awesomeProject/comm"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//视图层渲染
	router.LoadHTMLGlob("view/*")
	//加载静态文件/static
	router.Static("/static", "./static")
	//注册路由
	registerRouters(router)
	//绑定端口启动
	router.Run(comm.GetConfigValue("ServerAddress")).Error()
}
