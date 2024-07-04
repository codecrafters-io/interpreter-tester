package assertions

import (
	"fmt"
)

type StderrAssertion struct {
	ExpectedLines []string
}

func NewStderrAssertion(expectedLines []string) StderrAssertion {
	return StderrAssertion{ExpectedLines: expectedLines}
}

func (a StderrAssertion) Run(stderr []string) ([]string, error) {
	var output []string

	for i, expectedLine := range a.ExpectedLines {
		if i >= len(stderr) {
			output = append(output, fmt.Sprintf("? %s", expectedLine))
			return output, fmt.Errorf("Expected line #%d on stderr to be %q, but didn't find line", i+1, expectedLine)
		}
		actualValue := stderr[i]

		if actualValue != expectedLine {
			output = append(output, fmt.Sprintf("𐄂 %s", actualValue))
			return output, fmt.Errorf("Expected line #%d on stderr to be %q, got %q", i+1, expectedLine, actualValue)
		} else {
			output = append(output, fmt.Sprintf("✓ %s", actualValue))
		}
	}

	if len(stderr) > len(a.ExpectedLines) {
		output = append(output, fmt.Sprintf("! %s", stderr[len(a.ExpectedLines)]))
		return output, fmt.Errorf("Expected last line to be %q, but found %d more line(s)", stderr[len(a.ExpectedLines)-1], len(stderr)-len(a.ExpectedLines))
	}

	// If all lines match, we don't want to print all the lines again
	// We just want to print a single line summary
	successOutput := []string{fmt.Sprintf("✓ %d line(s) match on stderr", len(a.ExpectedLines))}
	return successOutput, nil
}
