package assertions

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/tester-utils/executable"
	"github.com/codecrafters-io/tester-utils/logger"
)

type StdoutAssertion struct {
	ExpectedLines []string
}

func NewStdoutAssertion(expectedLines []string) StdoutAssertion {
	return StdoutAssertion{ExpectedLines: expectedLines}
}

func (a StdoutAssertion) Run(result executable.ExecutableResult, logger *logger.Logger) error {
	var successLogs []string
	stdout := getStdoutLinesFromExecutableResult(result)

	for i, expectedLine := range a.ExpectedLines {
		if i >= len(stdout) {
			logAllSuccessLogs(successLogs, logger)
			logger.Errorf("? %s", expectedLine)
			return fmt.Errorf("Expected line #%d on stdout to be %q, but didn't find line", i+1, expectedLine)
		}
		actualValue := stdout[i]

		if actualValue != expectedLine {
			logAllSuccessLogs(successLogs, logger)
			logger.Errorf("𐄂 %s", actualValue)
			return fmt.Errorf("Expected line #%d on stdout to be %q, got %q", i+1, expectedLine, actualValue)
		} else {
			successLogs = append(successLogs, fmt.Sprintf("✓ %s", actualValue))
		}
	}

	if len(stdout) > len(a.ExpectedLines) {
		logAllSuccessLogs(successLogs, logger)
		logger.Errorf("! %s", stdout[len(a.ExpectedLines)])
		return fmt.Errorf("Expected nothing after last stdout line %q, but found extra line: %q", a.ExpectedLines[len(a.ExpectedLines)-1], stdout[len(a.ExpectedLines)])
	}

	// If all lines match, we don't want to print all the lines again
	// We just want to print a single line summary
	logger.Successf("✓ %d line(s) match on stdout", len(a.ExpectedLines))

	return nil
}

func logAllSuccessLogs(successLogs []string, logger *logger.Logger) {
	for _, line := range successLogs {
		logger.Successf(line)
	}
}

func getStdoutLinesFromExecutableResult(result executable.ExecutableResult) []string {
	stdout := strings.Split(strings.TrimRight(string(result.Stdout), "\n"), "\n")
	return stdout
}
