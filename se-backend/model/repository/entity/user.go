package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint64         `gorm:"column:us_id;primaryKey;autoIncrement"`
	CreatedAt time.Time      `gorm:"column:us_created_at;not null"`
	UpdatedAt time.Time      `gorm:"column:us_updated_at;not null"`
	DeletedAt gorm.DeletedAt `gorm:"column:us_deleted_at"`
	FirstName string         `gorm:"column:us_first_name;type:varchar(50);not null"`
	LastName  string         `gorm:"column:us_last_name;type:varchar(50);not null"`
	Email     string         `gorm:"column:us_email;type:varchar(255);not null"`
	Username  string         `gorm:"column:us_username;type:varchar(50);not null"`
	Password  string         `gorm:"column:us_password;type:varchar(255);not null"`
	Verified  bool           `gorm:"column:us_verified;type:bool;default:false;not null"`
}

func (*User) TableName() string {
	return "us_users"
}
