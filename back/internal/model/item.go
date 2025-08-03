package model

type Item struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	Want       string `gorm:"not null"`
	ImagePaths string `gorm:"not null"` // 複数の画像(カンマ区切りの文字列)
	UserID     uint   `gorm:"not null"` // 外部キー
}
