package assertions

import (
	"fmt"
)

type StdoutAssertion struct {
	ExpectedLines []string
}

func NewStdoutAssertion(expectedLines []string) StdoutAssertion {
	return StdoutAssertion{ExpectedLines: expectedLines}
}

func (a StdoutAssertion) Run(stdout []string) ([]string, error) {
	var output []string

	for i, expectedLine := range a.ExpectedLines {
		if i >= len(stdout) {
			output = append(output, fmt.Sprintf("? %s", expectedLine))
			return output, fmt.Errorf("Expected line #%d on stdout to be %q, but didn't find line", i+1, expectedLine)
		}
		actualValue := stdout[i]

		if actualValue != expectedLine {
			output = append(output, fmt.Sprintf("𐄂 %s", actualValue))
			return output, fmt.Errorf("Expected line #%d on stdout to be %q, got %q", i+1, expectedLine, actualValue)
		} else {
			output = append(output, fmt.Sprintf("✓ %s", actualValue))
		}
	}

	if len(stdout) > len(a.ExpectedLines) {
		output = append(output, fmt.Sprintf("! %s", stdout[len(a.ExpectedLines)]))
		return output, fmt.Errorf("Expected last line to be %q, but found %d more line(s)", stdout[len(a.ExpectedLines)-1], len(stdout)-len(a.ExpectedLines))
	}

	// If all lines match, we don't want to print all the lines again
	// We just want to print a single line summary
	successOutput := []string{fmt.Sprintf("✓ %d line(s) match on stdout", len(a.ExpectedLines))}
	return successOutput, nil
}
