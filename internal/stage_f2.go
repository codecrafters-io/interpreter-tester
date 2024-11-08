package internal

import (
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

// THis isn't even need - juust use createTestFor...
func testFunctionsNoArgs(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	runTestCases := testcases.MultiTestCase{TestCases: GetTestCasesForCurrentStageWithRandomValues("f2")}
	return runTestCases.RunAll(b, logger)
}
