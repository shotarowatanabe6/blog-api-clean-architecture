package handler_test

import (
	"blog-api-clean-architecture/internal/domain/models"
	mock_service "blog-api-clean-architecture/internal/domain/service/mock"
	"blog-api-clean-architecture/internal/handler"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func TestUserHandler_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name                 string
		userID               string
		expectedStatus       int
		expectedMockBehavior func(mock *mock_service.MockIUserService)
	}{
		{
			name:                 "異常系：URIにidが含まれていない場合は Bad Request とする",
			userID:               "",
			expectedStatus:       http.StatusBadRequest,
			expectedMockBehavior: nil,
		},
		{
			name:           "異常系：Serviceがエラーを返した場合は Internal Server Error とする",
			userID:         "abc123",
			expectedStatus: http.StatusInternalServerError,
			expectedMockBehavior: func(mock *mock_service.MockIUserService) {
				mock.EXPECT().FindByID("abc123").Return(nil, fmt.Errorf("error"))
			},
		},
		{
			name:           "正常系：Serviceが正常に実施された場合は OK とする",
			userID:         "abc123",
			expectedStatus: http.StatusOK,
			expectedMockBehavior: func(mock *mock_service.MockIUserService) {
				mock.EXPECT().FindByID("abc123").Return(&models.User{ID: "abc123", Name: "taro.yamada", Email: "taro.yamada@example.com", CreatedAt: 1}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックの振る舞いを定義
			userService := mock_service.NewMockIUserService(ctrl)
			if tt.expectedMockBehavior != nil {
				tt.expectedMockBehavior(userService)
			}

			// テスト用のHTTPコンテキストを構築
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodGet, "/users/"+tt.userID, nil)
			if tt.userID != "" {
				c.Params = gin.Params{{Key: "id", Value: tt.userID}}
			}

			// テストを実施
			h := handler.NewUserHandler(userService)
			h.FindByID(c)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestUserHandler_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name                 string
		requestBody          string
		expectedStatus       int
		expectedMockBehavior func(mock *mock_service.MockIUserService)
	}{
		{
			name:                 "異常系：リクエストにnameが含まれていない場合は Bad Request とする",
			requestBody:          `{"email": "taro.yamada@example.com"}`,
			expectedStatus:       http.StatusBadRequest,
			expectedMockBehavior: nil,
		},
		{
			name:                 "異常系：リクエストにemailが含まれていない場合は Bad Request とする",
			requestBody:          `{"name": "taro.yamada"}`,
			expectedStatus:       http.StatusBadRequest,
			expectedMockBehavior: nil,
		},
		{
			name:           "異常系：リクエストにname,emailが含まれているが、Serviceがエラーを返した場合は Internal Server Error とする",
			requestBody:    `{"name": "taro.yamada", "email": "taro.yamada@example.com"}`,
			expectedStatus: http.StatusInternalServerError,
			expectedMockBehavior: func(mock *mock_service.MockIUserService) {
				mock.EXPECT().Save(&models.User{Name: "taro.yamada", Email: "taro.yamada@example.com"}).Return(nil, fmt.Errorf("error"))
			},
		},
		{
			name:           "正常系：リクエストにname,emailが含まれており、Serviceが正常に実施された場合は Created とする",
			requestBody:    `{"name": "taro.yamada", "email": "taro.yamada@example.com"}`,
			expectedStatus: http.StatusCreated,
			expectedMockBehavior: func(mock *mock_service.MockIUserService) {
				mock.EXPECT().Save(&models.User{Name: "taro.yamada", Email: "taro.yamada@example.com"}).Return(&models.User{ID: "1", Name: "taro.yamada", Email: "taro.yamada@example.com", CreatedAt: 1}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックの振る舞いを定義
			userService := mock_service.NewMockIUserService(ctrl)
			if tt.expectedMockBehavior != nil {
				tt.expectedMockBehavior(userService)
			}

			// テスト用のHTTPコンテキストを構築
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(tt.requestBody))
			c.Request.Header.Set("Content-Type", "application/json")

			// テストを実施
			h := handler.NewUserHandler(userService)
			h.Save(c)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}
