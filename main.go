// main.go

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hackerearth/safety_check/handlers"
	"github.com/hackerearth/safety_check/middleware"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	//router.Static("/assets", "./assets")
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	router.Run()
}

func initializeRoutes() {

	// Use the SetUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(middleware.SetUserStatus())

	// Handle the index route
	router.GET("/", handlers.ShowIndexPage)

	router.Static("/assets", "assets/")

	router.Static("/slick", "slick/")

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Group user related routes together
	userRoutes := router.Group("/u")
	{
		// Handle the GET requests at /u/login
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/login", middleware.EnsureNotLoggedIn(), handlers.ShowLoginPage)

		// Handle POST requests at /u/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", middleware.EnsureNotLoggedIn(), handlers.PerformLogin)

		// Handle GET requests at /u/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", middleware.EnsureLoggedIn(), handlers.Logout)

		// Handle the GET requests at /u/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", middleware.EnsureNotLoggedIn(), handlers.ShowRegistrationPage)

		// Handle POST requests at /u/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", middleware.EnsureNotLoggedIn(), handlers.Register)
	}

	// Group review related routes together
	reviewRoutes := router.Group("/review")
	{
		// Handle GET requests at /review/view/some_review_id
		reviewRoutes.GET("/view/:placeid/:long/:lat", handlers.GetReview)

		// Handle the GET requests at /review/create
		// Show the review creation page
		// Ensure that the user is logged in by using the middleware
		//reviewRoutes.GET("/create", EnsureLoggedIn(), showreviewCreationPage)
		reviewRoutes.GET("/create", handlers.ShowWriteReviewPage)
		// Handle POST requests at /review/create
		// Ensure that the user is logged in by using the middleware
		//	reviewRoutes.POST("/create", EnsureLoggedIn(), createreview)
		reviewRoutes.POST("/create", handlers.CreateReview)

		reviewRoutes.GET("/score/:placeid/:long/:lat", handlers.GetSafetyScore)
	}
}
