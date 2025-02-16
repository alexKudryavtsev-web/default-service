package v1

import (
	"github.com/alexKudryavtsev-web/default-service/internal/usecase"
	"github.com/alexKudryavtsev-web/default-service/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, logger logger.Interface, todoUseCase usecase.Todos) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	router := handler.Group("/api")

	newCommonRoutes(router)
	newTodosRoutes(router, logger, todoUseCase)
}
