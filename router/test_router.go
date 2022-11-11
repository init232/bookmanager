package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.GET("/api/test", TestHandler)
}

func TestHandler(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
