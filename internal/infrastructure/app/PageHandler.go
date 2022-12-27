package app

import (
	"net/http"
	"wordleGame/internal/ui/views"

	"github.com/gin-gonic/gin"
)

func (s *ApplicationServer) HomePageHandler() func(*gin.Context) {
	// return func(c *gin.Context) {
	// 	// routePath := strings.Split(c.Request.URL.Path, "/")
	// 	// fmt.Println(routePath)
	// 	//c.Writer.WriteString("Test")
	// 	c.HTML(http.StatusOK, "web/components/Board.go", gin.H{
	// 		"title": "hello world",
	// 	})
	// }
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
		_ = views.HomeView(c.DefaultQuery("name", "Guest")).Render(c.Writer)
	}
}
