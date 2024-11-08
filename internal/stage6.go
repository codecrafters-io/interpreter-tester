package internal

import (
	"slices"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEquality(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger
	// A test string of length 5 is longer than the length of EQUALS (2), so it's certain the test case raises error.
	shuffledString1 := "((" + strings.Join(random.RandomElementsFromArray(slices.Concat(LEXICAL_ERRORS, EQUALS), 5), "") + "))"
	tokenizeTestCases := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.TokenizeTestCase{FileContents: "=", ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: "==", ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: "({=}){==}", ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: shuffledString1, ExpectsError: true},
		},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
