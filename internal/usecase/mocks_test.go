// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/usecase/interfaces.go

// Package usecase_test is a generated GoMock package.
package usecase_test

import (
	context "context"
	reflect "reflect"

	entity "github.com/alexKudryavtsev-web/default-service/internal/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockTodos is a mock of Todos interface.
type MockTodos struct {
	ctrl     *gomock.Controller
	recorder *MockTodosMockRecorder
}

// MockTodosMockRecorder is the mock recorder for MockTodos.
type MockTodosMockRecorder struct {
	mock *MockTodos
}

// NewMockTodos creates a new mock instance.
func NewMockTodos(ctrl *gomock.Controller) *MockTodos {
	mock := &MockTodos{ctrl: ctrl}
	mock.recorder = &MockTodosMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodos) EXPECT() *MockTodosMockRecorder {
	return m.recorder
}

// SaveTodo mocks base method.
func (m *MockTodos) SaveTodo(ctx context.Context, task string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTodo", ctx, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTodo indicates an expected call of SaveTodo.
func (mr *MockTodosMockRecorder) SaveTodo(ctx, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTodo", reflect.TypeOf((*MockTodos)(nil).SaveTodo), ctx, task)
}

// TodoByID mocks base method.
func (m *MockTodos) TodoByID(ctx context.Context, id uint64) (*entity.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TodoByID", ctx, id)
	ret0, _ := ret[0].(*entity.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TodoByID indicates an expected call of TodoByID.
func (mr *MockTodosMockRecorder) TodoByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TodoByID", reflect.TypeOf((*MockTodos)(nil).TodoByID), ctx, id)
}

// Todos mocks base method.
func (m *MockTodos) Todos(ctx context.Context) ([]entity.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Todos", ctx)
	ret0, _ := ret[0].([]entity.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Todos indicates an expected call of Todos.
func (mr *MockTodosMockRecorder) Todos(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Todos", reflect.TypeOf((*MockTodos)(nil).Todos), ctx)
}

// MockTodosRepo is a mock of TodosRepo interface.
type MockTodosRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTodosRepoMockRecorder
}

// MockTodosRepoMockRecorder is the mock recorder for MockTodosRepo.
type MockTodosRepoMockRecorder struct {
	mock *MockTodosRepo
}

// NewMockTodosRepo creates a new mock instance.
func NewMockTodosRepo(ctrl *gomock.Controller) *MockTodosRepo {
	mock := &MockTodosRepo{ctrl: ctrl}
	mock.recorder = &MockTodosRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodosRepo) EXPECT() *MockTodosRepoMockRecorder {
	return m.recorder
}

// GetAllTodos mocks base method.
func (m *MockTodosRepo) GetAllTodos(ctx context.Context) ([]entity.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTodos", ctx)
	ret0, _ := ret[0].([]entity.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTodos indicates an expected call of GetAllTodos.
func (mr *MockTodosRepoMockRecorder) GetAllTodos(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTodos", reflect.TypeOf((*MockTodosRepo)(nil).GetAllTodos), ctx)
}

// GetTodoByID mocks base method.
func (m *MockTodosRepo) GetTodoByID(ctx context.Context, id uint64) (*entity.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodoByID", ctx, id)
	ret0, _ := ret[0].(*entity.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodoByID indicates an expected call of GetTodoByID.
func (mr *MockTodosRepoMockRecorder) GetTodoByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodoByID", reflect.TypeOf((*MockTodosRepo)(nil).GetTodoByID), ctx, id)
}

// SaveTodo mocks base method.
func (m *MockTodosRepo) SaveTodo(ctx context.Context, task string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTodo", ctx, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTodo indicates an expected call of SaveTodo.
func (mr *MockTodosRepoMockRecorder) SaveTodo(ctx, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTodo", reflect.TypeOf((*MockTodosRepo)(nil).SaveTodo), ctx, task)
}
