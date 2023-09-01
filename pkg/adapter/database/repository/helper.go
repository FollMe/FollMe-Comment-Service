package repository

import (
	"fmt"
	"strings"
)

func buildParamsStruct(quantity int) string {
	params := []string{}
	if quantity <= 0 {
		return ""
	}

	for i := 1; i <= quantity; i++ {
		params = append(params, fmt.Sprintf("$%v", i))
	}

	return strings.Join(params, ", ")
}
