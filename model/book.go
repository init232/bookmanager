package model

// 定义book结构体
type Book struct {
	ID   int64  `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	User []User `gorm:"many2many:book_users"`
}

// 自定义表名
func (*Book) TableName() string {
	return "book"
}
