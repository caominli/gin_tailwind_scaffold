package controller

import (
	model "btcmai/models"
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

// // 根据链接自动加语言代码
// func Jump(c *gin.Context) {
// 	//判断url是否存在
// 	ifurl := c.Query("url")
// 	if ifurl == "" {
// 		c.JSON(400, gin.H{"error": "url Query is null"})
// 		c.Abort()
// 		return
// 	}
// 	//取上下文
// 	localizer, ok := c.MustGet("localizer").(*i18n.Localizer)
// 	if !ok {
// 		localizer = &i18n.Localizer{}
// 		log.Print("Jump视图函数i18n中间件返回了空")
// 	}
// 	lang := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "LangTag"})

// 	url := "/" + lang + ifurl
// 	c.Redirect(302, url) //跳转

// }
