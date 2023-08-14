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

type LarkAppUpdatePartialRequest struct {
	Name              string `json:"name"`
	Status            uint8  `json:"status"`
	AppID             string `json:"app_id"`
	AppSecret         string `json:"app_secret"`
	EncryptKey        string `json:"encrypt_key"`
	VerificationToken string `json:"verification_token"`
}

type LarkAppUpdateRequest struct {
	Name              string `json:"name" binding:"required"`
	Status            uint8  `json:"status" binding:"required"`
	AppID             string `json:"app_id" binging:"required"`
	AppSecret         string `json:"app_secret" binging:"required"`
	EncryptKey        string `json:"encrypt_key" binging:"required"`
	VerificationToken string `json:"verification_token" binging:"required"`
}
