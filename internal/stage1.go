package internal

import (
	"github.com/codecrafters-io/grep-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/grep-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEOF(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	commandTestCases := testcases.TokenizeOutputTestCase{
		FileContents: "",
	}
	if err := commandTestCases.Run(b, logger); err != nil {
		return err
	}

	return nil
}
