package routers

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	authRouter := r.Group("/users")
	{
		// Create/Register User
		authRouter.POST("/register", controllers.UserRegister)

		// Login User
		authRouter.POST("/login", controllers.UserLogin)
	}

	userRouter := r.Group("/users")
	{
		userRouter.Use(middlewares.Authentication())

		// Get User For Profile
		userRouter.GET("/", controllers.UserGet)

		// Get all User for Search any user
		userRouter.GET("/all", controllers.UserGetAll)

		// Update Profile User
		userRouter.PUT("/", controllers.UserUpdate)

		// Delete Account user
		userRouter.DELETE("/", controllers.UserDelete)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())

		// Create New Photo
		photoRouter.POST("/", controllers.CreatePhoto)

		// Get Account Photo
		photoRouter.GET("/", controllers.PhotoGet)

		// Update Account photo
		photoRouter.PUT("/:photoId", controllers.PhotoUpdate)

		// Delete Photo Account
		photoRouter.DELETE("/:photoId", controllers.PhotoDeleted)
	}

	sosmedRouter := r.Group("/socialmedias")
	{
		sosmedRouter.Use(middlewares.Authentication())

		// Create New Sosmed
		sosmedRouter.POST("/", controllers.CreateSosmed)

		// Get Account Sosmed
		sosmedRouter.GET("/", controllers.SosmedGet)

		// Get Data sosmed by user id
		sosmedRouter.GET("/:userIdSos", controllers.SosmedGetByIdUser)

		// Update Account sosmed
		sosmedRouter.PUT("/:socialMediaId", controllers.SosmedUpdate)

		// Delete Account Sosmed
		sosmedRouter.DELETE("/:socialMediaId", controllers.SosmedDeleted)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())

		// Create New Comment
		commentRouter.POST("/", controllers.CreateComment)

		// Get Comment
		commentRouter.GET("/:photoId", controllers.GetComment)

		// Update Comment
		commentRouter.PUT("/:commentId", controllers.UpdateComment)

		// Delete Comment
		commentRouter.DELETE("/:commentId", controllers.CommentDeleted)
	}
	return r
}
