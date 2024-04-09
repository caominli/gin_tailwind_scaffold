package models //定义在models包下

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10" //导入自定义验证器
	// "log"
	"reflect" //验证器定义错误需要
	"strings"
	"time"
)

// 定义一个User的数据库模型
type User struct {
	ID        uint      `gorm:"primaryKey"`           //id
	Email     string    `gorm:"uniqueIndex;not null"` //邮箱
	Password  []byte    `gorm:"not null"`             //密码
	CreatedAt time.Time `gorm:"autoCreateTime"`       //创建时间
}

// 注册表单验证器
type Register struct {
	Email    string `form:"email" binding:"required,email,max=200" msg:"不支持的邮箱地址"`          //用户名
	Password string `form:"password" binding:"required,min=6,max=25" msg:"密码不符合规范，6-25个字符"` //密码,必填，小于5，大于10位数
	Code     string `form:"code" binding:"required,alphanum,len=6" msg:"验证码是6个字母和数字组成"`     //验证码
}

// 注册表单验证器
type Login struct {
	Email    string `form:"email" binding:"required,email,max=200" msg:"不支持的邮箱地址"`          //用户名
	Password string `form:"password" binding:"required,min=6,max=25" msg:"密码不符合规范,6-25个字符"` //密码,必填，小于5，大于10位数
}

// 定义一个验证码库
type Captcha struct {
	ID    uint      `gorm:"primaryKey"`           //id
	Email string    `gorm:"uniqueIndex;not null"` //邮箱
	Code  string    `gorm:"not null"`             //代码
	Date  time.Time `gorm:"not null"`             //生成时间
}

// 自定义验证器结构体的错误信息，接收err错误对象和结构体的指针，返回str类型
func GetValidMsg(err error, obj interface{}) string {
	//获取结构体的类型信息
	getObj := reflect.TypeOf(obj)
	// 检查错误是否属于validator.ValidationErrors类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			// 检查是否可以通过字段名获取结构体字段信息
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				return f.Tag.Get("msg") //获取自定义的结构体msg参数
			}
		}
	}
	// 如果无法通过验证器错误获取自定义消息，返回默认错误消息
	return err.Error()
}

// 验证码核验，传入邮箱，表单验证码，返回bool
func ValidateCode(email string, code string) bool {
	//查询数据
	var user Captcha                    //定义数据库查询变量
	DB.First(&user, "email = ?", email) //执行查询
	if user.ID == 0 {
		//没有查询到用户
		return false
	}

	if user.Code == code {
		return true
	}
	return false
}

// 给予闪现消息
// 参数：gin对象，消息内容str，类型str
func FlashSet(c *gin.Context, msg string, msgtype string) {
	//创建消息,用冒号分开
	value := msg + ":" + msgtype
	c.SetCookie("msg", value, 40, "/", Domain, false, true)
}

// 读取闪现消息，返回包含两个值的数组，消息内容和类型
func FlashGet(c *gin.Context) []string {
	var values []string
	//读取cookie
	msg, _ := c.Cookie("msg")
	// 如果没有 token 字段，返回 401 错误
	if msg != "" {
		values = strings.Split(msg, ":")
	} else {
		values = []string{}
	}
	//删除cookie
	c.SetCookie("msg", "", -1, "/", "", false, true)
	return values
}
