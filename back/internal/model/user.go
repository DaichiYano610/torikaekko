package model

type User struct {
	Username string
	Password string
}

type Users struct {
	Username string `gorm:"primaryKey"`
	Password []byte `gorm:"not null"` // パスワードのハッシュ（bcryptなど）
}
