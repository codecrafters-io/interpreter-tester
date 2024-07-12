package internal

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testParseEquality(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	randomWord := random.RandomElementFromArray(QUOTEDSTRINGS)
	equalityExpr1 := strings.Join(random.RandomElementsFromArray(QUOTEDSTRINGS, 2), "!=")
	equalityExpr2 := randomWord + " == " + randomWord
	equalityExpr3 := fmt.Sprintf("%d == %d", getRandInt(), getRandInt())
	equalityExpr4 := fmt.Sprintf("(%d != %d) == ((-%d + %d) >= (%d * %d))", getRandInt(), getRandInt(), getRandInt(), getRandInt(), getRandInt(), getRandInt())
	parseTestCase := testcases.MultiParseTestCase{
		FileContents: []string{equalityExpr1, equalityExpr2, equalityExpr3, equalityExpr4},
	}
	return parseTestCase.RunAll(b, logger)
}
