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
	// A test string of length 4 is longer than the length of WHITESPACES (3), making sure the test case raises error.
	multiLineErrors2 := strings.Join(random.RandomElementsFromArray(slices.Concat(LEXICAL_ERRORS, WHITESPACES), 4), "")
	multiLineErrors3 := `()  #	{}
@
$
+++
// Let's Go!
+++
#`
	multiLineErrors4 := "({" + random.RandomElementFromArray(SINGLE_CHAR_OPERATORS) + random.RandomElementFromArray(WHITESPACES) + random.RandomElementFromArray(LEXICAL_ERRORS) + "})"
	tokenizeTestCases := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.TokenizeTestCase{FileContents: multiLineErrors1, ExpectsError: true},
			&testcases.TokenizeTestCase{FileContents: multiLineErrors2, ExpectsError: true},
			&testcases.TokenizeTestCase{FileContents: multiLineErrors3, ExpectsError: true},
			&testcases.TokenizeTestCase{FileContents: multiLineErrors4, ExpectsError: true},
		},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
