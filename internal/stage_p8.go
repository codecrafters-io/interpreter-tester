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
	comparisonExpr1 := fmt.Sprintf("%d > %d", randomInteger1, randomInteger1-randomInteger2)
	comparisonExpr2 := fmt.Sprintf("%d <= %d", randomInteger2, randomInteger1+randomInteger2)
	comparisonExpr3 := fmt.Sprintf("%d < %d < %d", randomInteger1, randomInteger1+randomInteger2, randomInteger1+2*randomInteger2)
	comparisonExpr4 := fmt.Sprintf("(%d - %d) >= -(%d / %d + %d)", getRandInt(), getRandInt(), getRandInt(), getRandInt(), getRandInt())
	parseTestCase := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.ParseTestCase{FileContents: comparisonExpr1, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: comparisonExpr2, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: comparisonExpr3, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: comparisonExpr4, ExpectsError: false},
		},
	}
	return parseTestCase.RunAll(b, logger)
}
