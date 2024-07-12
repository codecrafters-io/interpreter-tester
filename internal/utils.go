package internal

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/tester-utils/random"
)

func getRandInt() int {
	return random.RandomInt(10, 100)
}

func joinWith[T any](arr []T, sep string) string {
	return strings.Join(convertToStringSlice(arr), sep)
}

func convertToStringSlice[T any](arr []T) []string {
	result := make([]string, len(arr))
	for i, v := range arr {
		result[i] = fmt.Sprintf("%v", v)
	}
	return result
}
