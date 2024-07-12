package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testParseUnary(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	unaryExpr1 := "!" + random.RandomElementFromArray(BOOLEANS)
	unaryExpr2 := "-" + fmt.Sprint(random.RandomInt(10, 100))
	unaryExpr3 := "!!" + random.RandomElementFromArray(BOOLEANS)
	unaryExpr4 := "(!!(" + random.RandomElementFromArray(BOOLEANS) + "))"
	parseTestCase := testcases.MultiParseTestCase{
		FileContents: []string{unaryExpr1, unaryExpr2, unaryExpr3, unaryExpr4},
	}
	return parseTestCase.RunAll(b, logger)
}
