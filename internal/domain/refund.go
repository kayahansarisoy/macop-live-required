package domain

import "github.com/kayahansarisoy/macop-live-required/internal/infrastructure"

// Refund records money returned to a customer for an order.
type Refund struct {
	OrderID string
	Amount  int64
}

// StorageKey returns where the refund record is stored.
func (r Refund) StorageKey() string {
	return infrastructure.Key(r.OrderID) + "/refund"
}
