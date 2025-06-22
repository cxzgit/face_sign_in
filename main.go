package main

import (
	_ "face-signIn/inits" // 空导入，初始化
	"face-signIn/internal/server"
)

func main() {
	server.Run()
}
