package inits

import (
	"face-signIn/pkg/globals"
	"log"

	"github.com/spf13/viper"
)

// ServerInit 服务初始化
func ServerInit() {
	// viper 提取配置文件中的app 即提取ip和端口
	if err := viper.UnmarshalKey("app", &globals.AppConfig.App); err != nil {
		log.Fatalf("无法解码为结构: %s", err)
	}
}
