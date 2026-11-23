package domain

// MetresToKilometres converts a route length from metres to kilometres.
func MetresToKilometres(metres int) float64 {
	return float64(metres) / 1000
}
