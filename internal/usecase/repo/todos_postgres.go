package repo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/alexKudryavtsev-web/default-service/internal/entity"
	"github.com/alexKudryavtsev-web/default-service/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

const (
	_defaultListCap = 64
)

type TodosRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *TodosRepo {
	return &TodosRepo{
		pg,
	}
}

func (t *TodosRepo) GetAllTodos(ctx context.Context) ([]entity.Todo, error) {
	query, _, err := t.Builder.Select("id", "task").From("todos").ToSql()
	if err != nil {
		return []entity.Todo{}, fmt.Errorf("can't create sql query: %w", err)
	}

	rows, err := t.Pool.Query(ctx, query)
	if err != nil {
		return []entity.Todo{}, fmt.Errorf("can't query request: %w", err)
	}
	defer rows.Close()

	todos := make([]entity.Todo, 0, _defaultListCap)

	for rows.Next() {
		todo := entity.Todo{}

		err := rows.Scan(&todo.ID, &todo.Task)
		if err != nil {
			return []entity.Todo{}, fmt.Errorf("can't scan rows: %w", err)
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (t *TodosRepo) GetTodoByID(ctx context.Context, id uint64) (*entity.Todo, error) {
	query, args, err := t.Builder.
		Select("id", "task").
		From("todos").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("can't create sql query: %w", err)
	}

	row := t.Pool.QueryRow(ctx, query, args...)

	result := entity.Todo{}

	err = row.Scan(&result.ID, &result.Task)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entity.ErrTodoNotFound
		}

		return nil, fmt.Errorf("can't to scan row: %w", err)
	}

	return &result, nil
}

func (t *TodosRepo) SaveTodo(ctx context.Context, task string) error {
	query, args, err := t.Builder.
		Insert("todos").
		Columns("task").
		Values(task).
		ToSql()
	if err != nil {
		return fmt.Errorf("can't create sql query: %w", err)
	}

	if _, err = t.Pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("can't insert into table: %w", err)
	}

	return nil
}
