package entity

import "time"

type Session struct {
	ID             string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key;column:id"`
	UserID         string    `gorm:"type:uuid;not null;column:user_id" json:"user_id"`
	ExpirationDate time.Time `gorm:"type:timestamp;not null;column:expiration_date"`
}
