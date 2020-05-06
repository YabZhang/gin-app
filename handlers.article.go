package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func getArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := getArticleByID(articleID); err == nil {
			// call the HTML method of the Content to render a template
			c.HTML(
				// set http status code
				http.StatusOK,
				// use the index.html template
				"article.html",
				// pass the data that the page uses
				gin.H{
					"title": article.Title,
					"payload": article,
				},
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithError(http.StatusNotFound, err)
	}
}