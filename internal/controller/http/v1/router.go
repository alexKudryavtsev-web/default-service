package v1

import "github.com/gin-gonic/gin"

func NewRouter(engine *gin.Engine) {
	router := engine.Group("/api")

	newCommonRoutes(router)
}