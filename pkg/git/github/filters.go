package github

import (
	"fmt"
)

//buildFilter is a helper function to extend a list of filters.
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
