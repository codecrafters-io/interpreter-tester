package internal

import (
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var Strings = []string{"\"hello\"", "\"world\"", "\"foo\"", "\"bar\"", "\"baz\""}

func testStrings(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	string1 := randomSelection(2, Strings, " ")
	string2 := `"foo 	bar 123 // hello world!"`
	shuffledString2 := (randomSelection(2, Strings, "+")) + `"perseverance" && "Success" != "Failure"`

	tokenizeTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{Strings[0], string1, string2, shuffledString2},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
