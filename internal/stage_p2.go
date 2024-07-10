package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testParseNumbers(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	numberLiteral1 := fmt.Sprint(random.RandomInt(10, 100))
	numberLiteral2 := fmt.Sprint(random.RandomInt(10, 100)) + "." + fmt.Sprint(random.RandomInt(10, 100))
	parseTestCase := testcases.MultiParseTestCase{
		FileContents: []string{numberLiteral1, "0.0", numberLiteral2},
	}
	return parseTestCase.RunAll(b, logger)
}
