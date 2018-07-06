package math

// Calculate average number
func Average(xs []float64) float64 {
	sum := 0.0
	for _, v := range xs {
		sum += v
	}
	return sum / float64(len(xs))
}

// A function that can be used
func CanUse() {
}

// A function that cannot be used
func cannotUse() {
}
