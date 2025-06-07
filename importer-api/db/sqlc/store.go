package db

import (
	"context"
	"database/sql"
)

type Store interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetUserById(ctx context.Context, id int32) (User, error)
	GetFornecedoresById(ctx context.Context, id int32) (FornecedoresDado, error)
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
