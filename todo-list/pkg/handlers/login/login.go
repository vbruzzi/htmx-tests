package login

import (
	"crypto"
	"fmt"
	"net/http"
	db "vbruzzi/todo-list/db/sqlc"
	loginservice "vbruzzi/todo-list/pkg/services/loginService"

	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	loginService *loginservice.LoginService
}

type FormErrors struct {
	Todo string
}

type FormValues struct {
	Username string
	Password string
}

type Form struct {
	Values FormValues
	Errors FormErrors
}

func newLogin(username, password string) FormValues {
	return FormValues{
		username,
		password,
	}

}

func (h *LoginHandler) login(c echo.Context) error {
	salt := []byte{}

	for _, i := range "abc" {
		fmt.Println(i)
		salt = append(salt, byte(i))
	}

	h2 := crypto.MD5.New()

	return c.Render(http.StatusOK, "todoList", nil)
}

func NewLoginHandler(g *echo.Group, q *db.Queries) {
	handler := &LoginHandler{
		loginservice.NewLoginService(q),
	}

	g.POST("", handler.login)
}
