package app

import (
	"net/http"
	"wordleGame/internal/ui/views/components/errorPages"

	"github.com/gin-gonic/gin"
	g "github.com/maragudk/gomponents"
)

var errorPage = map[int]g.Node{
	400: errorPages.Error_400(),
	404: errorPages.Error_404(),
	500: errorPages.Error_500(),
}

func (s *ApplicationServer) err404PageHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (s *ApplicationServer) errhandler() func(*gin.Context) {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() >= 400 {
			val, ok := errorPage[c.Writer.Status()]
			if !ok {
				val = errorPages.Error_500()
			}
			val.Render(c.Writer)
		}

	}
}
