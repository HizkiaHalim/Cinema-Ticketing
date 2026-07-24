package main

import (
	"github.com/HizkiaHalim/Cinema-Ticketing/controllers"
	"github.com/HizkiaHalim/Cinema-Ticketing/initializers"
	"github.com/HizkiaHalim/Cinema-Ticketing/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	// Health check routes
	r.GET("/", controllers.HealthCheck)
	r.GET("/health", controllers.HealthCheck)
	r.GET("/ping", controllers.HealthCheck)

	// User routes
	initUserRoutes(r)

	// Protected routes
	initMovieAdminRoutes(r)
	initMovieUserRoutes(r)

	r.Run(":8080")
}

func initUserRoutes(r *gin.Engine) {
	{
		r.POST("/sign-up", controllers.SignUp)
		r.POST("/login", controllers.Login)
	}

}

func initMovieAdminRoutes(r *gin.Engine) {
	adminAuthorized := r.Group("/")
	adminAuthorized.Use(middleware.RequireAdminAuth)

	{
		adminAuthorized.POST("/register-admin", controllers.RegisterAdmin)
		adminAuthorized.POST("/add-movie", controllers.RegisterMovie)
		adminAuthorized.POST("/update-movie", controllers.UpdateMovie)
		adminAuthorized.POST("/delete-movie", controllers.DeleteMovie)
	}

}

func initMovieUserRoutes(r *gin.Engine) {
	authorized := r.Group("/")
	authorized.Use(middleware.RequireAuth)

	{
		authorized.GET("/movie-list", controllers.GetMovieList)
		authorized.GET("/movie-detail", controllers.GetMovieDetailOnDate)
	}
}
