package HTTP

import "github.com/gin-gonic/gin"

var G *gin.Engine

func Ser() {
	g := gin.Default()
	G = g
	InitRouter()
	g.Run(":2060")
}
func InitRouter() {
	G.GET("/callback", func(c *gin.Context) {
		c.Redirect()
	})
}
