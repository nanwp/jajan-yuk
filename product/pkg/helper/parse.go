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

func StringToInt64(str string) int64 {
	integer, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}

	return integer
}
