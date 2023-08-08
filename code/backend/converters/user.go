package converters

import (
	"github.com/tietiexx/bot/code/backend/database"
	"github.com/tietiexx/bot/code/backend/models"
)

func UserToUserInfoResponse(user *database.User) models.UserInfoResponse {
	userResponse := models.UserInfoResponse{
		ID:       user.ID,
		Username: user.Username,
		Status:   user.Status,
	}

	return userResponse
}
