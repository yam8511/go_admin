package main

import (
	"GO_Admin/global"
	route "GO_Admin/router"
	"log"
)

func init() {
	// 顯示日期 ｜ 顯示時間（台灣） ｜ 顯示檔案名稱
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	global.Start()

	r := route.SetupRouter()

	// Listen and Server in 0.0.0.0:8081
	r.Run(":8081")
}
