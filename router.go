/**
* @Time : 2019/11/15 9:11 下午
* @Author : GaoYuanMing
* @Package : main
* @FileName : router.go
 */
package main

import "github.com/gin-gonic/gin"

func registerRouters(router *gin.Engine) {

	router.GET("/", homeHandler)

	router.POST("/receive", receiveHandler)
}
