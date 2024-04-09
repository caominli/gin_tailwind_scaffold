package models  //定义在models包下

import (
	"gorm.io/gorm"
	// "gorm.io/driver/sqlite" //官方sqlite
	"github.com/glebarez/sqlite" //纯go实现的sqlite
)

//定义一个常量用来操作数据库，常量就是初始化后返回的db
var DB = Init()

var(
	Port string //端口
	Domain string  //域名，用作cookie作用域
)

//初始化数据库连接
func Init() *gorm.DB {
	//如果要sqlite支持外键，使用sqlite.Open("test.db?_fk=1")参数启用数据库层外键约束
	db, err := gorm.Open(sqlite.Open("btcmai.db?_fk=1"),&gorm.Config{})
	if err != nil {
		panic("无法连接数据库")
	}
	// 自动创建表
	db.AutoMigrate(&User{},&Captcha{})
	return db
}

func init() {
	// 开发环境
	Port = "80"
	Domain = "127.0.0.1"

	// // 生产环境
    // Domain = "b.usmm.cn"
    // Port = "5280"
}