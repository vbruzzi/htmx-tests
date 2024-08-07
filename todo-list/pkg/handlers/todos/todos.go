package todos

import (
	"net/http"
	"strconv"
	db "vbruzzi/todo-list/db/sqlc"
	errors "vbruzzi/todo-list/pkg/error"
	todoservice "vbruzzi/todo-list/pkg/services/todoService"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	todoService *todoservice.TodoService
}

type List struct{}

type FormErrors struct {
	Todo string
}

// todo: move to models.go
type FormValues struct {
	Id   int
	Todo string
	Done bool
}

type Form struct {
	Values FormValues
	Errors FormErrors
}

func newTodo(id int, todo string) FormValues {
	return FormValues{
		id,
		todo,
		false,
	}
}

func (h *TodoHandler) getTodos(c echo.Context) error {
	todos, err := h.todoService.ListTodos()

	if err != nil {
		return err.Err
	}

	return c.Render(http.StatusOK, "todoList", todos)
}

func (h *TodoHandler) createTodo(c echo.Context) error {
	todo := c.FormValue("value")
	newEntry, err := h.todoService.CreateTodo(todo)

	if err != nil {
		if err.Code != errors.EINVALID {
			return err.Err
		}

		res := Form{
			newTodo(0, todo),
			FormErrors{"Value cannot be empty"},
		}

		return c.Render(err.Status, "todoForm", res)
	}

	c.Render(http.StatusOK, "todoForm", nil)
	return c.Render(http.StatusOK, "oobItem", newEntry)
}

func (h *TodoHandler) deleteTodo(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr[1:])

	if err != nil {
		return c.Render(http.StatusOK, "error", nil)
	}

	deleteErr := h.todoService.DeleteTodo(int32(id))

	if deleteErr != nil {
		return c.Render(http.StatusOK, "error", nil)
	}

	return c.NoContent(http.StatusOK)
}

func NewTodoHandler(g *echo.Group, q *db.Queries) {
	handler := &TodoHandler{
		todoservice.NewTodoService(q),
	}

	g.GET("", handler.getTodos)
	g.POST("", handler.createTodo)
	g.DELETE(":id", handler.deleteTodo)
}
