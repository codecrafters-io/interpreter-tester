package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEvaluateTerm(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	n1 := getRandInt()
	term1 := fmt.Sprintf("%d - %d", getRandInt(), getRandInt())
	term2 := fmt.Sprintf("%d + %d - %d", getRandInt(), getRandInt(), getRandInt())
	term3 := fmt.Sprintf("%d + %d - (-(%d - %d))", getRandInt(), getRandInt(), getRandInt(), getRandInt())
	term4 := fmt.Sprintf("(-%d + %d) * (%d * %d) / (1 + 4)", n1, n1, getRandInt(), getRandInt())

	evaluateTestCases := testcases.MultiEvaluateTestCase{
		FileContents: []string{term1, term2, term3, term4},
	}
	return evaluateTestCases.RunAll(b, logger)
}
