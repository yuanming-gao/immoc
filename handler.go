/**
* @Time : 2019/11/15 9:06 下午
* @Author : GaoYuanMing
* @Package : 超星学习通
* @FileName : handler.go
 */
package main

import (
	"awesomeProject/comm"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PageData struct {
	Result string
}

var Data = new(PageData)

func homeHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", gin.H{
		"result": Data.Result,
	})
}

func receiveHandler(ctx *gin.Context) {
	question := ctx.PostForm("question")
	//敏感词过滤
	question = comm.FilterSensitiveWord(question)
	url := "http://zhannei.baidu.com/cse/search?s=10823138076610196716&entry=1&q=" + question

	//1.获取最佳回答的连接URL
	//test pass!
	answerURL := regularParse(url, FetchAnswerExpression)
	//2.根据最佳回答的URL获取结果
	//test pass!
	questionResult := regularParseGBString(answerURL, FetchResultExpression)
	//3.返回最终答案
	//test pass!
	Data.Result = questionResult

	ctx.Redirect(http.StatusMovedPermanently, "/")
}
