package internal

import (
	"slices"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var Relational = []string{"<", ">", "<=", ">="}

func testRelational(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	shuffledString1 := joinWith(random.RandomElementsFromArray(Relational, 5), "")
	shuffledString2 := "(){" + joinWith(random.RandomElementsFromArray(slices.Concat(LexicalErrors, Equals, Negation, Relational), 3), "") + "}"

	tokenizeTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{">=", "<<<=>>>=", shuffledString1, shuffledString2},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
