package assertions

import (
	"fmt"
	"strconv"

	"github.com/codecrafters-io/tester-utils/executable"
	"github.com/codecrafters-io/tester-utils/logger"
)

type IntegerAssertion struct {
	minValue int
	maxValue int
}

func NewIntegerAssertion(minValue int, maxValue int) IntegerAssertion {
	return IntegerAssertion{minValue: minValue, maxValue: maxValue}
}

func (a IntegerAssertion) Run(result executable.ExecutableResult, logger *logger.Logger) error {
	// We expect a single line of output, which would be cast to an integer and compared to the min and max values
	stdOutput := getStdoutLinesFromExecutableResult(result)

	if len(stdOutput) != 1 {
		return fmt.Errorf("Expected a single line of output (%d), got %d lines", a.minValue, len(stdOutput))
	}

	actualValue, err := strconv.Atoi(stdOutput[0])
	if err != nil {
		return fmt.Errorf("Expected a single line of output to be an integer (%d), got %q", a.minValue, stdOutput[0])
	}

	if actualValue < a.minValue || actualValue > a.maxValue {
		logger.Errorf("𐄂 %d", actualValue)
		return fmt.Errorf("Expected integer to be between %d and %d, got %d", a.minValue, a.maxValue, actualValue)
	} else {
		logger.Successf("✓ %d", actualValue)
	}

	// Keep parity with all other stage outputs
	logger.Successf("✓ %d line(s) match on stdout", 1)

	return nil
}
