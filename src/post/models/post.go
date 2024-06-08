package userPostModel

import (
	userModel "github.com/NabinGrz/SocialMedia/src/authentication/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Post represents a social media post.
type Post struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Caption      string    `gorm:"not null"`
	UserID       uuid.UUID
	User         userModel.User `gorm:"foreignKey:UserID"`
	MediaDetails []MediaDetail  `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE;" json:"media_details"`
	Likes        []Like
	Shares       []Share
	// Shares      []userModel.User `gorm:"foreignKey:UserID"`
	// Comments    []CommentDetail  `gorm:"foreignKey:CommentID"`
}
type MediaDetail struct {
	MediaDetailID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	PostID        uuid.UUID `gorm:"type:uuid" json:"post_id"`
	Post          Post      `gorm:"foreignKey:PostID;references:ID" json:"-"`
	PostType      string    `gorm:"not null" json:"post_type"`
	Url           string    `gorm:"not null" json:"url"`
}

// Input Definitions
type CreatePostInput struct {
	Caption      string                   `json:"caption" binding:"required"`
	UserID       string                   `json:"user_id" binding:"required,uuid"`
	MediaDetails []CreateMediaDetailInput `json:"media_details" binding:"required,dive"`
}

type CreateMediaDetailInput struct {
	PostType string `json:"post_type" binding:"required"`
	Url      string `json:"url" binding:"required"`
}
type Like struct {
	ID     uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	PostID uuid.UUID      `gorm:"type:uuid" json:"post_id"`
	Post   Post           `gorm:"foreignKey:PostID;references:ID" json:"-"`
	UserID uuid.UUID      `gorm:"type:uuid" json:"-"`
	User   userModel.User `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

type Share struct {
	ID     uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	PostID uuid.UUID      `gorm:"type:uuid" json:"post_id"`
	Post   Post           `gorm:"foreignKey:PostID;references:ID" json:"-"`
	UserID uuid.UUID      `gorm:"type:uuid" json:"-"`
	User   userModel.User `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

type CommentDetail struct {
	gorm.Model
	CommentID    uuid.UUID `gorm:"primaryKey"`
	Comment      string    `gorm:"not null"`
	CommentUsers string    `gorm:"foreignKey:UserID;not null"`
}

type PostByUser struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email      string    `gorm:"not null"`
	ProfileUrl string
	Username   string `gorm:"not null"`
	FullName   string `gorm:"not null"`
}
