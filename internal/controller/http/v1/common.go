package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func newCommonRoutes(handler *gin.RouterGroup) {
	handler.GET("/healthz", doHealthz)
}

func doHealthz(c *gin.Context) {
	c.Status(http.StatusOK)
}
