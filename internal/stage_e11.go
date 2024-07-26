package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEvaluateMultErrors(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	error1 := fmt.Sprintf("%d * \"%s\"", getRandInt(), getRandString())
	error2 := fmt.Sprintf("\"%s\" / %d", getRandString(), getRandInt())
	error3 := fmt.Sprintf("%s / %s", getRandBoolean(), getRandBoolean())
	error4 := fmt.Sprintf("(\"%s\" + \"%s\") * (\"%s\" + \"%s\")", getRandString(), getRandString(), getRandString(), getRandString())

	evaluateTestCases := testcases.MultiEvaluateTestCase{
		FileContents: []string{error1, error2, error3, error4},
	}
	return evaluateTestCases.RunAll(b, logger)
}
