package test

import (
	"log"
	"os"
	"testing"
)

func RunWithProjectRoot(m *testing.M, projectRoot string) {
	originalDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current directory: %v", err)
	}

	// カレントディレクトリをプロジェクトのルートに変更
	if err := os.Chdir(projectRoot); err != nil {
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
