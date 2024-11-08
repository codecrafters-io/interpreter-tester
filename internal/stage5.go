package internal

import (
	"slices"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testErrors(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	shuffledString1 := strings.Join(random.RandomElementsFromArray(LEXICAL_ERRORS, 5), "")
	// A test string of length 7 is longer than the length of SINGLE_CHAR_OPERATORS (6), making sure the test case raises error.
	shuffledString2 := "{(" + strings.Join(random.RandomElementsFromArray(slices.Concat(SINGLE_CHAR_OPERATORS, LEXICAL_ERRORS), 7), "") + ")}"
	tokenizeTestCases := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.TokenizeTestCase{FileContents: "@", ExpectsError: true},
			&testcases.TokenizeTestCase{FileContents: ",.$(#", ExpectsError: true},
			&testcases.TokenizeTestCase{FileContents: shuffledString1, ExpectsError: true},
			&testcases.TokenizeTestCase{FileContents: shuffledString2, ExpectsError: true},
		},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
