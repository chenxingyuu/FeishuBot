package models

import "github.com/tietiexx/bot/code/backend/constant"

type LarkBotListRequest struct {
}

type LarkBotInfoResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type LarkBotInfoListResponse []LarkBotInfoResponse

type LarkBotPaginationResponse struct {
	Count   int64                   `json:"count"`
	Results LarkBotInfoListResponse `json:"results"`
}

type LarkBotCreateRequest struct {
	LarkAppUUID string               `json:"lark_app_uuid"`
	Name        string               `json:"name"`
	LarkBotType constant.LakeBotType `json:"lark_bot_type"`
}
