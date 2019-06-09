package secrets

type StoreRequest struct {
	Hash string `json:"hash" binding:"required"`
	SecretText string `json:"secretText" binding:"required"`
	CreatedAt string `json:"createdAt" binding:"required"`
	ExpiresAt string `json:"expiresAt" binding:"required"`
	RemainingViews uint64 `json:"remainingViews" binding:"required"`
}