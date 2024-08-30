package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEvaluateConcat(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	concat1 := fmt.Sprintf("\"%s\" + \"%s\"", getRandString(), getRandString())
	concat2 := fmt.Sprintf("\"%s\" + \"%s\"", getRandString(), getRandIntAsString())
	concat3 := fmt.Sprintf("\"%s\" + \"%s\" + \"%s\"", getRandString(), getRandString(), getRandString())
	concat4 := fmt.Sprintf("(\"%s\" + \"%s\") + (\"%s\" + \"%s\")", getRandString(), getRandString(), getRandString(), getRandString())

	evaluateTestCases := testcases.MultiEvaluateTestCase{
		TestCases: []testcases.EvaluateTestCase{
			{FileContents: concat1, ExpectsError: false},
			{FileContents: concat2, ExpectsError: false},
			{FileContents: concat3, ExpectsError: false},
			{FileContents: concat4, ExpectsError: false},
		},
	}
	return evaluateTestCases.RunAll(b, logger)
}
