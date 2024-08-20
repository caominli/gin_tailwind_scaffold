package common

import (
	"golang.org/x/crypto/bcrypt" //密码加密解密
	"log"
	"time"
	"strings"
    "math/rand"
)



//字符串转大写，用于验证码可以不区分大小写
//参数：str ，返回：str
func Da_xie(code string) (code2 string) {
    return strings.ToUpper(code)
}
   
//生成随机码,参数:长度
func Captcha(length int) string {
    const charset = "ABCDEFGHJKMNPQRSTUVWXYZ123456789"
    var sb strings.Builder
    sb.Grow(length)

	// 创建一个伪随机数生成器的种子源
    source := rand.NewSource(time.Now().UnixNano())

    // 使用种子源创建一个伪随机数生成器
	random := rand.New(source)

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}
	return string(b)

}

//加密密码，参数：str，返回：[]byte
func SetPassword(text string) []byte {
	// 对密码进行哈希加密
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		log.Print("密码加密错误：",err)
		return nil
	}
	return passwordHash
}

//效验密码，参数：数据库读取的[]byte密码，表单获取的str密码，返回布尔
func Cpassword(modelupassword []byte,password string) bool {
	err := bcrypt.CompareHashAndPassword(modelupassword, []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
	

}

