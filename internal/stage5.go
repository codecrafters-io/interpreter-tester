package internal

import (
	"slices"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var LexicalErrors = []string{"@", "$", "#", "%"}

func testErrors(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	shuffledString1 := randomStringFromCharacters(5, LexicalErrors)
	shuffledString2 := "{(" + randomStringFromCharacters(5, slices.Concat(SingleCharOperators, LexicalErrors)) + ")}"
	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{"@", ",.$(#", shuffledString1, shuffledString2},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
