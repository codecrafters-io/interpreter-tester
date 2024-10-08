package internal

import (
	"slices"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testSingleChars(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	shuffledString1 := strings.Join(random.RandomElementsFromArray(SINGLE_CHAR_OPERATORS, 7), "")
	shuffledString2 := "({" + strings.Join(random.RandomElementsFromArray(slices.Concat(SINGLE_CHAR_OPERATORS), 5), "") + "})"
	tokenizeTestCases := testcases.MultiTokenizeTestCase{
		TestCases: []testcases.TokenizeTestCase{
			{FileContents: "+-", ExpectsError: false},
			{FileContents: "++--**..,,;;", ExpectsError: false},
			{FileContents: shuffledString1, ExpectsError: false},
			{FileContents: shuffledString2, ExpectsError: false},
		},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
