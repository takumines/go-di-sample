package jsonplaceholder

import (
	"context"
	"encoding/json"
	"go-di-sample/domain/model"
	"go-di-sample/domain/repository"
	"log"
	"os"
)

type MockClient struct{}

func NewMockClient() repository.UserRepository {
	return &MockClient{}
}

func (c *MockClient) GetUserList(ctx context.Context) ([]*model.User, error) {
	file, err := os.Open("mock/api/jsonplaceholder/fixture/user_list.json")
	if err != nil {
		return nil, err
	}
	log.Println("file", file)

	defer file.Close()

	var users []*model.User
	if err = json.NewDecoder(file).Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func (c *MockClient) GetUserById(ctx context.Context, id string) (*model.User, error) {
	file, err := os.Open("mock/api/jsonplaceholder/fixture/user.json")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var user model.User
	if err = json.NewDecoder(file).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
