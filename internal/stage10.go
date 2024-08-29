package internal

import (
	"slices"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testWhitespace(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	ws1 := SP
	ws2 := SP + TAB + LF + SP
	ws3 := "{" + strings.Join(random.RandomElementsFromArray(WHITESPACES, 2), "") + "}" + LF + "((" + strings.Join(random.RandomElementsFromArray(slices.Concat(SINGLE_CHAR_OPERATORS, WHITESPACES), 5), "") + "))"
	ws4 := "{" + strings.Join(random.RandomElementsFromArray(WHITESPACES, 5), "") + "}" + LF + "((" + strings.Join(random.RandomElementsFromArray(slices.Concat(SINGLE_CHAR_OPERATORS, RELATIONALS, WHITESPACES), 5), "") + "))"
	tokenizeTestCases := testcases.MultiTokenizeTestCase{
		TestCases: []testcases.TokenizeTestCase{
			{FileContents: ws1, ExpectsError: false},
			{FileContents: ws2, ExpectsError: false},
			{FileContents: ws3, ExpectsError: false},
			{FileContents: ws4, ExpectsError: false},
		},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
