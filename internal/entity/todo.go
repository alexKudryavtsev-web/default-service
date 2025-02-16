package entity

import "errors"

type Todo struct {
	ID   uint64    `json:"id"`
	Task string `json:"task"`
}

var (
	ErrTodoNotFound = errors.New("todo not found")
)
