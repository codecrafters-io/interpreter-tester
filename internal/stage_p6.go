package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testParseFactor(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	factorExpr1 := fmt.Sprintf("%d * %d / %d", getRandInt(), getRandInt(), getRandInt())
	factorExpr2 := fmt.Sprintf("%d / %d / %d", getRandInt(), getRandInt(), getRandInt())
	factorExpr3 := fmt.Sprintf("%d * %d * %d / %d", getRandInt(), getRandInt(), getRandInt(), getRandInt())
	factorExpr4 := fmt.Sprintf("(%d * -%d / (%d * %d))", getRandInt(), getRandInt(), getRandInt(), getRandInt())

	parseTestCase := testcases.MultiParseTestCase{
		FileContents: []string{factorExpr1, factorExpr2, factorExpr3, factorExpr4},
	}
	return parseTestCase.RunAll(b, logger)
}
