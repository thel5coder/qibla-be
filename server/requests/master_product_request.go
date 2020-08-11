package requests

type MasterProductRequest struct {
	Name             string `json:"name"`
	SubscriptionType string `json:"subscription_type"`
}
