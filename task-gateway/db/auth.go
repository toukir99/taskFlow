package db

import (
	"log/slog"
	"task-gateway/db"
	"task-gateway/web/model"
)

type AuthRepo struct {
	authTable string
}

var authRepo *AuthRepo

func initAuthRepo() {
	authRepo = &AuthRepo{
		authTable: "auth",
	}
}

func GetAuthRepo() *AuthRepo {
	return authRepo
}

func (r *AuthRepo) CreateUser(user *model.User) (int, error) {
	query, args, err := GetQueryBuilder().
		Insert(r.authTable).
		Columns("email", "password", "is_active").
		Values(user.Email, user.Password, false).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		slog.Error("Failed to create user insert query", "err", err)
		return 0, err
	}

	var newId int
	err = db.GetWriteDB().QueryRow(query, args...).Scan(&newId)
	if err != nil {
		slog.Error("Error executing create user query", "err", err)
		return 0, err
	}

	return newId, nil
}
