package service_test

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"blog-api-clean-architecture/internal/domain/models"
	mock_repository "blog-api-clean-architecture/internal/domain/repository/mock"
	"blog-api-clean-architecture/internal/domain/service"

	"go.uber.org/mock/gomock"
)

func TestSave(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	inputUser := &models.User{
		Name:  "taro.yamada",
		Email: "taro.yamada@test.com",
	}

	type args struct {
		user *models.User
	}
	tests := []struct {
		name                 string
		args                 args
		expectedMockBehavior func(dbRepo *mock_repository.MockIDBRepository)
		wantErr              bool
	}{
		{
			name: "正常系: ユーザーを保存できる場合",
			args: args{user: inputUser},
			expectedMockBehavior: func(dbRepo *mock_repository.MockIDBRepository) {
				dbRepo.EXPECT().Set(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "異常系: リポジトリでエラーが発生した場合",
			args: args{user: inputUser},
			expectedMockBehavior: func(dbRepo *mock_repository.MockIDBRepository) {
				dbRepo.EXPECT().Set(gomock.Any(), gomock.Any()).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbRepo := mock_repository.NewMockIDBRepository(ctrl)
			if tt.expectedMockBehavior != nil {
				tt.expectedMockBehavior(dbRepo)
			}

			userService := service.NewUserService(dbRepo)

			got, err := userService.Save(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				if got != nil {
					t.Errorf("Save() = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Error("Save() returned nil, want non-nil")
				return
			}
			if got.ID == "" {
				t.Error("Save() returned user with empty ID")
			}
			if got.Name != tt.args.user.Name {
				t.Errorf("Save() Name = %v, want %v", got.Name, tt.args.user.Name)
			}
			if got.Email != tt.args.user.Email {
				t.Errorf("Save() Email = %v, want %v", got.Email, tt.args.user.Email)
			}
			if got.CreatedAt == 0 {
				t.Error("Save() returned user with zero CreatedAt")
			}
		})
	}
}

func TestFindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existingUser := &models.User{
		ID:        "uuid",
		Name:      "taro.yamada",
		Email:     "taro.yamada@test.com",
		CreatedAt: 1000000,
	}
	existingUserJSON, _ := json.Marshal(existingUser)
	existingUserJSONStr := string(existingUserJSON)

	invalidUserData := "invalid_user_data"

	type args struct {
		id string
	}
	tests := []struct {
		name                 string
		args                 args
		expectedMockBehavior func(dbRepo *mock_repository.MockIDBRepository)
		wantUser             *models.User
		wantErr              bool
	}{
		{
			name: "正常系: ユーザーが存在する場合",
			args: args{id: "uuid"},
			expectedMockBehavior: func(dbRepo *mock_repository.MockIDBRepository) {
				dbRepo.EXPECT().Get("uuid").Return(&existingUserJSONStr, nil)
			},
			wantUser: existingUser,
			wantErr:  false,
		},
		{
			name: "正常系: ユーザーが存在しない場合",
			args: args{id: "not_found_uuid"},
			expectedMockBehavior: func(dbRepo *mock_repository.MockIDBRepository) {
				dbRepo.EXPECT().Get("not_found_uuid").Return(nil, nil)
			},
			wantUser: nil,
			wantErr:  false,
		},
		{
			name: "異常系: リポジトリでエラーが発生した場合",
			args: args{id: "uuid"},
			expectedMockBehavior: func(dbRepo *mock_repository.MockIDBRepository) {
				dbRepo.EXPECT().Get("uuid").Return(nil, errors.New("db error"))
			},
			wantUser: nil,
			wantErr:  true,
		},
		{
			name: "異常系: リポジトリから取得したユーザーが不正な形式である場合",
			args: args{id: "uuid"},
			expectedMockBehavior: func(dbRepo *mock_repository.MockIDBRepository) {
				dbRepo.EXPECT().Get("uuid").Return(&invalidUserData, nil) // go1.26からは `new("invaliduserdata")` で書けるはず
			},
			wantUser: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbRepo := mock_repository.NewMockIDBRepository(ctrl)
			if tt.expectedMockBehavior != nil {
				tt.expectedMockBehavior(dbRepo)
			}

			userService := service.NewUserService(dbRepo)

			got, err := userService.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.wantUser) {
				t.Errorf("FindByID() = %v, want %v", got, tt.wantUser)
			}
		})
	}
}
