package service

import (
	"context"
	"go-di-sample/mock/api/jsonplaceholder"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// NOTE: テスト実行時はカレントディレクトリがテスト実行場所となるため、テストコード内でファイルを読み込む処理がある場合、
	//	参照できなくなるケースを回避するようカレントディレクトリをプロジェクトのルートに変更する

	// カレントディレクトリを取得
	originalDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current directory: %v", err)
	}

	// カレントディレクトリをプロジェクトのルートに変更
	if err := os.Chdir("../../"); err != nil {
		log.Fatalf("failed to change directory: %v", err)
	}

	// テスト完了後に元のディレクトリに戻す
	defer func() {
		if err := os.Chdir(originalDir); err != nil {
			log.Fatalf("failed to restore original directory: %v", err)
		}
	}()

	os.Exit(m.Run())
}

func TestUserService_GetUserList(t *testing.T) {
	mockRepo := jsonplaceholder.NewMockClient()

	s := NewUserService(mockRepo)

	users, err := s.GetUserList(context.Background())
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

	user, err := s.GetUserById(context.Background(), "1")
	if err != nil {
		t.Fatal(err)
	}

	if user.ID != 1 {
		t.Errorf("expected 1 but got %d", user.ID)
	}
}
