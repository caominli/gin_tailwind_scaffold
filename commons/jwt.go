package common

import (
	model "ginTailwindcssBase/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

// 定义 JWT 签名密钥
var jwtKey = []byte("s1D5e+sUc8eF+2rNexP7IM7B6Nt6_kr7L3e4y")

// 定义 JWT 签发函数
func GenerateJWT(email string) (string, error) {
	// 创建 JWT Claims 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		//设置超时时间15天
		"exp":      time.Now().Add(time.Hour * 360).Unix(),
		"useremail": email, //这是存入的用户名
	})
	// 生成 JWT Token
	tokenString, err := token.SignedString(jwtKey)
	// 使用密钥进行签名生成字符串 token
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 定义 JWT 认证中间件函数，必须登陆用户才可以
func AuthUser(c *gin.Context) {
	//获取cookie中的token字段
	tokenString, _ := c.Cookie("token")

	// 如果没有提供信息
	if tokenString == "" {
		//创建消息
		model.FlashSet(c, "请先登陆", "2")
		//跳转页面
		c.Redirect(302, "/login")
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//这里的参数就是方法，方法直接返回jwtKey令牌
		return jwtKey, nil
	})
	// 检查token是否有效
	if !token.Valid {
		//如果无效
		model.FlashSet(c, "请先登陆", "2")
		//跳转页面
		c.Redirect(302, "/login")
		c.Abort()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		log.Println("找到用户：", claims["useremail"])
		// 将解析后的claims存储到上下文中，以便后续处理程序可以访问
		//例如返回用户名，权限
		c.Set("useremail", claims["useremail"])
	} else {
		c.JSON(500, gin.H{"token_err": err})
		c.Abort()
		return
	}
	c.Next()
}

//jwt中间件，检查是否登录和返回用邮箱,返回str,没有返回空字符串
func IsLogin(c *gin.Context) {
	//获取cookie中的token字段
	tokenString, _ := c.Cookie("token")

	// 如果没有提供返回空
	if tokenString == "" {
		c.Set("useremail", nil)
		return
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//这里的参数就是方法，方法直接返回jwtKey令牌
		return jwtKey, nil
	})
	// 检查token是否有效
	if !token.Valid {
		//如果无效
		c.Set("useremail", nil)
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//找到用户
		c.Set("useremail", claims["useremail"])
	} else {
		//否则系统错误返回空，或者打印信息
		c.Set("useremail", nil)
		return
	}
	c.Next()
}
