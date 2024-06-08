package router

import (
	authController "github.com/NabinGrz/SocialMedia/src/authentication/controller"
	authServices "github.com/NabinGrz/SocialMedia/src/authentication/services"
	cloudinaryController "github.com/NabinGrz/SocialMedia/src/cloudinary/controller"
	postController "github.com/NabinGrz/SocialMedia/src/post/controller"
	profileController "github.com/NabinGrz/SocialMedia/src/profile/controller"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Router(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	//!! GENERATING JWT TOKEN AFTER LOGIN
	router.POST("/login", func(ctx *gin.Context) { authController.LoginHandler(ctx, db) })
	router.POST("/register", func(ctx *gin.Context) { authController.RegisterHandler(ctx, db) })

	authorized := router.Group("/api")

	authorized.Use(authServices.JWTMiddleware())
	{
		// 	authorized.GET("/", func(ctx *gin.Context) {
		// 		ctx.IndentedJSON(http.StatusOK, "Hello World")
		// 	})
		// 	//!! USER
		authorized.GET("/userDetails", func(ctx *gin.Context) { profileController.GetUserProfile(ctx, db) })
		authorized.PUT("/updateProfileUrl", func(ctx *gin.Context) { profileController.UpdateProfileImage(ctx, db) })
		authorized.PUT("/updateProfileDetail", func(ctx *gin.Context) { profileController.UpdateDetails(ctx, db) })

		// 	//!! POST
		authorized.GET("/posts", func(ctx *gin.Context) { postController.GetAllPost(ctx, db) })
		authorized.GET("/posts/own", func(ctx *gin.Context) { postController.GetAllOwnPost(ctx, db) })
		authorized.POST("/post", func(ctx *gin.Context) { postController.CreatePost(ctx, db) })
		authorized.GET("/post/:id", func(ctx *gin.Context) { postController.GetPostDetails(ctx, db) })
		authorized.DELETE("/post/:id", func(ctx *gin.Context) { postController.DeletePost(ctx, db) })
		authorized.PUT("/post/:id", func(ctx *gin.Context) { postController.UpdatePost(ctx, db) })
		authorized.POST("/post/like/:id", func(ctx *gin.Context) { postController.LikePost(ctx, db) })
		authorized.POST("/post/share/:id", func(ctx *gin.Context) { postController.SharePost(ctx, db) })
		authorized.PUT("/post/comment/:id", func(ctx *gin.Context) { postController.CommentOnPost(ctx, db) })

		// 	//!! FILE UPLOAD
		authorized.POST("/upload-file", func(ctx *gin.Context) { cloudinaryController.UpdatePOSTImage(ctx) })
	}

	return router
}
