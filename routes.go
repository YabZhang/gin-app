package main

func initializeRoutes() {
	// index page
	router.GET("/", showIndexPage)

	// single article detail
	router.GET("/article/view/:article_id", getArticle)
}


