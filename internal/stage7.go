package internal

import (
	"slices"

	"github.com/codecrafters-io/grep-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/grep-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var Negation = []string{"!", "!="}

func testNegation(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	shuffledString1 := "{(" + random.RandomStringFromCharacters(5, slices.Concat(LexicalErrors, Equals, Negation)) + ")}"
	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{"!=", "!!===", "!{!}(!===)=", shuffledString1},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
