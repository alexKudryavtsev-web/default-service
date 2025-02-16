package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/alexKudryavtsev-web/default-service/internal/entity"
	"github.com/alexKudryavtsev-web/default-service/internal/usecase"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var errInternalServErr = errors.New("internal server error")
var errTodoNotFound = entity.ErrTodoNotFound

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func todos(t *testing.T) (*usecase.TodosUseCase, *MockTodosRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockTodosRepo(mockCtl)
	todos := usecase.NewTodosUseCase(repo)

	return todos, repo
}

func TestTodos(t *testing.T) {
	t.Parallel()

	todos, repo := todos(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				repo.EXPECT().GetAllTodos(context.Background()).Return(nil, nil)
			},
			res: []entity.Todo(nil),
			err: nil,
		},
		{
			name: "result with error",
			mock: func() {
				repo.EXPECT().GetAllTodos(context.Background()).Return(nil, errInternalServErr)
			},
			res: []entity.Todo(nil),
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := todos.Todos(context.Background())

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestTodoByID(t *testing.T) {
	t.Parallel()

	todos, repo := todos(t)

	tests := []struct {
		name   string
		todoID uint64
		mock   func(uint64)
		res    *entity.Todo
		err    error
	}{
		{
			name:   "todo found",
			todoID: 1,
			mock: func(id uint64) {
				expectedTodo := &entity.Todo{ID: id, Task: "Test Task"}
				repo.EXPECT().GetTodoByID(context.Background(), id).Return(expectedTodo, nil)
			},
			res: &entity.Todo{ID: 1, Task: "Test Task"},
			err: nil,
		},
		{
			name:   "todo not found",
			todoID: 2,
			mock: func(id uint64) {
				repo.EXPECT().GetTodoByID(context.Background(), id).Return(nil, errTodoNotFound)
			},
			res: nil,
			err: errTodoNotFound,
		},
		{
			name:   "internal server error",
			todoID: 3,
			mock: func(id uint64) {
				repo.EXPECT().GetTodoByID(context.Background(), id).Return(nil, errInternalServErr)
			},
			res: nil,
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock(tc.todoID)

			res, err := todos.TodoByID(context.Background(), tc.todoID)

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestSaveTodo(t *testing.T) {
	t.Parallel()

	todos, repo := todos(t)

	tests := []test{
		{
			name: "successful save",
			mock: func() {
				repo.EXPECT().SaveTodo(context.Background(), "New Task").Return(nil)
			},
			res: nil,
			err: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			err := todos.SaveTodo(context.Background(), "New Task")

			require.ErrorIs(t, err, tc.err)
		})
	}
}
