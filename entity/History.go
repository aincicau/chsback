package entity

import "time"

type History struct {
	ID       string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key;column:id"`
	FileName string    `gorm:"type:varchar(255);not null"`
	Date     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null"`
	Result   string    `gorm:"type:varchar(255);not null"`
	UserID   string    `gorm:"type:uuid;not null;column:user_id" json:"user_id"`
}