package helper

import "strconv"

func ConvertStringSliceToInt64Slice(stringSlice []string) (int64Slice []int64, err error) {
	for _, stringID := range stringSlice {
		intID, err := strconv.ParseInt(stringID, 10, 64)
		if err != nil {
			return []int64{}, err
		}

		int64Slice = append(int64Slice, intID)
	}

	return int64Slice, nil
}
