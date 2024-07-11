package internal

import (
	"slices"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var Whitespace = []string{sp, tab, lf}
var sp = " "
var tab = "	"
var lf = "\n"

func testWhitespace(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	ws1 := sp
	ws2 := sp + tab + lf + sp
	ws3 := "{" + joinWith(random.RandomElementsFromArray(Whitespace, 2), "") + "}" + lf + "((" + joinWith(random.RandomElementsFromArray(slices.Concat(SingleCharOperators, Whitespace), 5), "") + "))"
	ws4 := "{" + joinWith(random.RandomElementsFromArray(Whitespace, 5), "") + "}" + lf + "((" + joinWith(random.RandomElementsFromArray(slices.Concat(SingleCharOperators, Relational, Whitespace), 5), "") + "))"
	tokenizeTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{ws1, ws2, ws3, ws4},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
