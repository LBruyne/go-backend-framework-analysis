package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func redirect() *gin.Engine {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	return r
}
