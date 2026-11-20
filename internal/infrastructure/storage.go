// Package infrastructure maps domain records to where they are stored.
package infrastructure

// Key returns the storage key of an order.
func Key(id string) string {
	return "orders/" + id
}
