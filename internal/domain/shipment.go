package domain

import "github.com/kayahansarisoy/macop-live-required/internal/infrastructure"

// Shipment is the delivery of one order.
type Shipment struct {
	OrderID string
	Carrier string
}

// StorageKey returns where the shipment is stored.
func (s Shipment) StorageKey() string {
	return infrastructure.Key(s.OrderID) + "/shipment"
}
