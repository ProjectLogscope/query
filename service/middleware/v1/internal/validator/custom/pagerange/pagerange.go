package pagerange

import (
	"fmt"
	"strconv"
)

func Validate(m map[string]string) error {
	for k, v := range m {
		switch k {
		case "paginationPage":
			page, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("paginationPage value must be integer: %w", err)
			}
			if page < MinPageValue || page > MaxPageValue {
				return fmt.Errorf("paginationPage value must be gte %d and lte %d", MinPageValue, MaxPageValue)
			}
		case "paginationCount":
			count, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("paginationPage value must be integer: %w", err)
			}
			if count < MinCountValue || count > MaxCountValue {
				return fmt.Errorf("paginationPage value must be gte %d and lte %d", MinPageValue, MaxPageValue)
			}
		}
	}
	return nil
}
