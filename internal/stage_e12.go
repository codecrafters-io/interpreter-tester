package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEvaluateAddErrors(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	error1 := fmt.Sprintf("\"%s\" + %s", getRandString(), getRandBoolean())
	error2 := fmt.Sprintf("%d + \"%s\" + %d", getRandInt(), getRandString(), getRandInt())
	error3 := fmt.Sprintf("%d - %s", getRandInt(), getRandBoolean())
	error4 := fmt.Sprintf("%s - (\"%s\" + \"%s\")", getRandBoolean(), getRandString(), getRandString())

	evaluateTestCases := testcases.MultiEvaluateTestCase{
		FileContents: []string{error1, error2, error3, error4},
	}
	return evaluateTestCases.RunAll(b, logger)
}
