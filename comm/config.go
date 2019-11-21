package comm

import (
	"log"

	"github.com/unknwon/goconfig"
)

var (
	config *goconfig.ConfigFile
)

func init() {
	var err error
	config, err = goconfig.LoadConfigFile("config.ini")
	if err != nil {
		panic(err)
	}
}

//GetConfigValue return a config Value
func GetConfigValue(key string) string {
	value, err := config.GetValue("", key)
	if err != nil {
		log.Println("读取配置文件失败", err)
		return ""
	}
	return value
}
