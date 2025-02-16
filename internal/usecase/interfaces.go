package usecase

import (
	"context"

	"github.com/alexKudryavtsev-web/default-service/internal/entity"
)

type (
	Todos interface {
		Todos(ctx context.Context) ([]entity.Todo, error)
		TodoByID(ctx context.Context, id uint64) (*entity.Todo, error)
		SaveTodo(ctx context.Context, task string) error
	}

	TodosRepo interface {
		GetAllTodos(ctx context.Context) ([]entity.Todo, error)
		GetTodoByID(ctx context.Context, id uint64) (*entity.Todo, error)
		SaveTodo(ctx context.Context, task string) error
	}
)
