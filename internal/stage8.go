package internal

import (
	"slices"
	"strings"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testRelational(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	shuffledString1 := strings.Join(random.RandomElementsFromArray(RELATIONALS, 5), "")
	shuffledString2 := "(){" + strings.Join(random.RandomElementsFromArray(slices.Concat(LEXICALERRORS, EQUALS, NEGATIONS, RELATIONALS), 3), "") + "}"

	tokenizeTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{">=", "<<<=>>>=", shuffledString1, shuffledString2},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
