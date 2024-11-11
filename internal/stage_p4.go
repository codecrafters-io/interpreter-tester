package internal

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testParseParens(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	parens1 := "(\"foo\")"
	parens2 := "((true))"
	parens3 := "(nil)"
	parens4 := fmt.Sprintf("(%d.%d)", getRandInt(), getRandInt())
	parseTestCase := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.ParseTestCase{FileContents: parens1, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: parens2, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: parens3, ExpectsError: false},
			&testcases.ParseTestCase{FileContents: parens4, ExpectsError: false},
		},
	}
	return parseTestCase.RunAll(b, logger)
}
