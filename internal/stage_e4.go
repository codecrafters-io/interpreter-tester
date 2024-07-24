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

	evaluateTestCases := testcases.MultiEvaluateTestCase{
		FileContents: []string{unary1, unary2, unary3, unary4},
	}
	return evaluateTestCases.RunAll(b, logger)
}
