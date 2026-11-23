package domain

// LitresOf returns the volume of a box in litres from its sides in centimetres.
func LitresOf(width, height, depth int) float64 {
        return float64(width*height*depth)/1000
}
