package assertions

import (
	"fmt"
	"strconv"

	"github.com/codecrafters-io/tester-utils/executable"
	"github.com/codecrafters-io/tester-utils/logger"
)

type NumberWithinRangeAssertion struct {
	minValue int
	maxValue int
}

func NewNumberWithinRangeAssertion(minValue int, maxValue int) NumberWithinRangeAssertion {
	return NumberWithinRangeAssertion{minValue: minValue, maxValue: maxValue}
}

func (a NumberWithinRangeAssertion) Run(result executable.ExecutableResult, logger *logger.Logger) error {
	// We expect a single line of output, which would be cast to an integer and compared to the min and max values
	stdOutput := getStdoutLinesFromExecutableResult(result)

	if len(stdOutput) != 1 {
		return fmt.Errorf("Expected a single line of output (%d), got %d lines", a.minValue, len(stdOutput))
	}

	actualValue, err := strconv.ParseFloat(stdOutput[0], 64)
	if err != nil {
		return fmt.Errorf("Expected a single line of output to be an integer (%d), got %q", a.minValue, stdOutput[0])
	}
	actualValueInt := int(actualValue)

	if actualValueInt < a.minValue || actualValueInt > a.maxValue {
		logger.Errorf("𐄂 %d", actualValueInt)
		return fmt.Errorf("Expected integer to be between %d and %d, got %d", a.minValue, a.maxValue, actualValueInt)
	} else {
		logger.Successf("✓ %f", actualValue)
	}

	// Keep parity with all other stage outputs
	logger.Successf("✓ %d line(s) match on stdout", 1)

	return nil
}
