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
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := todos.Todos(context.Background())

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}

}
