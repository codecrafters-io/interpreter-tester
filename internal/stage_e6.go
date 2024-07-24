package internal

import (
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEvaluateTerm(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	term1 := "10 - 20"
	term2 := "23 + 27 - 20"
	term3 := "2 + 23 - (-(25 - 50))"
	term4 := "-(-4 + 8) * (6 * 2) / (12 + 12)"

	evaluateTestCases := testcases.MultiEvaluateTestCase{
		FileContents: []string{term1, term2, term3, term4},
	}
	return evaluateTestCases.RunAll(b, logger)
}
