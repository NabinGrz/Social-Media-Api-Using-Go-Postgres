package profileController

import (
	"net/http"

	userModel "github.com/NabinGrz/SocialMedia/src/authentication/models"
	authServices "github.com/NabinGrz/SocialMedia/src/authentication/services"
	cloudinaryController "github.com/NabinGrz/SocialMedia/src/cloudinary/controller"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func GetUserProfile(c *gin.Context, db *gorm.DB) {
	userIDString := c.GetString("userid")
	var existingUser userModel.User
	if err := db.Find(&existingUser, "id = ?", userIDString).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
		return
	}

	c.JSON(http.StatusOK, existingUser)
}
func UpdateDetails(c *gin.Context, db *gorm.DB) {

	userIDString := c.GetString("userid")
	var existingUser userModel.User
	if err := db.Find(&existingUser, "id = ?", userIDString).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
		return
	}
	var user userModel.UpdateUserInput
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if user.FullName != nil {
		existingUser.FullName = *user.FullName
	}
	if user.Email != nil {
		isValid := authServices.IsValidEmail(*user.Email)
		if isValid {
			existingUser.Email = *user.Email
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
			return
		}
	}
	if user.Username != nil {
		existingUser.Username = *user.Username
	}

	if err := db.Save(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile url"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile details updated successfully"})
}
func UpdateProfileImage(c *gin.Context, db *gorm.DB) {
	userIDString := c.GetString("userid")
	userID, _ := uuid.Parse(userIDString)

	fileUrl, err := cloudinaryController.UploadFile(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var existingUser userModel.User
	if err := db.First(&existingUser, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post Not Found"})
		return
	}
	existingUser.ProfileUrl = fileUrl
	if err := db.Save(&existingUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile url"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile URL updated successfully"})
}
