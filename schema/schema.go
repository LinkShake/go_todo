package schema

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Email string `gorm:"unique;not null"`
	Pwd string	`gorm:"not null"`
	Todos []Todo `gorm:"foreignKey:UserId;default:[]"`
}

type Todo struct {
	gorm.Model
	ID uint `gorm:"primaryKey"`
	Text string `gorm:"not null"`
	UserId uuid.UUID `gorm:"type:uuid;not null"`
}