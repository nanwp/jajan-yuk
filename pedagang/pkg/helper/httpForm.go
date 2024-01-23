package helper

import "strconv"

func GetValue(m map[string][]string, key string, defaultValue string) string {
	if val, ok := m[key]; ok {
		return val[0]
	}
	return defaultValue
}

func StringToFloat(str string) float64 {
	float, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}

	return float
}

func StringToBool(str string) bool {
	bool, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}

	return bool
}
