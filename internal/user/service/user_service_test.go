package service

import (
	"github.com/golang/mock/gomock"
	mock_repository "save-a-buddy-api/internal/mocks"
	"save-a-buddy-api/model"
	"testing"
)

func TestFinUsersIsEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repository.NewMockIUserRepository(ctrl)
	userService := New(mockRepository)
	mockRepository.EXPECT().FindUsersDb().Return(model.Users{}, nil)

	users, err := userService.FindUsers()

	if err != nil {
		t.Error("Error happens with find users")
	}

	if len(users) == 0 {
		t.Log("Find Users success is an empty list")
	}
}
