package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *ApplicationServer) err404PageHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
