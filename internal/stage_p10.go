package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testParseErrors(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	error1 := fmt.Sprintf("\"%s", random.RandomElementFromArray(STRINGS))
	error2 := fmt.Sprintf("(%d +)", getRandInt())
	error3 := "+"
	parseTestCase := testcases.MultiParseTestCase{
		TestCases: []testcases.ParseTestCase{
			{FileContents: error1, ExpectsError: true},
			{FileContents: error2, ExpectsError: true},
			{FileContents: error3, ExpectsError: true},
		},
	}
	return parseTestCase.RunAll(b, logger)
}
