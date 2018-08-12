package middleware

import (
	"github.com/gin-gonic/gin"
)

type Handler func() (statusCode int, err error, output interface{})

func Warp(f Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
