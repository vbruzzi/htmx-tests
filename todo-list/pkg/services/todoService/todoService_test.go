package todoservice_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	db "vbruzzi/todo-list/db/sqlc"
	err "vbruzzi/todo-list/pkg/error"
	todoservice "vbruzzi/todo-list/pkg/services/todoService"

	"github.com/stretchr/testify/assert"
)

type mockTodoServiceQueries struct{}

func (m *mockTodoServiceQueries) ListTodos(ctx context.Context) ([]db.Todo, error) {
	return []db.Todo{{ID: 1, Todo: "foo"}}, nil
}
func (m *mockTodoServiceQueries) CreateTodo(ctx context.Context, todo string) (db.Todo, error) {
	return db.Todo{ID: 1, Todo: todo}, nil
}
func (m *mockTodoServiceQueries) DeleteTodo(ctx context.Context, id int32) error {
	return nil
}

type mockFailedTodoServiceQueries struct{}

func (m *mockFailedTodoServiceQueries) ListTodos(ctx context.Context) ([]db.Todo, error) {
	return nil, errors.New("foo")
}
func (m *mockFailedTodoServiceQueries) CreateTodo(ctx context.Context, todo string) (db.Todo, error) {
	var t db.Todo
	return t, errors.New("bar")
}
func (m *mockFailedTodoServiceQueries) DeleteTodo(ctx context.Context, id int32) error {
	return errors.New("baz")
}

func createService(t *testing.T, testingErrs bool) (todoservice.TodoService, *assert.Assertions) {
	var q todoservice.TodoServiceQueries
	if testingErrs {
		q = &mockFailedTodoServiceQueries{}
	} else {
		q = &mockTodoServiceQueries{}
	}
	service := todoservice.NewTodoService(q)
	assert := assert.New(t)
	return *service, assert
}

func TestList(t *testing.T) {
	service, assert := createService(t, false)
	listRes, err := service.ListTodos()

	assert.Len(listRes, 1)
	assert.Nil(err)
}

func TestListErr(t *testing.T) {
	service, assert := createService(t, true)
	listRes, er := service.ListTodos()

	assert.NotNil(er)
	assert.Equal(er.Code, err.EINTERNAL)
	assert.Nil(listRes)
}

func TestCreate(t *testing.T) {
	service, assert := createService(t, false)
	createRes, err := service.CreateTodo("foo")

	assert.Nil(err)
	assert.Equal(createRes.ID, int32(1))
	assert.Equal(createRes.Todo, "foo")
}

func TestCreateValidation(t *testing.T) {
	service, assert := createService(t, false)
	createRes, er := service.CreateTodo("")

	assert.Nil(createRes)
	assert.Equal(er.Code, err.EINVALID)
	assert.Equal(er.Status, http.StatusUnprocessableEntity)
}

func TestCreateError(t *testing.T) {
	service, assert := createService(t, true)
	createRes, er := service.CreateTodo("foo")

	assert.Nil(createRes)
	assert.Equal(er.Code, err.EINTERNAL)
}

func TestDelete(t *testing.T) {
	service, assert := createService(t, false)
	er := service.DeleteTodo(1)

	assert.Nil(er)
}

func TestFailedDelete(t *testing.T) {
	service, assert := createService(t, true)
	er := service.DeleteTodo(1)

	assert.Equal(er.Code, err.EINTERNAL)
}
