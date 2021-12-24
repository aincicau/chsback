package entity

type User struct {
	ID       string `gorm:"type:uuid;default:uuid_generate_v4();primary_key;column:id"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`
}
