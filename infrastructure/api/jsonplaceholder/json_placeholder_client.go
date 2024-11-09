package jsonplaceholder

import (
	"context"
	"encoding/json"
	"go-di-sample/domain/model"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient JsonPlaceHolderClientのコンストラクタ
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
		baseURL:    "https://jsonplaceholder.typicode.com",
	}
}

// GetUserList ユーザ一覧を取得する
func (c *Client) GetUserList(ctx context.Context) ([]*model.User, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var users []*model.User
	if err = json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserById IDに紐づくユーザ情報を取得する
func (c *Client) GetUserById(ctx context.Context, id string) (*model.User, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/users/"+id, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var user model.User
	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
