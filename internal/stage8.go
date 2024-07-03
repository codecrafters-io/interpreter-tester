package internal

import (
	"slices"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var Relational = []string{"<", ">", "<=", ">="}

func testRelational(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	shuffledString1 := randomStringFromCharacters(5, Relational)
	shuffledString2 := "(){" + randomStringFromCharacters(3, slices.Concat(LexicalErrors, Equals, Negation, Relational)) + "}"

	tokenizeTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{">=", "<<<=>>>=", shuffledString1, shuffledString2},
	}
	return tokenizeTestCases.RunAll(b, logger)
}
