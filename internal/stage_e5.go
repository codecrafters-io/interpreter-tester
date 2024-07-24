package internal

import (
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEvaluateFactor(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	factor1 := "12 * 2"
	factor2 := "66 / 11"
	factor3 := "14 * 4 / 7 / 2"
	factor4 := "(18 * -2 / (3 * 6))"

	evaluateTestCases := testcases.MultiEvaluateTestCase{
		FileContents: []string{factor1, factor2, factor3, factor4},
	}
	return evaluateTestCases.RunAll(b, logger)
}
