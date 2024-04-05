package models  //定义在models包下

import (
	"time"
)

//定义一个User的数据库模型
type User struct {
	ID           uint  `gorm:"primaryKey"`//id
	Name         string  //用户名
	Email        string  //邮箱
	Password	string  //密码
	CreatedAt time.Time  //创建时间
  }