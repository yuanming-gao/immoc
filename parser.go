/**
* @Time : 2019/11/16 11:58 上午
* @Author : GaoYuanMing
* @Package : main
* @FileName : reguler.go
 */
package main

import (
	"awesomeProject/comm"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const (
	FetchAnswerExpression = `<a rpos="" cpos="title" href="([^\"]+)"`
	FetchResultExpression = `正确答案：(.+)</p>`
)

//获取内容
func fetcherContent(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer response.Body.Close()
	return data
}

func regularParse(url string, expression string) string {
	reg := regexp.MustCompile(expression)
	matches := reg.FindSubmatch(fetcherContent(url))
	return string(matches[1])
}

func regularParseGBString(url string, expression string) string {
	reg := regexp.MustCompile(expression)
	data := comm.ConvertToUTF8String(fetcherContent(url), "gbk")
	matches := reg.FindSubmatch(data)
	return string(matches[1])
}
