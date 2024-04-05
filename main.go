package main

import (
	model "btcmai/models"
	router "btcmai/routers"

)


func main() {

	r := router.Router()
	r.Run(":"+model.Port) // 启动HTTP服务器，监听端口80
}