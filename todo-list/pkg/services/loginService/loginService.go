package loginservice

import (
	"context"
	db "vbruzzi/todo-list/db/sqlc"
	errors "vbruzzi/todo-list/pkg/error"

	"github.com/jackc/pgx/v5/pgtype"
)

type LoginServiceQueries interface {
	GetUserByCredentials(ctx context.Context, params db.GetUserByCredentialsParams) (db.GetUserByCredentialsRow, error)
	GetUserIdFromLoginKey(ctx context.Context, key pgtype.Text) (int32, error)
	SetUserKey(ctx context.Context, params db.SetUserKeyParams) error
}

type LoginService struct {
	queries LoginServiceQueries
}

func (ts *LoginService) GetUserKey(username, password string) *errors.Error {
	_, err := ts.queries.GetUserByCredentials(context.Background(), db.GetUserByCredentialsParams{
		Username: username,
		Password: password,
	})

	if err != nil {
		// todo: need to handle 404
		return errors.NewError(err, errors.EINTERNAL)
	}

	return nil
}

func NewLoginService(q LoginServiceQueries) *LoginService {
	return &LoginService{q}
}
