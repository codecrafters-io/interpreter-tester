package internal

import (
	"fmt"

	"github.com/codecrafters-io/tester-utils/random"
)

func getRandInt() int {
	return random.RandomInt(10, 100)
}

func getRandIntAsString() string {
	return fmt.Sprintf("%d", getRandInt())
}

func getRandString() string {
	return random.RandomElementFromArray(STRINGS)
}
