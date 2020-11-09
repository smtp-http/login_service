package dao

import (
	"testing"

	"login_service/model"
)

func TestUserDAOMongo_FindUser(t *testing.T) {
	tests := []struct {
		name     string
		username string
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"Valid user find", "vbansal", false},
		{"Invalid user find", "vbans", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uDao := GetUserDAOMongoInstance()
			_, err := uDao.FindUser(tt.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserDAOMongo.FindUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserDAOMongo_AddNewUser(t *testing.T) {
	tests := []struct {
		name    string
		user    *model.UserModel
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Add user", &model.UserModel{
			Username:  "vbansal2",
			Password:  "Test123",
			Email:     "test@gmail.com",
			FirstName: "VarunTest",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uDao := GetUserDAOMongoInstance()
			if err := uDao.AddNewUser(tt.user); (err != nil) != tt.wantErr {
				t.Errorf("UserDAOMongo.AddNewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
