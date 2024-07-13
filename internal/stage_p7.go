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
	termExpr4 := fmt.Sprintf("-(-%d + %d) * (%d * %d) / (%d + %d)", getRandInt(), getRandInt(), getRandInt(), getRandInt(), getRandInt(), getRandInt())
	parseTestCase := testcases.MultiParseTestCase{
		FileContents: []string{termExpr1, termExpr2, termExpr3, termExpr4},
	}
	return parseTestCase.RunAll(b, logger)
}
