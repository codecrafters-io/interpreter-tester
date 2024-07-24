package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEvaluateFactor(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	n1 := random.RandomInt(2, 5)
	factor1 := fmt.Sprintf("%d * %d", getRandInt(), getRandInt())
	factor2 := fmt.Sprintf("%d / 5", getRandInt())
	factor3 := fmt.Sprintf("7 * %d / 7 / 1", n1)
	factor4 := fmt.Sprintf("(18 * %d / (3 * 6))", n1)

	evaluateTestCases := testcases.MultiEvaluateTestCase{
		FileContents: []string{factor1, factor2, factor3, factor4},
	}
	return evaluateTestCases.RunAll(b, logger)
}
