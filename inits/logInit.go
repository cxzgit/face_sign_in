package inits

import (
	"face-signIn/pkg/globals"
	"face-signIn/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LogInit() {
	if err := viper.UnmarshalKey("log", &globals.AppConfig.Log); err != nil {
		globals.Log.Panicf("无法解码为结构: %s", err)
	}

	logPath := globals.AppConfig.Log.LogPath
	appName := globals.AppConfig.Log.AppName

	writeSyncer := logger.GetLogWriter(logPath, appName)
	encoder := logger.GetEncoder()

	// 将日志输出到文件
	fileCore := zapcore.NewCore(encoder, writeSyncer, zapcore.ErrorLevel)

	//合并控制台输出和文件输出
	core := zapcore.NewTee(fileCore)

	log := zap.New(core, zap.AddCaller())

	globals.Log = log.Sugar()
}
