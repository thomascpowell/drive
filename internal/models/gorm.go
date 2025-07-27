package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`           // primary source of identification
	Username  string    `gorm:"uniqueIndex;not null"` // auth purposes
	CreatedAt time.Time `gorm:"autoCreateTime"`
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
