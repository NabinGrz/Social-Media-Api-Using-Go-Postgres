package userPostModel

import (
	userModel "github.com/NabinGrz/SocialMedia/src/authentication/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Post represents a social media post.
// Post represents a social media post.
type Post struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Caption string    `gorm:"not null"`
	UserID  uuid.UUID
	User    userModel.User `gorm:"foreignKey:UserID"`
	// MediaDetail []MediaDetail    `gorm:"foreignKey:MediaDetailID"`
	Likes  []Like
	Shares []Share
	// Shares      []userModel.User `gorm:"foreignKey:UserID"`
	// Comments    []CommentDetail  `gorm:"foreignKey:CommentID"`
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

type MediaDetail struct {
	gorm.Model
	MediaDetailID uuid.UUID `gorm:"primaryKey"`
	PostType      string    `gorm:"not null"`
	Url           string    `gorm:"not null"`
}
