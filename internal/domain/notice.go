package domain

import "net/http"

// Notice tells a customer what happened to an order.
type Notice struct {
	OrderID string
	Status  int
}

// Delivered reports whether the notice reached the customer.
func (n Notice) Delivered() bool {
	return n.Status == http.StatusOK
}
