package converters

import (
	"github.com/tietiexx/bot/code/backend/database"
	"github.com/tietiexx/bot/code/backend/models"
)

func ToLarkAppInfoResponse(larkApp *database.LarkApp) models.LarkAppInfoResponse {
	larkAppInfoResponse := models.LarkAppInfoResponse{
		UUID: larkApp.UUID,
		Name: larkApp.Name,
	}
	return larkAppInfoResponse
}

func ToLarkAppListResponse(list []*database.LarkApp) (listResponse models.LarkAppInfoListResponse) {
	for _, larkApp := range list {
		listResponse = append(listResponse, ToLarkAppInfoResponse(larkApp))
	}
	return
}

func ToLarkAppPaginationResponse(count int64, list []*database.LarkApp) models.LarkAppPaginationResponse {
	listResponse := ToLarkAppListResponse(list)
	return models.LarkAppPaginationResponse{Count: count, Results: listResponse}
}
