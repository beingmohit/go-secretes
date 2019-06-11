package requests 

type StoreRequest struct {
	SecretText string `json:"secretText" binding:"required"`
	ExpireAfterViews int32 `json:"expireAfterViews"`
	ExpireAfter int32 `json:"expireAfter"`
}