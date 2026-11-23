package domain

// GramsToKilos converts a parcel weight from grams to kilograms.
func GramsToKilos(grams int) float64 {
        return float64(grams)/1000
}
