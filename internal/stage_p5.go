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

	unaryExpr1 := "!" + getRandBoolean()
	unaryExpr2 := "-" + fmt.Sprint(random.RandomInt(10, 100))
	unaryExpr3 := "!!" + getRandBoolean()
	unaryExpr4 := "(!!(" + getRandBoolean() + "))"
	parseTestCase := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.ParseTestCase{FileContents: unaryExpr1, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: unaryExpr2, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: unaryExpr3, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: unaryExpr4, ExpectsError: false},
		},
	}
	return parseTestCase.RunAll(b, logger)
}
