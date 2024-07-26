package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEvaluateCompErrors(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	error1 := fmt.Sprintf("\"%s\" < %s", getRandString(), getRandBoolean())
	error2 := fmt.Sprintf("%s <= (%d + %d)", getRandBoolean(), getRandInt(), getRandInt())
	error3 := fmt.Sprintf("%d > (\"%s\" + \"%s\")", getRandInt(), getRandString(), getRandString())
	error4 := fmt.Sprintf("%s >= %s", getRandBoolean(), getRandBoolean())

	evaluateTestCases := testcases.MultiEvaluateTestCase{
		FileContents: []string{error1, error2, error3, error4},
	}
	return evaluateTestCases.RunAll(b, logger)
}
