package handlers

import (
	"github.com/labstack/echo/v4"
	"go-di-sample/application/service"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler UserHandlerのコンストラクタ
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetUserList ユーザ一覧を取得する
func (h *UserHandler) GetUserList(c echo.Context) error {
	users, err := h.userService.GetUserList(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

// GetUserById IDに紐づくユーザ情報を取得する
func (h *UserHandler) GetUserById(c echo.Context) error {
	id := c.Param("id")
	user, err := h.userService.GetUserById(c.Request().Context(), id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "server error"})
	}

	// IDが0の場合は、ユーザーが存在しないとして404を返す
	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	return c.JSON(http.StatusOK, user)
}
