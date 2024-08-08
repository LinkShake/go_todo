package schema

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username string
	Pwd string
	Todos []Todo `gorm:"foreignKey:UserId"`
}

type Todo struct {
	gorm.Model
	Text string
	UserId uuid.UUID `gorm:"type:uuid"`
}