package internal

import (
	"slices"

	"github.com/codecrafters-io/interpreter-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/interpreter-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testErrorsMulti(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	multiLineErrors1 := `() 
	@`
	multiLineErrors2 := random.RandomStringFromCharacters(3, slices.Concat(LexicalErrors, Whitespace))
	multiLineErrors3 := `()  #	{}
@
$
+++
// Let's Go!
+++
#`
	multiLineErrors4 := "({" + random.RandomSelection(1, SingleCharOperators, "") + random.RandomSelection(1, Whitespace, "") + random.RandomSelection(1, LexicalErrors, "") + "})"
	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{multiLineErrors1, multiLineErrors2, multiLineErrors3, multiLineErrors4},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
