package util

import "strconv"

func ConvertStringToIntPointer(value string) (*int, error) {
	if value == "" {
		return nil, nil
	}

	result, err := strconv.Atoi(value)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
