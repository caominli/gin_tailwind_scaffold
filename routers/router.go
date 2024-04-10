package router

import (
	controller "btcmai/controllers"
	service "btcmai/services"
	"github.com/gin-gonic/gin"
)

// type Menu struct {
//     LangTag     string
//     HelloPerson    string
// }

func Router() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode) //正式版模式
	r := gin.Default() //初始化gin

	r.LoadHTMLGlob("templates/*")   //加载html模板全局，指定templates目录下的所有文件
	r.Static("/static", "./static") //静态文件处理,参数：路由，本地路径

	r.GET("/", service.IsLogin,controller.Index)

	r.GET("/register", controller.Register)
	r.POST("/register", controller.Register)
	r.GET("/login", controller.Login)
	r.POST("/login", controller.Login)
	r.GET("/logout", controller.Logout) //退出登录

	r.GET("/changepassword", controller.ChangePassword) //修改密码
	r.POST("/changepassword", controller.ChangePassword) //修改密码

	// r.GET("/class/:page/:name", controllers.Param)

	r.GET("/acc", service.AuthUser, controller.Acc) //必须登陆的路由

	api := r.Group("/api")
	{
		api.POST("/getcode", controller.Getcode) //获取验证码
	}

	return r
}
