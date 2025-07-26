package utils

import(
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	Password  string    `gorm:"not null"`
}

type File struct {
	ID         uint      `gorm:"primaryKey"`
	Filename   string    `gorm:"not null"`
	Path       string    `gorm:"not null;uniqueIndex"`
	Size       int64     `gorm:"not null"`
	UploadedAt time.Time `gorm:"autoCreateTime"`
	UploadedBy uint      `gorm:"not null"`
}
