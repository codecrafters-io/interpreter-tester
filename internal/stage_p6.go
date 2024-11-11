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

	parseTestCase := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.ParseTestCase{FileContents: factorExpr1, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: factorExpr2, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: factorExpr3, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: factorExpr4, ExpectsError: false},
		},
	}
	return parseTestCase.RunAll(b, logger)
}
