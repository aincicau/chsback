package entity

type Session struct {
	ID     string `gorm:"type:uuid;default:uuid_generate_v4();primary_key;column:id"`
	UserID string `gorm:"type:uuid;not null;column:user_id" json:"user_id"`
}
