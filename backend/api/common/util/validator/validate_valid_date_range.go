package validator

import (
	"errors"
	"sistem-manajemen-armada/api/common/constant"
)

func IsValidDateRange(start *int, end *int) error {
	if (start != nil && end == nil) || (start == nil && end != nil) {
		return errors.New("invalid date range, both start and end must be provided together")
	}

	if start != nil && end != nil {
		if *start > *end {
			return errors.New("invalid date range, end date must be greater than start date")
		}

		if *start < constant.MinimumTimestamp || *start > constant.MaximumTimestamp {
			return errors.New("invalid date range, start timestamp out of range")
		}

		if *end < constant.MinimumTimestamp || *end > constant.MaximumTimestamp {
			return errors.New("invalid date range, end timestamp out of range")
		}
	}

	return nil
}
