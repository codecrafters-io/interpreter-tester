package internal

import (
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEvaluateBooleans(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	evaluateTestCases := testcases.MultiEvaluateTestCase{
		TestCases: []testcases.EvaluateTestCase{
			{FileContents: "true", ExpectsError: false},
			{FileContents: "false", ExpectsError: false},
			{FileContents: "nil", ExpectsError: false},
		},
	}
	return evaluateTestCases.RunAll(b, logger)
}
