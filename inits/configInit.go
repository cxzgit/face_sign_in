package inits

import (
	"face-signIn/pkg/globals"
	"fmt"
	"github.com/spf13/viper"
)

// ConfigInit 初始化配置文件
func ConfigInit() {
	fmt.Println("当前项目环境为", globals.Env)

	viper.SetConfigName(globals.Env) // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")      // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("./configs") // 查找配置文件所在的路径

	viper.AddConfigPath(".")    // 还可以在工作目录中查找配置
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		globals.Log.Panicf("Fatal errors config file: %s \n", err)
	}
}
