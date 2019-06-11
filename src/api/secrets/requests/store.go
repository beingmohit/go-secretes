package requests 

type StoreRequest struct {
	SecretText string `json:"secretText"`
	ExpireAfterViews int32 `json:"expireAfterViews"`
	ExpireAfter int32 `json:"expireAfter"`
}