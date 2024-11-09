package repository

import (
	"context"
	"go-di-sample/domain/model"
)

type UserRepository interface {
	GetUserList(ctx context.Context) ([]*model.User, error)
	GetUserById(ctx context.Context, id string) (*model.User, error)
}
