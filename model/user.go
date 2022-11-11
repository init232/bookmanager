package model

// 定义User结构体
type User struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// 自定义表名
func (*User) TableName() string {
	return "user"
}
