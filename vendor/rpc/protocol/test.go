package protocol

import (
	"time"
)

//用户信息
type User struct {
	Id       int64
	Username string    //用户名
	Password string    //密码
	Date     time.Time //日期
}
