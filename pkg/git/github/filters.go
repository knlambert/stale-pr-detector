package github

import (
	"fmt"
)

func buildFilter(
	parameter string,
	values *[]string,
	queryFilters *[]string,
) {
	if values != nil {
		var newFilters []string
		for _, value := range *values {
			newFilters = append(newFilters, fmt.Sprintf("%s:%s", parameter, value))
		}
		*queryFilters = append(*queryFilters, newFilters...)
	}
}
