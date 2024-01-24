package helper

import "strconv"

// StringToFloat64
func StringToFloat64(str string) float64 {
	float, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}

	return float
}
