package service

import (
	"context"
	"go-di-sample/domain/model"
	"go-di-sample/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

// NewUserService UserServiceのコンストラクタ
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUserList ユーザ一覧を取得する
func (s *UserService) GetUserList(ctx context.Context) ([]*model.User, error) {
	return s.repo.GetUserList(ctx)
}

// GetUserById IDに紐づくユーザ情報を取得する
func (s *UserService) GetUserById(ctx context.Context, id string) (*model.User, error) {
	return s.repo.GetUserById(ctx, id)
}
