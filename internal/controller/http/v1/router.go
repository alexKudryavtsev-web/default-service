package v1

import "github.com/gin-gonic/gin"

func NewRouter(handler *gin.Engine) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	router := handler.Group("/api")

	newCommonRoutes(router)
}