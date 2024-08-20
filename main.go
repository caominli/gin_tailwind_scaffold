package main

import (
	router "ginTailwindcssBase/routers"
	config "ginTailwindcssBase/config"

)


func main() {

	r := router.Router()
	r.Run(":"+config.Config.Port) // 启动HTTP服务器，监听端口
}