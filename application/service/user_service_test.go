package service

import (
	"context"
	"go-di-sample/mock/api/jsonplaceholder"
	"go-di-sample/utils/test"
	"testing"
)

func TestMain(m *testing.M) {
	test.RunWithProjectRoot(m, "../../")
}

func TestUserService_GetUserList(t *testing.T) {
	mockRepo := jsonplaceholder.NewMockClient()
	s := NewUserService(mockRepo)

	// ファイル内容をコンテキストに格納
	ctx := context.Background()
	ctx = context.WithValue(ctx, "path", "mock/api/jsonplaceholder/fixture/user_list.json")

	users, err := s.GetUserList(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(users) != 2 {
		t.Errorf("expected 2 but got %d", len(users))
	}

	if users[0].ID != 1 {
		t.Errorf("expected 1 but got %d", users[0].ID)
	}

	if users[1].ID != 2 {
		t.Errorf("expected 2 but got %d", users[1].ID)
	}
}

func TestUserService_GetUserById(t *testing.T) {
	mockRepo := jsonplaceholder.NewMockClient()
	s := NewUserService(mockRepo)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "path", "mock/api/jsonplaceholder/fixture/user.json")

	user, err := s.GetUserById(ctx, "1")
	if err != nil {
		t.Fatal(err)
	}

	if user.ID != 1 {
		t.Errorf("expected 1 but got %d", user.ID)
	}
}
