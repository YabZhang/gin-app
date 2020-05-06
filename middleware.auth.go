package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setUserStatus() gin.HandlerFunc {
	return func (c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil && token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}}

func ensureLoggedIn() gin.HandlerFunc {
	return func (c *gin.Context) {
		authIntface, _ := c.Get("is_logged_in")
		if isLoggedIn, _ := authIntface.(bool); !isLoggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func ensureNotLoggedIn() gin.HandlerFunc {
	return func (c *gin.Context) {
		authIntface, _ := c.Get("is_logged_in")
		if isLoggedIn, _ := authIntface.(bool); isLoggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}