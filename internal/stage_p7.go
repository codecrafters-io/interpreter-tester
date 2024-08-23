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
	termExpr2 := fmt.Sprintf("%d - %d - %d", getRandInt(), getRandInt(), getRandInt())
	termExpr3 := fmt.Sprintf("%d + %d - %d", getRandInt(), getRandInt(), getRandInt())
	termExpr4 := fmt.Sprintf("(-%d + %d) * (%d * %d) / (%d + %d)", getRandInt(), getRandInt(), getRandInt(), getRandInt(), getRandInt(), getRandInt())
	parseTestCase := testcases.MultiParseTestCase{
		TestCases: []testcases.ParseTestCase{
			{FileContents: termExpr1, ExpectsError: false},
			{FileContents: termExpr2, ExpectsError: false},
			{FileContents: termExpr3, ExpectsError: false},
			{FileContents: termExpr4, ExpectsError: false},
		},
	}
	return parseTestCase.RunAll(b, logger)
}
