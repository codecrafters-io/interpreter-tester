package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEvaluateRelational(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	n1, n2, n3 := getRandInt(), getRandInt(), getRandInt()
	relational1 := fmt.Sprintf("%d > -%d", n1, n1+n2)
	relational2 := fmt.Sprintf("%d <= %d", n1, n1+n2+n3)
	relational3 := fmt.Sprintf("%d >= %d", n3, n3)
	relational4 := fmt.Sprintf("(%d - %d) >= -(%d / %d + %d)", getRandInt(), getRandInt(), n1*2, n1, getRandInt())

	evaluateTestCases := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.EvaluateTestCase{FileContents: relational1, ExpectsError: false},
			&testcases.EvaluateTestCase{FileContents: relational2, ExpectsError: false},
			&testcases.EvaluateTestCase{FileContents: relational3, ExpectsError: false},
			&testcases.EvaluateTestCase{FileContents: relational4, ExpectsError: false},
		},
	}
	return evaluateTestCases.RunAll(b, logger)
}
