package controller

import (
	model "btcmai/models"
	service "btcmai/services"
	"gorm.io/gorm"
	"log"
	"time"

	"github.com/gin-gonic/gin"         //导入gin包
	"github.com/gin-gonic/gin/binding" //绑定表单
)

// 注册
func Register(c *gin.Context) {
	// 获取当前请求的HTTP方法
	method := c.Request.Method
	//如果是post
	if method == "POST" {
		//先检查表单错误
		var form model.Register //定义一个表单变量

		err := c.ShouldBindWith(&form, binding.Form) //执行绑定
		if err != nil {                              //如果验证失败
			//创建消息
			model.FlashSet(c, model.GetValidMsg(err, &form), "1")
			//返回注册页
			c.Redirect(302, "/register")
			return
		}

		//执行数据库查询用户是否已存在
		var user model.User                            //定义数据库查询变量
		model.DB.First(&user, "email = ?", form.Email) //执行查询
		if user.ID != 0 {                              //如果查询到用户
			//创建消息
			model.FlashSet(c, "邮箱已被注册，请直接登录，若忘记密码请修改密码", "1")
			//返回注册页
			c.Redirect(302, "/register")
			return
		}

		code := service.Da_xie(form.Code) //验证码转大写
		// 核验验证码
		if !model.ValidateCode(form.Email, code) { //核验验证码
			//创建消息
			model.FlashSet(c, "验证码错误", "1")
			//返回注册页
			c.Redirect(302, "/register")
			return
		}

		//开始写入数据库逻辑
		err = model.DB.Create(&model.User{Email: form.Email, Password: service.SetPassword(form.Password)}).Error

		if err != nil {
			c.JSON(400, gin.H{"msg": "Database error, please try again later or report this issue to our customer support team"})
			log.Print("注册用户写入数据库错误：", err)
			return
		}

		//创建消息
		model.FlashSet(c, "用户创建成功", "3")
		//返回登陆页
		c.Redirect(302, "/login")
		return

	}
	msg := model.FlashGet(c)

	//如果是get直接返回登录页
	c.HTML(200, "register.html", gin.H{"title": "注册", "msg": msg})
}

// 登录页页
func Login(c *gin.Context) {
	// 获取当前请求的HTTP方法
	method := c.Request.Method
	//如果是post
	if method == "POST" {
		//先检查表单错误
		var form model.Login                         //定义一个表单变量
		err := c.ShouldBindWith(&form, binding.Form) //执行绑定
		if err != nil {                              //如果验证失败
			//创建消息
			model.FlashSet(c, model.GetValidMsg(err, &form), "1")
			//跳转页面
			c.Redirect(302, "/login")
			return
		}

		var user model.User
		//查询数据库
		model.DB.First(&user, "email = ?", form.Email)
		if user.ID == 0 { //如果未查询到用户
			//创建消息
			model.FlashSet(c, "账户或密码错误", "1")
			//跳转页面
			c.Redirect(302, "/login")
			return
		}
		// 验证密码如果没通过
		if ok := service.Cpassword(user.Password, form.Password); !ok {
			//创建消息
			model.FlashSet(c, "账户或密码错误", "1")
			//跳转页面
			c.Redirect(302, "/login")
			return
		}

		//密码正确后的逻辑
		//调用jwt签发函数并保存到cookie，也可以直接返回json的token让客户端自己处理
		tokenString, err := service.GenerateJWT(form.Email)
		if err != nil {
			c.JSON(500, gin.H{"error": "token generation failed"})
			return
		}
		//设置cookie有效期15天
		c.SetCookie("token", tokenString, 1296000, "/", model.Domain, false, true)
		//跳转到后台路由
		model.FlashSet(c, "登陆成功", "3")
		c.Redirect(302, "/")
		return

	}
	msg := model.FlashGet(c)
	//如果是get直接返回登录页
	c.HTML(200, "login.html", gin.H{"title": "登陆", "msg": msg})
}

// 定义验证码的json结构体
type EmailJson struct {
	Email string `json:"email" binding:"required,email,max=200" msg:"不支持的邮箱地址"`
}

