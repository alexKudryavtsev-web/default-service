package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/alexKudryavtsev-web/default-service/internal/entity"
	"github.com/alexKudryavtsev-web/default-service/internal/usecase"
	"github.com/alexKudryavtsev-web/default-service/pkg/logger"
	"github.com/gin-gonic/gin"
)

type todosRoutes struct {
	u usecase.Todos
	l logger.Interface
}

func newTodosRoutes(handler *gin.RouterGroup, logger logger.Interface, todosUseCase usecase.Todos) {
	r := todosRoutes{todosUseCase, logger}

	handler.GET("/todos", r.doGetAllTodos)
	handler.GET("/todos/:id", r.doGetTodoByID)
	handler.POST("/todos", r.doSaveTodo)
}

// @Summary     Get all toos
// @Description Get all todos
// @ID          get-all-todos
// @Tags  	    todos
// @Accept      json
// @Success     200
// @Failure     500
// @Produce     json
// @Router      /todos/ [get]
func (t *todosRoutes) doGetAllTodos(ctx *gin.Context) {
	todos, err := t.u.Todos(ctx.Request.Context())

	if err != nil {
		t.l.Error(err, "http - v1 - doGetAllAdmins")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": todos,
	})
}

// @Summary     Get todo by id
// @Description Get todo by id
// @ID          get-todo-by-id
// @Tags  	    todos
// @Param       id   path      int  true  "Todo ID"
// @Accept      json
// @Success     200
// @Failure     400
// @Failure     404
// @Failure     500
// @Produce     json
// @Router      /todos/{id} [get]
// @Security    BearerAuth
func (t *todosRoutes) doGetTodoByID(ctx *gin.Context) {
	id := ctx.Param("id")

	todoID, err := strconv.Atoi(id)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid ID")

		return
	}

	todo, err := t.u.TodoByID(ctx.Request.Context(), uint64(todoID))

	if errors.Is(err, entity.ErrTodoNotFound) {
		errorResponse(ctx, http.StatusNotFound, "todo not found")

		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": todo,
	})
}

type doSaveTodoRequest struct {
	Task string `json:"task" binding:"required"`
}

// @Summary     Create todo
// @Description Create todo
// @ID          create-todo
// @Tags  	    todos
// @Param 			request body doSaveTodoRequest true "query params"
// @Accept      json
// @Success     200
// @Failure     500
// @Produce     json
// @Router      /todos [post]
func (t *todosRoutes) doSaveTodo(ctx *gin.Context) {
	var request doSaveTodoRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		t.l.Error(err, "http - v1 - doSaveTodo")
		errorResponse(ctx, http.StatusBadRequest, "invalid request body")

		return
	}

	if err := t.u.SaveTodo(ctx.Request.Context(), request.Task); err != nil {
		t.l.Error(err, "http - v1 - doSaveTodo")
		errorResponse(ctx, http.StatusInternalServerError, "internal service problems")

		return
	}

	ctx.JSON(http.StatusOK, nil)
}
