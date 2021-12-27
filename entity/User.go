package entity

type User struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primary_key;column:id"`
	Username  string    `gorm:"type:varchar(255);not null;unique"`
	Email     string    `gorm:"type:varchar(255);not null;unique"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Histories []History `gorm:"foreignKey:UserID"`
	Sessions  []Session `gorm:"foreignKey:UserID"`
}
