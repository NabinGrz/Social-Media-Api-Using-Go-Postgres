package userPostModel

import (
	userModel "github.com/NabinGrz/SocialMedia/src/authentication/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// SocialMediaPost represents a social media post.
// SocialMediaPost represents a social media post.
type SocialMediaPost struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Caption string    `gorm:"not null"`
	UserID  uuid.UUID
	User    userModel.User `gorm:"foreignKey:UserID"`
	// MediaDetail []MediaDetail    `gorm:"foreignKey:MediaDetailID"`
	// LikeBy      []userModel.User `gorm:"many2many:post_likes;"`
	// Shares      []userModel.User `gorm:"many2many:post_shares;"`
	// Comments    []CommentDetail  `gorm:"foreignKey:CommentID"`
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
