package internal

import (
	"slices"

	"github.com/codecrafters-io/grep-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/grep-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var Strings = []string{"\"hello\"", "\"world\"", "\"foo\"", "\"bar\"", "\"baz\""}

func testStrings(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	string1 := `"foo baz"`
	string2 := `"foo 	bar 123 // hello world!"`
	shuffledString2 := "(({" + random.RandomStringFromCharacters(5, slices.Concat(Whitespace, Strings)) + "}))"

	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{Strings[0], string1, string2, shuffledString2},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
