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

	if skippedLines > 0 {
		logger.Plainf("[stderr] Skipped %d lines that didn't start with [line N]", skippedLines)
	}

	for i, expectedLine := range a.ExpectedLines {
		if i >= len(stderr) {
			logAllSuccessLogs(successLogs, logger)

			return fmt.Errorf(`
[stderr] Missing line #%d from stderr: %q
[stderr] Perhaps it's printed to stdout? It should be printed to stderr.
 `, i+1, expectedLine)
		}
		actualValue := stderr[i]

		if actualValue != expectedLine {
			logAllSuccessLogs(successLogs, logger)

			return fmt.Errorf(`
[stderr] Mismatch on line #%d of stderr:
[stderr] Expected: %q
[stderr] Actual  : %q
 `, i+1, expectedLine, actualValue)
		} else {
			successLogs = append(successLogs, fmt.Sprintf("✓ %s", actualValue))
		}
	}

	if len(stderr) > len(a.ExpectedLines) {
		logAllSuccessLogs(successLogs, logger)
		return fmt.Errorf(`
𐄂 [stderr] Extra unexpected line in stderr: %q
 `, stderr[len(a.ExpectedLines)])
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
	trimmedStderr := strings.TrimRight(string(result.Stderr), "\n")

	// Handle the case where strings.Split("", "\n") returns [""]
	if trimmedStderr == "" {
		return 0
	}

	unfilteredStderrLines := strings.Split(trimmedStderr, "\n")
	return len(unfilteredStderrLines) - len(getStderrLinesFromExecutableResult(result))
}
