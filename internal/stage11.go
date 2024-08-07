package internal

import (
	"slices"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testErrorsMulti(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	multiLineErrors1 := `() 
	@`
	multiLineErrors2 := strings.Join(random.RandomElementsFromArray(slices.Concat(LEXICAL_ERRORS, WHITESPACES), 3), "")
	multiLineErrors3 := `()  #	{}
@
$
+++
// Let's Go!
+++
#`
	multiLineErrors4 := "({" + random.RandomElementFromArray(SINGLE_CHAR_OPERATORS) + random.RandomElementFromArray(WHITESPACES) + random.RandomElementFromArray(LEXICAL_ERRORS) + "})"
	tokenizeTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{multiLineErrors1, multiLineErrors2, multiLineErrors3, multiLineErrors4},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
