package converters

import (
	"github.com/tietiexx/bot/code/backend/database"
	"github.com/tietiexx/bot/code/backend/models"
)

func ToLarkBotInfoResponse(larkApp *database.LarkBot) models.LarkBotInfoResponse {
	larkAppInfoResponse := models.LarkBotInfoResponse{
		UUID: larkApp.UUID,
		Name: larkApp.Name,
	}
	return larkAppInfoResponse
}

func ToLarkBotListResponse(list []*database.LarkBot) (listResponse models.LarkBotInfoListResponse) {
	for _, larkApp := range list {
		listResponse = append(listResponse, ToLarkBotInfoResponse(larkApp))
	}
	return
}

func ToLarkBotPaginationResponse(count int64, list []*database.LarkBot) models.LarkBotPaginationResponse {
	listResponse := ToLarkBotListResponse(list)
	return models.LarkBotPaginationResponse{Count: count, Results: listResponse}
}
