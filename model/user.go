package model

import (
	"time"
)

type User struct {
	ID        int64  `gorm:"primary_key;autoIncrement;type:BIGINT"`
	Username  string `gorm:"type:VARCHAR(32)"`
	Code      string `gorm:"type:VARCHAR(7)"`
	Fullname  string `gorm:"type:VARCHAR(255)"`
	Email     string `gorm:"type:VARCHAR(100)"`
	RoleId    int64  `gorm:"primary_key;type:BIGINT"`
	Password  string `gorm:"type:VARCHAR(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
