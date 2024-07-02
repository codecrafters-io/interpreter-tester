package internal

import (
	"slices"

	"github.com/codecrafters-io/grep-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/grep-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testErrorsMulti(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	multiLineErrors1 := `() 
	@`
	multiLineErrors2 := randomStringFromCharactersNew(10, slices.Concat(LexicalErrors, Whitespace))
	multiLineErrors3 := `()  #	{}
	@ @ @
	$
	+++
	// Let's Go!
	+++
	#
	#
	#`
	multiLineErrors4 := randomStringFromCharactersNew(40, slices.Concat(Parens, Braces, SingleCharOperators, LexicalErrors, RepeatSlice(Whitespace, 5)))
	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{multiLineErrors1, multiLineErrors2, multiLineErrors3, multiLineErrors4},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
