package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testParseTerms(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	termExpr1 := "\"hello\" + \"world\""
	termExpr2 := fmt.Sprintf("%d - %d * %d - %d", getRandInt(), getRandInt(), getRandInt(), getRandInt())
	termExpr3 := fmt.Sprintf("%d + %d - %d / %d", getRandInt(), getRandInt(), getRandInt(), getRandInt())
	termExpr4 := fmt.Sprintf("(-%d + %d) * (%d * %d) / (%d + %d)", getRandInt(), getRandInt(), getRandInt(), getRandInt(), getRandInt(), getRandInt())
	parseTestCase := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.ParseTestCase{FileContents: termExpr1, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: termExpr2, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: termExpr3, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: termExpr4, ExpectsError: false},
		},
	}
	return parseTestCase.RunAll(b, logger)
}
