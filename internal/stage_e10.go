package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEvaluateUnaryErrors(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	error1 := fmt.Sprintf("-\"%s\"", getRandString())
	error2 := "-true"
	error3 := "-false"
	error4 := fmt.Sprintf("-(\"%s\" + \"%s\")", getRandString(), getRandString())

	evaluateTestCases := testcases.MultiEvaluateTestCase{
		FileContents: []string{error1, error2, error3, error4},
	}
	return evaluateTestCases.RunAll(b, logger)
}
