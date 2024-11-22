package jsonplaceholder

import (
	"context"
	"encoding/json"
	"go-di-sample/domain/model"
	"go-di-sample/domain/repository"
	"os"
)

type MockClient struct{}

func NewMockClient() repository.UserRepository {
	return &MockClient{}
}

func (c *MockClient) GetUserList(ctx context.Context) ([]*model.User, error) {
	// コンテキストの値からファイルを取得
	path, ok := ctx.Value("path").(string)
	if !ok {
		return nil, nil
	}
	// fixtureデータを取得してコンテキストに格納する
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []*model.User
	if err := json.NewDecoder(file).Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func (c *MockClient) GetUserById(ctx context.Context, id string) (*model.User, error) {
	path, ok := ctx.Value("path").(string)
	if !ok {
		return nil, nil
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var user model.User
	if err := json.NewDecoder(file).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
