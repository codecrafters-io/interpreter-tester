package internal

import (
	"slices"

	"github.com/codecrafters-io/grep-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/grep-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var Equals = []string{"=", "=="}

func testEquality(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	shuffledString1 := random.RandomStringFromCharacters(20, slices.Concat(Parens, Braces, SingleCharOperators, LexicalErrors, Equals))
	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{"==", "{=====}", "(=)#(=====)", shuffledString1},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
