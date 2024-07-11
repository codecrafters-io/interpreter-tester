package internal

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testStrings(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	string1 := joinWith(random.RandomElementsFromArray(QuotedStrings, 1), "") + " " + "\"unterminated"
	string2 := `"foo 	bar 123 // hello world!"`
	shuffledString2 := joinWith(random.RandomElementsFromArray(QuotedStrings, 2), "+") + `"perseverance" && "Success" != "Failure"`

	tokenizeTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{QuotedStrings[0], string1, string2, shuffledString2},
	}
	return tokenizeTestCases.RunAll(b, logger)
}

func joinWith[T any](arr []T, sep string) string {
	return strings.Join(convertToStringSlice(arr), sep)
}

func convertToStringSlice[T any](arr []T) []string {
	result := make([]string, len(arr))
	for i, v := range arr {
		result[i] = fmt.Sprintf("%v", v)
	}
	return result
}
