package authController

import (
	"net/http"

	userModel "github.com/NabinGrz/SocialMedia/src/authentication/models"
	authServices "github.com/NabinGrz/SocialMedia/src/authentication/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func LoginHandler(ctx *gin.Context, db *gorm.DB) {
	var user userModel.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		if user.Email == "" || user.Password == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "email/password field is required"})
			return
		} else {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	token, error := authServices.Login(user, db)

	if error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, token)

}

func RegisterHandler(ctx *gin.Context, db *gorm.DB) {
	var user userModel.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		emptyError := authServices.IsValid(user)
		if emptyError != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": emptyError.Error()})
			return
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	result, err := authServices.Register(user, db)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
