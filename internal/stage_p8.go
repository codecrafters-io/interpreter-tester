package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testParseComparison(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	randomInteger1 := getRandInt()
	randomInteger2 := getRandInt()
	// TODO: fix
	// comparisonExpr1 := fmt.Sprintf("%s > %s", randomInteger1, randomInteger1-randomInteger2)
	comparisonExpr2 := fmt.Sprintf("%s <= %s", randomInteger2, randomInteger1+randomInteger2)
	// comparisonExpr3 := fmt.Sprintf("%s < %s < %s", randomInteger1, randomInteger1+randomInteger2, randomInteger1+2*randomInteger2)
	comparisonExpr4 := fmt.Sprintf("(%s - %s) >= -(%s / %s + %s)", getRandInt(), getRandInt(), getRandInt(), getRandInt(), getRandInt())
	parseTestCase := testcases.MultiParseTestCase{
		// FileContents: []string{comparisonExpr1, comparisonExpr2, comparisonExpr3, comparisonExpr4},
		FileContents: []string{comparisonExpr2, comparisonExpr4},
	}
	return parseTestCase.RunAll(b, logger)
}
