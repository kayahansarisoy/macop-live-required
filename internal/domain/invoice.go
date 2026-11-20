package domain

import "github.com/kayahansarisoy/macop-live-required/internal/infrastructure"

// Invoice is the bill of one order.
type Invoice struct {
	OrderID string
	Total   int
}

// StorageKey returns where the invoice is stored.
func (i Invoice) StorageKey() string {
	return infrastructure.Key(i.OrderID) + "/invoice"
}
