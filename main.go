package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// set the default router
	router = gin.Default()

	// load templates
	router.LoadHTMLGlob("templates/*")

	// define route handlers
	initializeRoutes()

	// start to run
	router.Run()
}
