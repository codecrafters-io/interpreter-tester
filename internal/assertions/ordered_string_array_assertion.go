package assertions

import (
	"fmt"
)

// OrderedStringArrayAssertion : Order of the actual and expected values matters.
// We don't alter the ordering.
type OrderedStringArrayAssertion struct {
	ExpectedValue []string
	Type          string
}

func NewOrderedStringArrayAssertion(expectedValue []string, t string) OrderedStringArrayAssertion {
	return OrderedStringArrayAssertion{ExpectedValue: expectedValue, Type: t}
}

func (a OrderedStringArrayAssertion) Run(value []string) ([]string, error) {
	if len(value) != len(a.ExpectedValue) {
		return []string{}, fmt.Errorf("Expected %d lines of %s, got %d", len(a.ExpectedValue), a.Type, len(value))
	}

	var output []string
	for i, expectedValue := range a.ExpectedValue {
		actualValue := value[i]

		if actualValue != expectedValue {
			output = append(output, fmt.Sprintf("𐄂 %s", actualValue))
			return output, fmt.Errorf("Expected line #%d on %s to be %q, got %q", i+1, a.Type, expectedValue, actualValue)
		} else {
			output = append(output, fmt.Sprintf("✓ %s", actualValue))
		}
	}

	successOutput := []string{fmt.Sprintf("✓ %d line(s) match on %s", len(a.ExpectedValue), a.Type)}
	return successOutput, nil
}
