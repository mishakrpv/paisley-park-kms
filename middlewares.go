package main

import (
	"github.com/gin-gonic/gin"
)

func RequireAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
