package private

type CancelRequest struct {
	OrderID string `json:"order_id" mapstructure:"order_id"`
}