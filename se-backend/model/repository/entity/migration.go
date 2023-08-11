package entity

import "time"

type Migration struct {
	ID        uint64    `gorm:"column:mi_id;primaryKey;autoIncrement"`
	Name      string    `gorm:"column:mi_name;not null;type:varchar(255)"`
	CreatedAt time.Time `gorm:"column:mi_created_at;not null"`
}

func (Migration) TableName() string {
	return "mi_migrations"
}
