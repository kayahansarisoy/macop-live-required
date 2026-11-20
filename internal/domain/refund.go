package domain

// Refund records money returned to a customer for an order.
type Refund struct {
	OrderID string
	Amount  int64
}

// StorageKey returns where the refund record is stored, given where its order is stored.
func (r Refund) StorageKey(orderKey string) string {
	return orderKey + "/refund"
}
