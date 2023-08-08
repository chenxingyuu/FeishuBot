package models

type LarkAppListRequest struct {
}

type LarkAppInfoResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type LarkAppInfoListResponse []LarkAppInfoResponse

type LarkAppPaginationResponse struct {
	Count   int64                   `json:"count"`
	Results LarkAppInfoListResponse `json:"results"`
}

type LarkAppCreateRequest struct {
	Name string `json:"name" binding:"required"`
}
