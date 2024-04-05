package router

import (
	controller "btcmai/controllers"
	"github.com/gin-gonic/gin" 
)

// type Menu struct {
//     LangTag     string
//     HelloPerson    string
// }

func Router() *gin.Engine {
	r := gin.Default()              //初始化gin

	r.LoadHTMLGlob("templates/*")   //加载html模板全局，指定templates目录下的所有文件
	r.Static("/static", "./static") //静态文件处理,参数：路由，本地路径

	r.GET("/", controller.Index)

	// r.GET("/class/:page/:name", controllers.Param)

	// api := r.Group("/api")
	// {
	// 	// api.GET("/json", controllers.Json)             //调用json结构体

	// }
	return r
}
