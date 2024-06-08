package userModel

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Email      string    `gorm:"not null" json:"email" `
	Password   string    `gorm:"not null" json:"password"`
	ProfileUrl string    `json:"profile_url"`
	Username   string    `gorm:"not null" json:"username"`
	FullName   string    `gorm:"not null" json:"full_name"`
}

type UpdateUserInput struct {
	Email      *string `json:"email"`
	ProfileUrl *string `json:"profile_url"`
	Username   *string `json:"username"`
	FullName   *string `json:"full_name"`
}
