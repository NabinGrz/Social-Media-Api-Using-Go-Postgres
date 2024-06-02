package userModel

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email      string    `gorm:"not null"`
	Password   string    `gorm:"not null"`
	ProfileUrl string
	Username   string `gorm:"not null"`
	FullName   string `gorm:"not null"`
}
