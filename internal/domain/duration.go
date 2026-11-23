package domain

// MinutesToHours converts a delivery time from minutes to hours.
func MinutesToHours(minutes int) float64 {
	return float64(minutes) / 60
}
