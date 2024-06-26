package internal

import (
	"github.com/codecrafters-io/grep-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/grep-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testParen(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	shuffledString1 := random.RandShuffledStr(5, []rune("()"))
	shuffledString2 := random.RandShuffledStr(25, []rune("()"))
	shuffledString3 := random.RandShuffledStr(5, []rune("()")) + "##"
	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{"(", "))", shuffledString1, shuffledString2, shuffledString3},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
