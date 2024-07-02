package internal

import (
	"slices"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testNumbers(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	shuffledString2 := "{ 1234.00" + random.RandomStringFromCharacters(5, slices.Concat(SingleCharOperators, LexicalErrors, Whitespace)) + "00.1234 }"
	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{"1234", "1234.1234", "1234.1234.1234.", shuffledString2},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
