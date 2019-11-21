/**
* @Time : 2019/11/16 11:52 上午
* @Author : GaoYuanMing
* @Package : comm
* @FileName : filter.go
 */
package comm

import (
	"github.com/axgle/mahonia"
	"strings"
)

func FilterSensitiveWord(word string) string {
	SensitiveWord := []string{"习近平", "江泽民", "胡锦涛"}
	var s = word
	for i := 0; i < len(SensitiveWord); i++ {
		s = strings.Replace(s, SensitiveWord[i], "", -1)
	}
	return s
}

//字符转码
func ConvertToUTF8String(data []byte, code string) []byte {
	srcCoder := mahonia.NewDecoder(code)
	srcResult := srcCoder.ConvertString(string(data))
	tagCoder := mahonia.NewDecoder("utf-8")
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}
