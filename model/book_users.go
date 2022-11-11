package model

type BookUser struct {
	UserID int64 `gorm:parimary_key`
	BookID int64 `gorm:parimary_key`
}
