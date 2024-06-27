package internal

import (
	"github.com/codecrafters-io/grep-tester/internal/interpreter_executable"
	testcases "github.com/codecrafters-io/grep-tester/internal/test_cases"

	"github.com/codecrafters-io/tester-utils/random"
	"github.com/codecrafters-io/tester-utils/test_case_harness"
)

func testEquality(stageHarness *test_case_harness.TestCaseHarness) error {
	b := interpreter_executable.NewInterpreterExecutable(stageHarness)

	logger := stageHarness.Logger

	shuffledString1 := random.RandomStringFromCharacters(20, []rune("=(=)="))
	commandTestCases := testcases.MultiTokenizeTestCase{
		FileContents: []string{"Hello\n\n\n\n\n\n\nWORLD!", "{=}{====}", "(=)#(=====)", shuffledString1},
	}
	if err := commandTestCases.RunAll(b, logger); err != nil {
		return err
	}

	return nil
}
