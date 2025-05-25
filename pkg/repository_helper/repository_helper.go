package repository_helper

import (
	"fmt"
	"strings"
)

func BuildParamsStruct(quantity int) string {
	params := []string{}
	if quantity <= 0 {
		return ""
	}

	for i := 1; i <= quantity; i++ {
		params = append(params, fmt.Sprintf("$%v", i))
	}

	return strings.Join(params, ", ")
}
