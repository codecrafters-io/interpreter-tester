package internal

import (
	"github.com/codecrafters-io/grep-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/grep-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testSingleChars(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	// shuffledString1 := random.RandomStringFromCharacters(15, []rune("+-*.,;"))
	// shuffledString2 := random.RandomStringFromCharacters(20, []rune("+-*.,;(){}"))
	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{"+-", "+-*.,;"},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
