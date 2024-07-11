package internal

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/tester-utils/random"
)

func repeatSlice(slice []string, n int) []string {
	ret := make([]string, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, slice...)
	}
	return ret
}

func getRandInt() string {
	return fmt.Sprint(random.RandomInt(10, 100))
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
