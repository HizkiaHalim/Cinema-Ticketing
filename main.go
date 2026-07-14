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
	r.POST("/sign-up", controllers.SignUp)
	r.POST("/login", controllers.Login)

	// Protected routes
	authorized := r.Group("/")
	authorized.Use(middleware.RequireAuth)

	authorized.GET("/movie-list", controllers.GetMovieList)
	authorized.GET("/add-movie", controllers.RegisterMovie)
	authorized.GET("/update-movie", controllers.UpdateMovie)
	// authorized.GET("/delete-movie", controllers.RemoveMovie)

	r.Run(":8080")
}
