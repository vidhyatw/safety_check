// routes.go

package main

func initializeRoutes() {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())

	// Handle the index route
	router.GET("/", showIndexPage)

	// Group user related routes together
	userRoutes := router.Group("/u")
	{
		// Handle the GET requests at /u/login
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)

		// Handle POST requests at /u/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)

		// Handle GET requests at /u/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", ensureLoggedIn(), logout)

		// Handle the GET requests at /u/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)

		// Handle POST requests at /u/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
	}

	// Group review related routes together
	reviewRoutes := router.Group("/review")
	{
		// Handle GET requests at /review/view/some_review_id
		reviewRoutes.GET("/view/:review_id", getreview)

		// Handle the GET requests at /review/create
		// Show the review creation page
		// Ensure that the user is logged in by using the middleware
		reviewRoutes.GET("/create", ensureLoggedIn(), showreviewCreationPage)

		// Handle POST requests at /review/create
		// Ensure that the user is logged in by using the middleware
		reviewRoutes.POST("/create", ensureLoggedIn(), createreview)
	}
}
