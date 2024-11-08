package internal

import (
	"slices"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testComments(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	// As whitespace is not introduced yet, we skip multi-line comments here.
	comment1 := "//Comment"
	comment2 := "(///Unicode:£§᯽☺♣)"
	division1 := "/"
	division2 := "({(" + strings.Join(random.RandomElementsFromArray(slices.Concat(SINGLE_CHAR_OPERATORS, EQUALS, NEGATIONS, RELATIONALS), 3), "") + ")})" + "//Comment"
	tokenizeTestCases := testcases.MultiTestCase{
		TestCases: []testcases.TestCase{
			&testcases.TokenizeTestCase{FileContents: comment1, ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: comment2, ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: division1, ExpectsError: false},
			&testcases.TokenizeTestCase{FileContents: division2, ExpectsError: false},
		},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
