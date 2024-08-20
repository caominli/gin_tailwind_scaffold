package router

import (
	common "ginTailwindcssBase/commons"
	view "ginTailwindcssBase/views"

	// "github.com/gin-contrib/cors" //允许跨域，生产环境需注释
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode) //正式版模式
	r := gin.Default() //初始化gin

	// //允许跨域，生产环境需注释
	// r.Use(cors.New(cors.Config{ //生产环境需注释
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))
	r.LoadHTMLGlob("templates/*") //加载html模板全局，指定templates目录下的所有文件
	//生产环境需注释，因为使用反向代理来处理静态文件
	r.Static("/static", "./static") //静态文件处理,参数：路由，本地路径

	r.GET("/", common.IsLogin, view.Index)

	r.GET("/register", view.Register)
	r.POST("/register", view.Register)
	r.GET("/login", view.Login)
	r.POST("/login", view.Login)
	r.GET("/logout", view.Logout) //退出登录

	r.GET("/changepassword", view.ChangePassword)  //修改密码
	r.POST("/changepassword", view.ChangePassword) //修改密码

	r.GET("/acc", common.AuthUser, view.Acc) //必须登陆的路由

	api := r.Group("/api")
	{
		api.POST("/getcode", view.Getcode) //获取验证码
	}

	return r
}
