package assertions

import (
	"fmt"

	"github.com/codecrafters-io/tester-utils/logger"
)

type StderrAssertion struct {
	ExpectedLines []string
}

func NewStderrAssertion(expectedLines []string) StderrAssertion {
	return StderrAssertion{ExpectedLines: expectedLines}
}

func (a StderrAssertion) Run(stderr []string, logger *logger.Logger) error {
	var output []string

	for i, expectedLine := range a.ExpectedLines {
		if i >= len(stderr) {
			for _, line := range output {
				logger.Successf(line)
			}
			logger.Errorf("? %s", expectedLine)
			return fmt.Errorf("Expected line #%d on stderr to be %q, but didn't find line", i+1, expectedLine)
		}
		actualValue := stderr[i]

		if actualValue != expectedLine {
			for _, line := range output {
				logger.Successf(line)
			}
			logger.Errorf("𐄂 %s", actualValue)
			return fmt.Errorf("Expected line #%d on stderr to be %q, got %q", i+1, expectedLine, actualValue)
		} else {
			output = append(output, fmt.Sprintf("✓ %s", actualValue))
		}
	}

	if len(stderr) > len(a.ExpectedLines) {
		for _, line := range output {
			logger.Successf(line)
		}
		logger.Errorf("! %s", stderr[len(a.ExpectedLines)])
		return fmt.Errorf("Expected last line to be %q, but found %d more line(s)", stderr[len(a.ExpectedLines)-1], len(stderr)-len(a.ExpectedLines))
	}

	// If all lines match, we don't want to print all the lines again
	// We just want to print a single line summary
	logger.Successf("✓ %d line(s) match on stderr", len(a.ExpectedLines))
	return nil
}
