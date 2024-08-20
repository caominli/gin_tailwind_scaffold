package view

import (
	model "ginTailwindcssBase/models"
	// "log"
	// service "btcmai/services"
	// "gorm.io/gorm"
	// "log"
	// "time"

	"github.com/gin-gonic/gin" //导入gin包
)

// 主页
func Index(c *gin.Context) {
	useremail := c.MustGet("useremail") //传入是否登录信息，在路由那里使用了中间件

	c.HTML(200, "index.html", gin.H{"title": "测试首页","user":useremail})
}

// 我的账户
func Acc(c *gin.Context) {
	useremail := c.MustGet("useremail")
	
	msg := model.FlashGet(c)
	c.HTML(200, "acc.html", gin.H{"title": "我的账户", "msg": msg,"user":useremail})
}

