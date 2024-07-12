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
	error2 := fmt.Sprintf("(%s +)", getRandInt())
	error3 := "+"
	parseTestCase := testcases.MultiParseTestCase{
		FileContents: []string{error1, error2, error3},
	}
	return parseTestCase.RunAll(b, logger)
}
