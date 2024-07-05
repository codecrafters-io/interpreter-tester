package assertions

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/codecrafters-io/tester-utils/executable"
	"github.com/codecrafters-io/tester-utils/logger"
)

type StderrAssertion struct {
	ExpectedLines []string
}

func NewStderrAssertion(expectedLines []string) StderrAssertion {
	return StderrAssertion{ExpectedLines: expectedLines}
}

func (a StderrAssertion) Run(result executable.ExecutableResult, logger *logger.Logger) error {
	var successLogs []string
	stderr := getStderrLinesFromExecutableResult(result)
	skippedLines := getSkippedLinesCount(result)

	for i, expectedLine := range a.ExpectedLines {
		if i >= len(stderr) {
			logAllSuccessLogs(successLogs, logger)
			logger.Errorf("? %s", expectedLine)
			logger.Errorf("Skipped %d lines that didn't start with [line N]", skippedLines)
			return fmt.Errorf("Expected line #%d on stderr to be %q, but didn't find line", i+1, expectedLine)
		}
		actualValue := stderr[i]

		if actualValue != expectedLine {
			logAllSuccessLogs(successLogs, logger)
			logger.Errorf("𐄂 %s", actualValue)
			return fmt.Errorf("Expected line #%d on stderr to be %q, got %q", i+1, expectedLine, actualValue)
		} else {
			successLogs = append(successLogs, fmt.Sprintf("✓ %s", actualValue))
		}
	}

	if len(stderr) > len(a.ExpectedLines) {
		logAllSuccessLogs(successLogs, logger)
		logger.Errorf("! %s", stderr[len(a.ExpectedLines)])
		return fmt.Errorf("Expected last stderr line to be %q, but found extra line: %q", a.ExpectedLines[len(a.ExpectedLines)-1], stderr[len(a.ExpectedLines)])
	}

	// If all lines match, we don't want to print all the lines again
	// We just want to print a single line summary
	logger.Successf("✓ %d line(s) match on stderr", len(a.ExpectedLines))
	return nil
}

func getStderrLinesFromExecutableResult(result executable.ExecutableResult) []string {
	var filteredStdErr []string
	stderr := strings.Split(strings.TrimRight(string(result.Stderr), "\n"), "\n")
	regex := regexp.MustCompile(`\[line [0-9]+\]`)

	for _, line := range stderr {
		if regex.MatchString(line) {
			filteredStdErr = append(filteredStdErr, line)
		}
	}

	return filteredStdErr
}

func getSkippedLinesCount(result executable.ExecutableResult) int {
	unfilteredStderr := strings.Split(strings.TrimRight(string(result.Stderr), "\n"), "\n")
	return len(unfilteredStderr) - len(getStderrLinesFromExecutableResult(result))
}
