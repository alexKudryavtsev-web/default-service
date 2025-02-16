package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/alexKudryavtsev-web/default-service/docs"
)

func newCommonRoutes(handler *gin.RouterGroup) {
	handler.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	handler.GET("/healthz", doHealthz)
}

// @Summary     Check server healthz
// @ID          healthz
// @Tags  	    common
// @Success     200
// @Failure     500
// @Router      /healthz [get]
func doHealthz(c *gin.Context) {
	c.Status(http.StatusOK)
}
