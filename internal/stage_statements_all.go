package internal

import (
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

// For each stage, we need to create a test function that will be used to test the user submissions for that stage.
// This function will be called by the tester_definition.go file, to create the test function for that stage.
func createTestForRunCommandStage(stageId string) func(*test_case_harness.TestCaseHarness) error {
	return func(stageHarness *test_case_harness.TestCaseHarness) error {
		b := interpreter_executable.NewInterpreterExecutable(stageHarness)

		logger := stageHarness.Logger

		runTestCases := testcases.MultiTestCase{
			TestCases: GetTestCasesForCurrentStageWithRandomValues(stageId),
		}

		return runTestCases.RunAll(b, logger)
	}
}
