package assertions

import (
	"fmt"

	"github.com/codecrafters-io/tester-utils/logger"
)

type StdoutAssertion struct {
	ExpectedLines []string
}

func NewStdoutAssertion(expectedLines []string) StdoutAssertion {
	return StdoutAssertion{ExpectedLines: expectedLines}
}

func (a StdoutAssertion) Run(stdout []string, logger *logger.Logger) error {
	var output []string

	for i, expectedLine := range a.ExpectedLines {
		if i >= len(stdout) {
			logAllSuccessLogs(output, logger)
			logger.Errorf("? %s", expectedLine)
			return fmt.Errorf("Expected line #%d on stdout to be %q, but didn't find line", i+1, expectedLine)
		}
		actualValue := stdout[i]

		if actualValue != expectedLine {
			logAllSuccessLogs(output, logger)
			logger.Errorf("𐄂 %s", actualValue)
			return fmt.Errorf("Expected line #%d on stdout to be %q, got %q", i+1, expectedLine, actualValue)
		} else {
			output = append(output, fmt.Sprintf("✓ %s", actualValue))
		}
	}

	if len(stdout) > len(a.ExpectedLines) {
		logAllSuccessLogs(output, logger)
		logger.Errorf("! %s", stdout[len(a.ExpectedLines)])
		return fmt.Errorf("Expected last line to be %q, but found %d more line(s)", stdout[len(a.ExpectedLines)-1], len(stdout)-len(a.ExpectedLines))
	}

	// If all lines match, we don't want to print all the lines again
	// We just want to print a single line summary
	logger.Successf("✓ %d line(s) match on stdout", len(a.ExpectedLines))

	return nil
}

func logAllSuccessLogs(output []string, logger *logger.Logger) {
	for _, line := range output {
		logger.Successf(line)
	}
}
