package internal

import (
	"slices"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

var Division = []string{"/"}

func testComments(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	// As whitespace is not introduced yet, we skip multi-line comments here.
	comment1 := "//Comment"
	comment2 := "(///Unicode:£§᯽☺♣)"
	division1 := "/"
	division2 := "({(" + random.RandomStringFromCharacters(3, slices.Concat(SingleCharOperators, LexicalErrors, Equals, Negation, Relational)) + ")})" + "//Comment"
	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{comment1, comment2, division1, division2},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
