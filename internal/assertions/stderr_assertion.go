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
	if len(stderr) != len(a.ExpectedLines) {
		return []string{}, fmt.Errorf("Expected %d lines of stderr, got %d", len(a.ExpectedLines), len(stderr))
	}

	var output []string
	for i, expectedLines := range a.ExpectedLines {
		actualValue := stderr[i]

		if actualValue != expectedLines {
			output = append(output, fmt.Sprintf("𐄂 %s", actualValue))
			return output, fmt.Errorf("Expected line #%d on stderr to be %q, got %q", i+1, expectedLines, actualValue)
		} else {
			output = append(output, fmt.Sprintf("✓ %s", actualValue))
		}
	}

	successOutput := []string{fmt.Sprintf("✓ %d line(s) match on stderr", len(a.ExpectedLines))}
	return successOutput, nil
}