// 获取验证码API
func Getcode(c *gin.Context) {
	// 解析请求体中的 JSON 数据到结构体
	var emailjson EmailJson
	if err := c.BindJSON(&emailjson); err != nil {
		//返回错误信息
		c.JSON(400, gin.H{"msg": model.GetValidMsg(err, &emailjson)})
		return
	}

	// 查找或创建记录
	var user model.Captcha
	//查询一条数据
	result := model.DB.First(&user, "email = ?", emailjson.Email) //搜索条件为name字段，搜索值为传递的路由参数
	if result.Error != nil {
		// 判断是否查询到数据
		if result.Error != gorm.ErrRecordNotFound {
			//如果不是空值则代表数据库其他错误
			c.JSON(500, gin.H{"msg": "数据库错误，请晚点再试"}) //返回查询失败
			log.Print("验证码数据库错误查询错误：", result.Error)
			return
		}
		log.Print("没有这条数据开始创建")
		//如果未查询到数据则创建它
		code := service.Captcha(6) //获得验证码

		//写入数据
		result = model.DB.Create(&model.Captcha{Email: emailjson.Email, Code: code, Date: time.Now()})
		if result.Error != nil {
			c.JSON(500, gin.H{"msg": "数据库错误，请晚点再试"}) //返回查询失败
			log.Print("验证码数据库创建失败：", result.Error)
			return
		}
		//返回成功
		c.JSON(200, gin.H{"msg": true})
		return
	}

	// //给现在的时间加上2分钟
	// newTime := user.Date.Add(2 * time.Minute)
	log.Print("有这条数据开始更新")

	//计算当前时间和对比时间的时间差
	timeDiff := time.Since(user.Date)
	// 判断时间差是否超过 2 分钟
	if timeDiff < 2*time.Minute {
		// 如果没有超过2分钟,返回成功和通知
		c.JSON(400, gin.H{"msg": "获取间隔不能低于2分钟"})
		log.Println("时间差没有超过 2 分钟")
		return
	}

	//超过2分钟则更新验证码和时间
	code := service.Captcha(6) //获得验证码

	user.Code = code
	user.Date = time.Now()
	result = model.DB.Save(&user)
	if result.Error != nil {
		// 修改失败
		c.JSON(500, gin.H{"msg": "数据库错误，请稍后再试"})
		log.Println("修改验证码失败：", result.Error)
		return
	}
	//返回成功
	c.JSON(200, gin.H{"msg": true})
}

//修改密码
func ChangePassword(c *gin.Context) {
	// 获取当前请求的HTTP方法
	method := c.Request.Method
	//如果是post
	if method == "POST" {
		//先检查表单错误
		var form model.Register //定义一个表单变量

		err := c.ShouldBindWith(&form, binding.Form) //执行绑定
		if err != nil {                              //如果验证失败
			//创建消息
			model.FlashSet(c, model.GetValidMsg(err, &form), "1")
			c.Redirect(302, "/changepassword")
			return
		}

		//执行数据库查询用户是否有这个用户
		var user model.User                            //定义数据库查询变量
		model.DB.First(&user, "email = ?", form.Email) //执行查询
		if user.ID == 0 {                              //如果查询到用户
			//创建消息
			model.FlashSet(c, "该邮箱并未注册为本站用户，请检查", "1")
			c.Redirect(302, "/changepassword")
			return
		}

		code := service.Da_xie(form.Code) //验证码转大写
		// 核验验证码
		if !model.ValidateCode(form.Email, code) { //核验验证码
			//创建消息
			model.FlashSet(c, "验证码错误", "1")
			c.Redirect(302, "/changepassword")
			return
		}

		//更新数据库逻辑
		user.Password = service.SetPassword(form.Password)
		if model.DB.Save(&user).Error != nil {
			//创建消息
			model.FlashSet(c, "数据库遇到错误请稍后再试", "1")
			c.Redirect(302, "/changepassword")
			return
		}

		//更新密码成功
		model.FlashSet(c, "密码重置成功", "3")
		c.Redirect(302, "/login")
		return
	}
	msg := model.FlashGet(c)
	//如果是get直接渲染页面
	c.HTML(200, "changepassword.html", gin.H{"title": "重置密码","msg":msg})
}

//退出登录
func Logout(c *gin.Context) {
	//删除cookie
	c.SetCookie("token", "", -1, "/", "", false, true)
	//跳转到后台路由
	model.FlashSet(c, "您已退出登录", "3")
	c.Redirect(302, "/login")
}