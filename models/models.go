package models

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(15);not null"`
	CreatedAt time.Time `gorm:"not null"`
	DeletedAt time.Time
}

type Room struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	DeletedAt time.Time
}

type Message struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint `gorm:"not null"`
	RoomID    uint `gorm:"not null"`
	Message   string `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"not null"`
	DeletedAt time.Time
}
