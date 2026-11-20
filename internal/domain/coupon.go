package domain

import "github.com/kayahansarisoy/macop-live-required/internal/infrastructure"

// Coupon is a discount applied to one order.
type Coupon struct {
	OrderID string
	Percent int
}

// StorageKey returns where the coupon is stored.
func (c Coupon) StorageKey() string {
	return infrastructure.Key(c.OrderID) + "/coupon"
}
