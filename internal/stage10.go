package internal

import (
	"slices"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

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
	ws3 := "{" + randomStringFromCharacters(2, Whitespace) + "}" + lf + "((" + randomStringFromCharacters(5, slices.Concat(SingleCharOperators, Whitespace)) + "))"
	ws4 := "{" + randomStringFromCharacters(5, Whitespace) + "}" + lf + "((" + randomStringFromCharacters(5, slices.Concat(SingleCharOperators, Relational, Whitespace)) + "))"
	tokenizeTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{ws1, ws2, ws3, ws4},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
