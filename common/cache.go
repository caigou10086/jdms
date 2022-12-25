package common

import "time"

type User struct {
	Cookie    string `json:"cookie"`
	Name      string `json:"name"`
	IsLogin   bool   `json:"is_login"`
	UserAgent string `json:"user_agent"`
	// 毫秒
	Millisecond string `json:"millisecond"`
	// 秒杀时间
	StartTime time.Time `json:"start_time"`
	Eid       string    `json:"eid"`
	Fp        string    `json:"fp"`
}

// UserData 全局变量
var UserData *User

// GetUserData 获得全局变量
func GetUserData() *User {
	if UserData == nil {
		UserData = &User{}
	}
	return UserData
}

// IsLogin 判断是否登陆
func IsLogin() bool {
	user := GetUserData()
	return user.IsLogin
}
