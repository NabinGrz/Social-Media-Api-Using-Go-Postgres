package postController

import (
	"bytes"
	"io"
	"log"
	"net/http"

	userPostModel "github.com/NabinGrz/SocialMedia/src/post/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func GetAllPost(c *gin.Context, db *gorm.DB) {
	var posts []userPostModel.Post

	if err := db.Preload("User").Preload("MediaDetails").Preload("Likes").Preload("Likes.User").Preload("Shares").Preload("Shares.User").Preload("Comments").Preload("Comments.User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":        posts,
		"total_count": len(posts),
	})

}
func GetAllOwnPost(c *gin.Context, db *gorm.DB) {
	userIDString := c.GetString("userid")
	userID, _ := uuid.Parse(userIDString)

	var posts []userPostModel.Post

	if err := db.Preload("User").Preload("MediaDetails").Preload("Likes").Preload("Likes.User").Preload("Shares").Preload("Shares.User").Preload("Comments").Preload("Comments.User").Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":        posts,
		"total_count": len(posts),
	})

}

func GetPostDetails(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post Id must be valid UUID"})
		return
	}
	var post userPostModel.Post
	if err := db.Preload("User").Preload("MediaDetails").Preload("Likes").Preload("Likes.User").Preload("Shares").Preload("Shares.User").Preload("Comments").Preload("Comments.User").Find(&post, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}
func UpdatePost(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post Id must be valid UUID"})
		return
	}
	var post userPostModel.Post

	if err := uuid.Validate(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post Id must be valid UUID"})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	var existingPost userPostModel.Post
	if err := db.First(&existingPost, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post Not Found"})
		return
	}

	existingPost.Caption = post.Caption
	if err := db.Save(&existingPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
	})
}
func DeletePost(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post Id must be valid UUID"})
		return
	}
	var post userPostModel.Post

	if err := db.Delete(&post, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}

func CreatePost(c *gin.Context, db *gorm.DB) {

	var input userPostModel.Post

	// Print the raw request body for debugging
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read request body"})
		return
	}
	log.Printf("Raw body: %s", string(bodyBytes))

	// Reset the request body to allow Gin to read it again
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetString("userid")
	parsedUUID, _ := uuid.Parse(userID)
	// post.UserID = parsedUUID

	post := userPostModel.Post{
		Caption: input.Caption,
		UserID:  parsedUUID,
	}

	// Create the post
	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create media details
	for _, mediaInput := range input.MediaDetails {
		mediaDetail := userPostModel.MediaDetail{
			PostID:   post.ID,
			PostType: mediaInput.PostType,
			Url:      mediaInput.Url,
		}
		if err := db.Create(&mediaDetail).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		post.MediaDetails = append(post.MediaDetails, mediaDetail)
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func userLikedPost(userID, postID uuid.UUID, db *gorm.DB) (bool, error) {
	var count int64
	if err := db.Model(&userPostModel.Like{}).
		Where("user_id = ? AND post_id = ?", userID, postID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
func dislikePost(userID, postID uuid.UUID, db *gorm.DB) error {
	if err := db.Where("user_id = ? AND post_id = ?", userID, postID).
		Delete(&userPostModel.Like{}).Error; err != nil {
		return err
	}
	return nil
}

func LikePost(c *gin.Context, db *gorm.DB) {
	likeShare(c, db, true)
}
func SharePost(c *gin.Context, db *gorm.DB) {
	likeShare(c, db, false)
}
func likeShare(c *gin.Context, db *gorm.DB, isLiking bool) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post Id must be valid UUID"})
		return
	}
	userIDString := c.GetString("userid")
	userID, _ := uuid.Parse(userIDString)
	var post userPostModel.Post
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if isLiking {
		// Check if the user has already liked the post
		isLiked, _ := userLikedPost(userID, post.ID, db)
		if isLiked {
			dislikePost(userID, post.ID, db)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Post has been disliked"})
			return
		}
	}

	// Constructing the like/share object
	if isLiking {
		like := userPostModel.Like{
			PostID: post.ID,
			UserID: userID,
		}
		post.Likes = append(post.Likes, like)
	} else {
		share := userPostModel.Share{
			PostID: post.ID,
			UserID: userID,
		}
		post.Shares = append(post.Shares, share)
	}

	// Save the updated post back to the database
	if err := db.Save(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if isLiking {
		c.JSON(http.StatusOK, gin.H{"message": "Post has been liked"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Post has been shared"})
		return
	}

}
func CommentOnPost(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post Id must be valid UUID"})
		return
	}
	userIDString := c.GetString("userid")
	userID, _ := uuid.Parse(userIDString)
	var post userPostModel.Post
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var commented userPostModel.CommentDetail
	if err := c.ShouldBindJSON(&commented); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	comment := userPostModel.CommentDetail{
		PostID:  post.ID,
		UserID:  userID,
		Comment: commented.Comment,
	}
	post.Comments = append(post.Comments, comment)

	// Save the updated post back to the database
	if err := db.Save(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Commented successfully"})

}
