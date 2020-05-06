package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// set the default router
	router := gin.Default()

	// load templates
	router.LoadHTMLGlob("templates/*")

	// define route handlers
	router.GET("/", func(c *gin.Context) {
		// call the HTML method of the Content to render a template
		c.HTML(
			// set http status code
			http.StatusOK,
			// use the index.html template
			"index.html",
			// pass the data that the page uses
			gin.H{
				"title": "Home Page",
			},
		)
	})

	// start to run
	router.Run()
}
