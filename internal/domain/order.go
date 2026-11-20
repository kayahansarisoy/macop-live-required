package domain

// Order is a customer order.
type Order struct {
	ID    string
	Total int64
}

// Payable reports whether the order still has an amount to collect.
func (o Order) Payable() bool {
	return o.Total > 0
}
