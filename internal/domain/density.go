package domain

// GramsPerLitre returns how dense a parcel is from its weight and its volume in litres.
func GramsPerLitre(grams int, litres float64) float64 {
        return float64(grams)/litres
}
