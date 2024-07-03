package internal

import (
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var Strings = []string{"\"hello\"", "\"world\"", "\"foo\"", "\"bar\"", "\"baz\""}

func testStrings(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	string1 := random.RandomSelection(2, Strings, " ")
	string2 := `"foo 	bar 123 // hello world!"`
	shuffledString2 := (random.RandomSelection(2, Strings, "+")) + `"perseverance" && "Success" != "Failure"`

	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{Strings[0], string1, string2, shuffledString2},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
