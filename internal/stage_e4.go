package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEvaluateUnary(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	unary1 := fmt.Sprintf("-%d", getRandInt())
	unary2 := fmt.Sprintf("!%s", getRandBoolean())
	unary3 := "!nil"
	unary4 := fmt.Sprintf("(!!%d)", getRandInt())

	evaluateTestCases := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.EvaluateTestCase{FileContents: unary1, ExpectsError: false},
			&testcases.EvaluateTestCase{FileContents: unary2, ExpectsError: false},
			&testcases.EvaluateTestCase{FileContents: unary3, ExpectsError: false},
			&testcases.EvaluateTestCase{FileContents: unary4, ExpectsError: false},
		},
	}
	return evaluateTestCases.RunAll(b, logger)
}
