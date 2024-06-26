package todoservice

import (
	"context"
	db "vbruzzi/todo-list/db/sqlc"
	errors "vbruzzi/todo-list/pkg/error"
)

type TodoServiceQueries interface {
	ListTodos(ctx context.Context) ([]db.Todo, error)
	CreateTodo(ctx context.Context, todo string) (db.Todo, error)
}

type TodoService struct {
	queries TodoServiceQueries
}

func (ts *TodoService) ListTodos() ([]db.Todo, *errors.Error) {
	res, err := ts.queries.ListTodos(context.Background())

	if err != nil {
		return nil, errors.NewError(nil, errors.EINTERNAL)
	}

	return res, nil
}

func (ts *TodoService) CreateTodo(todo string) (*db.Todo, *errors.Error) {
	if todo == "" {
		return nil, errors.NewError(nil, errors.EINVALID)
	}

	res, err := ts.queries.CreateTodo(context.Background(), todo)
	if err != nil {
		return nil, errors.NewError(err, errors.EINTERNAL)
	}

	return &res, nil
}

func NewTodoService(q TodoServiceQueries) *TodoService {
	return &TodoService{q}
}
