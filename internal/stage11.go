package internal

import (
	"slices"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testErrorsMulti(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	multiLineErrors1 := `() 
	@`
	multiLineErrors2 := randomStringFromCharacters(3, slices.Concat(LexicalErrors, Whitespace))
	multiLineErrors3 := `()  #	{}
@
$
+++
// Let's Go!
+++
#`
	multiLineErrors4 := "({" + randomSelection(1, SingleCharOperators, "") + randomSelection(1, Whitespace, "") + randomSelection(1, LexicalErrors, "") + "})"
	tokenizeTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{multiLineErrors1, multiLineErrors2, multiLineErrors3, multiLineErrors4},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
