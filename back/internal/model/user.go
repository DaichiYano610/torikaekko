package model

type User struct {
	ID       uint   `gorm:"primaryKey" json:"-"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"` // JSONでは文字列、DBではハッシュ保存
}
