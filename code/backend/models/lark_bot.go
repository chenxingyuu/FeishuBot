package models

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
