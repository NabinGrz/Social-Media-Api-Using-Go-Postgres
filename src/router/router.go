package router

import (
	authController "github.com/NabinGrz/SocialMedia/src/authentication/controller"
	authServices "github.com/NabinGrz/SocialMedia/src/authentication/services"
	postController "github.com/NabinGrz/SocialMedia/src/post/controller"
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
		// 	authorized.PUT("/updateProfileUrl/:id", profileController.UpdateProfileImage)
		// 	authorized.PUT("/updateProfileDetail/:id", profileController.UpdateDetails)

		// 	//!! POST
		// 	authorized.GET("/posts", postController.GetAllPost)
		authorized.POST("/post", func(ctx *gin.Context) { postController.CreatePost(ctx, db) })
		authorized.GET("/post/:id", func(ctx *gin.Context) { postController.GetPostDetails(ctx, db) })
		// 	authorized.DELETE("/post/:id", postController.DeletePost)
		// 	authorized.PUT("/post/:id", postController.UpdatePost)
		// 	authorized.POST("/post/like/:id", postController.LikePost)
		// 	authorized.POST("/post/comment/:id", postController.CommentPost)
		// 	authorized.POST("/post/share/:id", postController.SharePost)

		// 	//!! FILE UPLOAD
		// 	authorized.POST("/uploadFile", cloudinaryController.UpdatePOSTImage)
	}

	return router
}
