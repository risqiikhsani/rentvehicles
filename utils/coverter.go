package utils

import "strconv"

func ConvertToUint(str string) (uint, error) {
	value, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}
