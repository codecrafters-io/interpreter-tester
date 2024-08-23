package internal

import (
	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testParseParens(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	parens1 := "(\"foo\")"
	parens2 := "((true))"
	parens3 := "()" // TODO: Fix
	parens4 := "(\"foo\""
	parseTestCase := testcases.MultiParseTestCase{
		TestCases: []testcases.ParseTestCase{
			{FileContents: parens1, ExpectsError: false},
			{FileContents: parens2, ExpectsError: false},
			{FileContents: parens3, ExpectsError: false},
			{FileContents: parens4, ExpectsError: true},
		},
	}
	return parseTestCase.RunAll(b, logger)
}
