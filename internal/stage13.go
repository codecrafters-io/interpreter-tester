package internal

import (
	"slices"

	"github.com/codecrafters-io/grep-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/grep-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testNumbers(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	shuffledString2 := random.RandomStringFromCharacters(35, slices.Concat(Parens, Braces, SingleCharOperators, LexicalErrors, Equals, Negation, Relational, Division, Whitespace, Strings, []string{"123.456", "456.789"}))
	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{"1234", "1234.1234", "1234.1234.1234.", shuffledString2},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
