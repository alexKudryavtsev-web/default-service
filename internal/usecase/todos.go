package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexKudryavtsev-web/default-service/internal/entity"
)

type TodosUseCase struct {
	repo TodosRepo
}

var _ Todos = (*TodosUseCase)(nil)

func NewTodosUseCase(repo TodosRepo) *TodosUseCase {
	return &TodosUseCase{
		repo,
	}
}

func (t *TodosUseCase) SaveTodo(ctx context.Context, task string) error {
	err := t.repo.SaveTodo(ctx, task)

	if err != nil {
		return fmt.Errorf("can't save todo: %w", err)
	}

	return nil
}

func (t *TodosUseCase) TodoByID(ctx context.Context, id uint64) (*entity.Todo, error) {
	todo, err := t.repo.GetTodoByID(ctx, id)

	if errors.Is(err, entity.ErrTodoNotFound) {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("can't get todo by id: %w", err)
	}

	return todo, nil
}

func (t *TodosUseCase) Todos(ctx context.Context) ([]entity.Todo, error) {
	todos, err := t.repo.GetAllTodos(ctx)

	if err != nil {
		return nil, fmt.Errorf("can't get all todos: %w", err)
	}

	return todos, nil
}
