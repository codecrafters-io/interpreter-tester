package internal

import (
	"slices"

	"github.com/codecrafters-io/grep-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/grep-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var Strings = []string{"hello", "world", "foo", "bar", "baz"}

func testStrings(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	shuffledString1 := random.RandomStringFromCharacters(20, slices.Concat(Whitespace, Strings))
	shuffledString2 := random.RandomStringFromCharacters(35, slices.Concat(Parens, Braces, SingleCharOperators, LexicalErrors, Equals, Negation, Relational, Division, Whitespace, Strings))

	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{"", ` foo	baz `, shuffledString1, shuffledString2},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
