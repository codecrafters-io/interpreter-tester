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
	if len(stdout) != len(a.ExpectedLines) {
		return []string{}, fmt.Errorf("Expected %d lines of stdout, got %d", len(a.ExpectedLines), len(stdout))
	}

	var output []string
	for i, expectedLines := range a.ExpectedLines {
		actualValue := stdout[i]

		if actualValue != expectedLines {
			output = append(output, fmt.Sprintf("𐄂 %s", actualValue))
			return output, fmt.Errorf("Expected line #%d on stdout to be %q, got %q", i+1, expectedLines, actualValue)
		} else {
			output = append(output, fmt.Sprintf("✓ %s", actualValue))
		}
	}

	successOutput := []string{fmt.Sprintf("✓ %d line(s) match on stdout", len(a.ExpectedLines))}
	return successOutput, nil
}
