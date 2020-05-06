package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// call the HTML method of the Content to render a template
	c.HTML(
		// set http status code
		http.StatusOK,
		// use the index.html template
		"index.html",
		// pass the data that the page uses
		gin.H{
			"title": "Home Page",
			"payload": articles,
		},
	)
}