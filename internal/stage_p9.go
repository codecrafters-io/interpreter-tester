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

	randomWord := random.RandomElementFromArray(QUOTED_STRINGS)
	equalityExpr1 := strings.Join(random.RandomElementsFromArray(QUOTED_STRINGS, 2), "!=")
	equalityExpr2 := randomWord + " == " + randomWord
	equalityExpr3 := fmt.Sprintf("%d == %d", getRandInt(), getRandInt())
	equalityExpr4 := fmt.Sprintf("(%d != %d) == ((-%d + %d) >= (%d * %d))", getRandInt(), getRandInt(), getRandInt(), getRandInt(), getRandInt(), getRandInt())
	parseTestCase := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.ParseTestCase{FileContents: equalityExpr1, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: equalityExpr2, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: equalityExpr3, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: equalityExpr4, ExpectsError: false},
		},
	}
	return parseTestCase.RunAll(b, logger)
}
