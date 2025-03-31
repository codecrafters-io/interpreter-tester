package internal

import (
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEOF(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	tokenizeTestCase := testcases.TokenizeTestCase{
		FileContents: "", ExpectsError: false,
	}

	logger.UpdateSecondaryPrefix(("test-1"))
	defer logger.ResetSecondaryPrefix()
	return tokenizeTestCase.Run(b, logger)
}
